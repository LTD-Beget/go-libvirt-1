package client

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"sync"

	"github.com/davecgh/go-xdr/xdr2"
	"github.com/vtolstov/go-libvirt"
)

// Libvirt implements LibVirt's remote procedure call protocol.
type Libvirt struct {
	conn net.Conn
	r    *bufio.Reader
	w    *bufio.Writer

	// callbacks
	cm        sync.Mutex
	callbacks map[uint32]chan libvirt.Message

	// streams
	sm      sync.Mutex
	streams map[uint32]*Stream

	// next request serial number
	s uint32
}

// Capabilities returns an XML document describing the host's capabilties.
func (l *Libvirt) Capabilities() ([]byte, error) {
	resp, err := l.send(libvirt.RemoteProcConnectGetCapabilities, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, nil)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	caps, _, err := dec.DecodeString()

	return []byte(caps), err
}

// Connect establishes communication with the libvirt server.
// The underlying libvirt socket connection must be previously established.
func (l *Libvirt) Connect() error {
	return l.connect()
}

func (l *Libvirt) Close() error {
	// inform libvirt we're done
	if err := l.disconnect(); err != nil {
		return err
	}
	return l.conn.Close()
}

// Version returns the version of the libvirt daemon.
func (l *Libvirt) Version() (string, error) {
	resp, err := l.send(libvirt.RemoteProcConnectGetLibVersion, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, nil)
	if err != nil {
		return "", err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return "", decodeError(r.Payload)
	}

	result := struct {
		Version uint64
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return "", err
	}

	// The version is provided as an int following this formula:
	// version * 1,000,000 + minor * 1000 + micro
	// See src/libvirt-host.c # virConnectGetLibVersion
	major := result.Version / 1000000
	result.Version %= 1000000
	minor := result.Version / 1000
	result.Version %= 1000
	micro := result.Version

	versionString := fmt.Sprintf("%d.%d.%d", major, minor, micro)
	return versionString, nil
}

// New configures a new Libvirt RPC connection.
func New(conn net.Conn) *Libvirt {
	l := &Libvirt{
		conn:      conn,
		s:         0,
		r:         bufio.NewReader(conn),
		w:         bufio.NewWriter(conn),
		callbacks: make(map[uint32]chan libvirt.Message),
		streams:   make(map[uint32]*Stream),
	}

	go l.listen()

	return l
}
