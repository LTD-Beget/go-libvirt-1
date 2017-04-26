package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strings"
	"sync/atomic"
	"time"

	"github.com/davecgh/go-xdr/xdr2"
	"github.com/vtolstov/go-libvirt"
)

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
	for _, s := range l.streams {
		s.Abort()
	}
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
		if err != nil {
			// When the underlying connection EOFs or is closed, stop
			// this goroutine
			if err == io.EOF || strings.Contains(err.Error(), "use of closed network connection") {
				return
			}
			if err, ok := err.(net.Error); ok && err.Timeout() {
				continue
			} else {
				panic(fmt.Sprintf("invalid packet %s", err))
			}
		}

		// response header
		h, err := extractHeader(l.r)
		if err != nil {
			panic(fmt.Sprintf("inv packet: %s\n", err))
		}
		// payload: packet length minus what was previously read
		size := int(length) - (libvirt.NetMessageHeaderXdrLen + libvirt.NetMessageHeaderMax)
		if size > 0 {
			buf := make([]byte, size)
			_, err = io.ReadFull(l.r, buf)
			if err != nil {
				panic(err)
			}
			l.route(h, buf)
		} else {
			l.route(h, nil)
		}
	}
}

// callback sends rpc responses to their respective caller.
func (l *Libvirt) callback(res libvirt.Message) {
	c, ok := l.callbacks[res.Header.Serial]
	if ok {
		c <- res
	}

	l.deregister(res.Header.Serial)
}

// streamRead reads rpc responses to their respective caller without deregister
func (l *Libvirt) streamRead(res libvirt.Message) {
	l.sm.Lock()
	s, ok := l.streams[res.Header.Serial]
	l.sm.Unlock()
	if ok {
		s.msg <- res
	}
}

// messageRead reads rpc responses to their respective caller without deregister
func (l *Libvirt) messageRead(res libvirt.Message) {
	event, err := decodeEvent(res)
	if err != nil {
		panic(err)
	}
	l.em.Lock()
	s, ok := l.messages[event.CallbackID]
	l.em.Unlock()
	if ok {
		s <- event
	}
}

