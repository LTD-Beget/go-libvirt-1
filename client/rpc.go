package client

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync/atomic"
	"time"

	"github.com/davecgh/go-xdr/xdr2"
	"github.com/vtolstov/go-libvirt"
)

// MessageHeader is a libvirt rpc packet header
type MessageHeader struct {
	// Program identifier
	Program uint32

	// Program version
	Version uint32

	// Remote procedure identifier
	Procedure libvirt.RemoteProcedure

	// Call type, e.g., Reply
	Type libvirt.MessageType

	// Call serial number
	Serial uint32

	// Request status, e.g., StatusOK
	Status libvirt.MessageStatus
}

// packet represents a RPC request or response.
type packet struct {
	// Size of packet, in bytes, including length.
	// Len + Header + Payload
	Len    uint32
	Header MessageHeader
}

type Message struct {
	Header  MessageHeader
	Payload []byte
}

func NewMessage(hdr *MessageHeader, payload []byte) Message {
	return Message{Payload: payload, Header: *hdr}
}

// libvirt error response
type libvirtError struct {
	Code     uint32
	DomainID uint32
	Padding  uint8
	Message  string
	Level    uint32
}

func (l *Libvirt) connect() error {
	payload := struct {
		Padding [3]byte
		Name    string
		Flags   uint32
	}{
		Padding: [3]byte{0x1, 0x0, 0x0},
		Name:    "qemu:///system",
		Flags:   0,
	}

	buf, err := encode(&payload)
	if err != nil {
		return err
	}

	// libvirt requires that we call auth-list prior to connecting,
	// event when no authentication is used.
	resp, err := l.send(libvirt.RemoteProcAuthList, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	resp, err = l.send(libvirt.RemoteProcConnectOpen, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r = <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

func (l *Libvirt) disconnect() error {
	resp, err := l.send(libvirt.RemoteProcConnectClose, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, nil)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// listen processes incoming data and routes
// responses to their respective callback handler.
func (l *Libvirt) listen() {
	for {
		l.conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		// response packet length
		length, err := pktlen(l.r)
		fmt.Printf("len %d\n", length)
		if err != nil {
			// When the underlying connection EOFs or is closed, stop
			// this goroutine
			if err == io.EOF || strings.Contains(err.Error(), "use of closed network connection") {
				return
			}

			panic("invalid packet")
		}

		// response header
		h, err := extractHeader(l.r)
		if err != nil {
			fmt.Printf("inv packet: %s\n", err)
			panic("error")
		}
		fmt.Printf("xxx %+v\n", h)
		// payload: packet length minus what was previously read
		size := int(length) - (libvirt.NetMessageHeaderXdrLen + libvirt.NetMessageHeaderMax)
		fmt.Printf("pkt size %d\n", size)
		buf := make([]byte, size)
		_, err = io.ReadFull(l.r, buf)
		if err != nil {
			fmt.Printf("ee %s\n", err)
			// invalid packet
			continue
		}
		l.route(h, buf)
	}
}

// callback sends rpc responses to their respective caller.
func (l *Libvirt) callback(res Message) {
	c, ok := l.callbacks[res.Header.Serial]
	if ok {
		c <- res
	}

	l.deregister(res.Header.Serial)
}

// streamRead sends rpc responses to their respective caller without deregister
func (l *Libvirt) streamRead(res Message) {
	l.sm.Lock()
	s, ok := l.streams[res.Header.Serial]
	l.sm.Unlock()
	if ok {
		s.msg <- res
	}
}

// route sends incoming packets to their listeners.
func (l *Libvirt) route(h *MessageHeader, payload []byte) {
	// route events to their respective listener
	//	if h.Program == constants.ProgramQEMU && h.Procedure == constants.QEMUDomainMonitorEvent {
	//		l.event()
	//		return
	//	}
	switch h.Type {
	case libvirt.MessageTypeStream:
		l.streamRead(NewMessage(h, payload))
	default:
		l.callback(NewMessage(h, payload))
	}
}

// serial provides atomic access to the next sequential request serial number.
func (l *Libvirt) serial() uint32 {
	return atomic.AddUint32(&l.s, 1)
}

// addStream add stream
func (l *Libvirt) addStream(id uint32, s *Stream) {
	l.sm.Lock()
	l.streams[id] = s
	l.sm.Unlock()
}

// delStream remove stream
func (l *Libvirt) delStream(id uint32) {
	l.sm.Lock()
	close(l.streams[id].done)
	delete(l.streams, id)
	l.sm.Unlock()
}

// register configures a method response callback
func (l *Libvirt) register(id uint32, c chan Message) {
	l.cm.Lock()
	l.callbacks[id] = c
	l.cm.Unlock()
}

// deregister destroys a method response callback
func (l *Libvirt) deregister(id uint32) {
	l.cm.Lock()
	close(l.callbacks[id])
	delete(l.callbacks, id)
	l.cm.Unlock()
}

// send performs a libvirt RPC request.
// The returned channel is used by the caller to receive the asynchronous
// call response. The channel is closed once a response has been sent.
func (l *Libvirt) send(proc libvirt.RemoteProcedure, serial uint32, mtype libvirt.MessageType, program uint32, status libvirt.MessageStatus, payload *bytes.Buffer) (<-chan Message, error) {
	var n int
	if serial == 0 {
		serial = l.serial()
	}
	if program == 0 {
		program = libvirt.RemoteProgram
	}

	c := make(chan Message)
	switch mtype {
	case libvirt.MessageTypeCall:
		l.register(serial, c)
	}
	fmt.Printf("RRR\n")
	size := libvirt.NetMessageHeaderXdrLen + libvirt.NetMessageHeaderMax
	fmt.Printf("%d %d %d\n", size, libvirt.NetMessageHeaderXdrLen, libvirt.NetMessageHeaderMax)
	if payload != nil {
		size += payload.Len()
	}

	p := packet{
		Len: uint32(size),
		Header: MessageHeader{
			Program:   program,
			Version:   libvirt.RemoteProtocolVersion,
			Procedure: proc,
			Type:      mtype,
			Serial:    serial,
			Status:    status,
		},
	}
	l.conn.SetWriteDeadline(time.Now().Add(20 * time.Second))
	fmt.Printf("EEE\n")
	// write header
	err := binary.Write(l.w, binary.BigEndian, p)
	if err != nil {
		return nil, err
	}

	// write payload
	if payload != nil {
		fmt.Printf("qqqqq\n")
		if mtype != libvirt.MessageTypeStream {
			fmt.Printf("non stream\n")
			err = binary.Write(l.w, binary.BigEndian, payload.Bytes())
		} else {
			fmt.Printf("stream \n")
			n, err = l.w.Write(payload.Bytes())
			fmt.Printf("stream copy %d %#+v\n", n, err)
		}
		fmt.Printf("yyyy\n")
		if err != nil {
			return nil, err
		}
	}

	if err := l.w.Flush(); err != nil {
		return nil, err
	}
	fmt.Printf("TTTT\n")
	switch mtype {
	case libvirt.MessageTypeCall:
		return c, nil
	}
	return nil, nil
}

// encode XDR encodes the provided data.
func encode(data interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	_, err := xdr.Marshal(&buf, data)

	return buf, err
}

// decodeError extracts an error message from the provider buffer.
func decodeError(buf []byte) error {
	var e libvirtError

	dec := xdr.NewDecoder(bytes.NewReader(buf))
	_, err := dec.Decode(&e)
	if err != nil {
		return err
	}

	if strings.Contains(e.Message, "unknown procedure") {
		return libvirt.ErrUnsupported
	}

	return errors.New(e.Message)
}

/*
// decodeEvent extracts an event from the given byte slice.
// Errors encountered will be returned along with a nil event.
func decodeEvent(buf []byte) (*DomainEvent, error) {
	var e DomainEvent

	dec := xdr.NewDecoder(bytes.NewReader(buf))
	_, err := dec.Decode(&e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
*/

// pktlen determines the length of an incoming rpc response.
// If an error is encountered reading the provided Reader, the
// error is returned and response length will be 0.
func pktlen(r io.Reader) (uint32, error) {
	var size uint32
	var err error

	if err = binary.Read(r, binary.BigEndian, &size); err == nil {
		if size > libvirt.NetMessageMax {
			err = fmt.Errorf("invalid packet %d > %d", size, libvirt.NetMessageMax)
		}
	}
	return size, err
}

// extractHeader returns the decoded header from an incoming response.
func extractHeader(r io.Reader) (*MessageHeader, error) {
	var h MessageHeader
	err := binary.Read(r, binary.BigEndian, &h)
	return &h, err
}
