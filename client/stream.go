package client

import (
	"bytes"
	"fmt"
	"io"

	libvirt "github.com/vtolstov/go-libvirt"
)

// Stream represents a stream in ibvirt.
type Stream struct {
	rwbuf     []byte
	buf       *bytes.Buffer
	procedure libvirt.RemoteProcedure
	serial    uint32
	l         *Libvirt
	r         chan []byte
	w         chan []byte
	done      chan bool
	msg       chan libvirt.Message
	err       error
}

// StreamNew creates new stream.
func (l *Libvirt) StreamNew() (*Stream, error) {
	s := &Stream{
		l:     l,
		buf:   bytes.NewBuffer(nil),
		r:     make(chan []byte),
		w:     make(chan []byte),
		done:  make(chan bool),
		msg:   make(chan libvirt.Message),
		rwbuf: make([]byte, libvirt.NetMessageLegacyPayloadMax),
	}
	go s.io()

	return s, nil
}

func (s *Stream) io() {
	for {
		select {
		case msg := <-s.msg:
			switch msg.Header.Status {
			case libvirt.MessageStatusError:
				fmt.Printf("status err\n")
				s.err = io.ErrUnexpectedEOF
				close(s.r)
				close(s.w)
				close(s.done)
				s.l.delStream(s.serial)
				return
			case libvirt.MessageStatusOK:
				fmt.Printf("status ok\n")
				s.err = io.EOF
				close(s.r)
				close(s.w)
				close(s.done)
				s.l.delStream(s.serial)
				return
			case libvirt.MessageStatusContinue:
				fmt.Printf("status cont\n")
				if len(msg.Payload) > 0 {
					s.err = nil
					s.r <- msg.Payload
				} else {
					s.Finish()
					s.err = io.EOF
					continue
				}
			}
		}
	}
}

// Abort forced closes stream.
func (s *Stream) Abort() error {
	_, err := s.l.send(s.procedure, s.serial, libvirt.MessageTypeStream, libvirt.RemoteProgram, libvirt.MessageStatusError, nil)
	if err != nil {
		return err
	}
	return nil
}

// Finish shutdown stream.
func (s *Stream) Finish() error {
	_, err := s.l.send(s.procedure, s.serial, libvirt.MessageTypeStream, libvirt.RemoteProgram, libvirt.MessageStatusOK, nil)
	if err != nil {
		return err
	}
	return nil
}

// Read reads from stream
func (s *Stream) Read(p []byte) (int, error) {
	var err error
	var n int
	var c int

	n, err = s.buf.Read(p)
	if n == len(p) {
		return n, err
	}

Loop:
	for {
		select {
		case <-s.done:
			err = s.err
			s.buf.Reset()
			break Loop
		case buf := <-s.r:
			s.buf.Write(buf)
			if s.buf.Len() >= len(p) {
				c, err = s.buf.Read(p[n:])
				n += c
				if n >= len(p) {
					break Loop
				}
			}
		}
	}
	return n, err
}

// Write writes to stream
func (s *Stream) Write(p []byte) (int, error) {
	n := len(p)
	s.buf.Write(p)
	defer s.buf.Reset()
Loop:
	for {
		select {
		case <-s.done:
			return 0, io.EOF
		default:
			_, err := s.l.send(s.procedure, s.serial, libvirt.MessageTypeStream, libvirt.RemoteProgram, libvirt.MessageStatusContinue, s.buf)
			if err != nil {
				return n - s.buf.Len(), err
			}
			break Loop
		}
	}
	return n, nil
}

// Close closes stream, to implement io.Closer interface
func (s *Stream) Close() error {
	err := s.Finish()
	fmt.Printf("finish\n")
	if err != nil {
		return err
	}
	return nil
}