// route sends incoming packets to their listeners.
func (l *Libvirt) route(h *libvirt.MessageHeader, payload []byte) {
	//fmt.Printf("route %d %s\n", len(payload), h)
	switch h.Type {
	case libvirt.MessageTypeStream:
		l.streamRead(libvirt.NewMessage(h, payload))
	case libvirt.MessageTypeMessage:
		l.messageRead(libvirt.NewMessage(h, payload))
	default:
		l.callback(libvirt.NewMessage(h, payload))
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
	//close(l.streams[id].done)
	delete(l.streams, id)
	l.sm.Unlock()
}

// addEvent add event channel
func (l *Libvirt) addEvent(id uint32, c chan *libvirt.Event) {
	l.em.Lock()
	l.messages[id] = c
	l.em.Unlock()
}

// delEvent del event channel
func (l *Libvirt) delEvent(id uint32) {
	l.em.Lock()
	delete(l.messages, id)
	l.em.Unlock()
}

// register configures a method response callback
func (l *Libvirt) register(id uint32, c chan libvirt.Message) {
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
func (l *Libvirt) send(proc libvirt.RemoteProcedure, serial uint32, mtype libvirt.MessageType, program uint32, status libvirt.MessageStatus, payload *bytes.Buffer) (<-chan libvirt.Message, error) {
	var c chan libvirt.Message = nil

	if serial == 0 {
		serial = l.serial()
	}
	if program == 0 {
		program = libvirt.RemoteProgram
	}

	switch mtype {
	case libvirt.MessageTypeCall:
		c = make(chan libvirt.Message)
		l.register(serial, c)
	}

	size := libvirt.NetMessageHeaderXdrLen + libvirt.NetMessageHeaderMax
	if payload != nil {
		size += payload.Len()
	}

	p := libvirt.Packet{
		Length: uint32(size),
		Header: libvirt.MessageHeader{
			Program:   program,
			Version:   libvirt.RemoteProtocolVersion,
			Procedure: proc,
			Type:      mtype,
			Serial:    serial,
			Status:    status,
		}}

	buf := bytes.NewBuffer(nil)

	// write header
	err := binary.Write(buf, binary.BigEndian, p)
	if err != nil {
		return nil, err
	}
	// write payload
	if payload != nil {
		if mtype != libvirt.MessageTypeStream {
			err = binary.Write(buf, binary.BigEndian, payload.Bytes())
		} else {
			_, err = buf.Write(payload.Bytes())
		}
		if err != nil {
			return nil, err
		}
	}

Loop:
	for {
		l.conn.SetWriteDeadline(time.Now().Add(20 * time.Second))
		_, err := l.w.Write(buf.Bytes())
		switch err {
		case nil, io.EOF:
			break Loop
		default:
			if err, ok := err.(net.Error); ok && err.Timeout() {
				continue
			} else {
				break Loop
			}
		}
	}

	if err := l.w.Flush(); err != nil {
		return nil, err
	}
	return c, nil
}

// encode XDR encodes the provided data.
func encode(data interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	_, err := xdr.Marshal(&buf, data)

	return buf, err
}

// decodeError extracts an error message from the provider buffer.
func decodeError(buf []byte) error {
	//	res := libvirt.RemoteError{}
	res := struct {
		Code     int
		DomainID int
		Message  *string `xdr:"optional"`
		Level    int
	}{}

	// TODO: fix error parsing
	dec := xdr.NewDecoder(bytes.NewReader(buf))
	_, err := dec.Decode(&res)
	if err != nil {
		return err
	}

	if strings.Contains(*res.Message, "unknown procedure") || strings.Contains(*res.Message, "unsupported event") {
		return libvirt.ErrUnsupported
	}

	return fmt.Errorf(*res.Message)
}

// decodeEvent extracts an event from the given byte slice.
// Errors encountered will be returned along with a nil event.
func decodeEvent(res libvirt.Message) (*libvirt.Event, error) {
	var err error
	evt := libvirt.Event{}

	dec := xdr.NewDecoder(bytes.NewReader(res.Payload))

	msg := libvirt.LookupMsgTypeByProc(res.Header.Procedure)
	if msg == nil {
		return nil, libvirt.ErrUnsupported
	}

	cid, _, err := dec.DecodeInt()
	if err != nil {
		return nil, err
	}

	evt.CallbackID = uint32(cid)

	_, err = dec.Decode(&msg)
	if err != nil {
		return nil, err
	}
	evt.Msg = &msg

	return &evt, nil
}

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
func extractHeader(r io.Reader) (*libvirt.MessageHeader, error) {
	var h libvirt.MessageHeader
	err := binary.Read(r, binary.BigEndian, &h)
	return &h, err
}

func decodeTyped(dec *xdr.Decoder) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	cnt, _, err := dec.DecodeUint()
	if err != nil {
		return nil, err
	}
	for idx := uint32(0); idx < cnt; idx++ {
		tpname, _, err := dec.DecodeString()
		if err != nil {
			return nil, err
		}
		tptype, _, err := dec.DecodeUint()
		if err != nil {
			return nil, err
		}
		switch libvirt.TypedParamTypes(tptype) {
		case libvirt.TypedParamTypeINT:
			params[tpname], _, err = dec.DecodeInt()
		case libvirt.TypedParamTypeUINT:
			params[tpname], _, err = dec.DecodeUint()
		case libvirt.TypedParamTypeSTRING:
			params[tpname], _, err = dec.DecodeString()
		case libvirt.TypedParamTypeBOOLEAN:
			params[tpname], _, err = dec.DecodeBool()
		case libvirt.TypedParamTypeDOUBLE:
			params[tpname], _, err = dec.DecodeDouble()
		case libvirt.TypedParamTypeLLONG:
			params[tpname], _, err = dec.DecodeHyper()
		case libvirt.TypedParamTypeULLONG:
			params[tpname], _, err = dec.DecodeUhyper()
		}
		if err != nil {
			return nil, err
		}
	}
	return params, nil
}
