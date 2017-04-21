package client

import (
	"bytes"
	"io"

	libvirt "github.com/vtolstov/go-libvirt"
)

// Stream represents a stream in ibvirt.
type Stream struct {
	//rwbuf     []byte
	rbuf      *bytes.Buffer
	wbuf      *bytes.Buffer
	procedure libvirt.RemoteProcedure
	serial    uint32
	l         *Libvirt
	r         chan []byte
	w         chan []byte
	done      chan bool
	msg       chan libvirt.Message
	err       error
	//	rm        sync.Mutex
	//	wm        sync.Mutex
}

// StreamNew creates new stream.
func (l *Libvirt) StreamNew() (*Stream, error) {
	s := &Stream{
		l:    l,
		err:  io.EOF,
		rbuf: bytes.NewBuffer(nil),
		wbuf: bytes.NewBuffer(nil),
		r:    make(chan []byte),
		w:    make(chan []byte),
		done: make(chan bool),
		msg:  make(chan libvirt.Message),
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
				s.err = io.ErrUnexpectedEOF
				if s.r != nil {
					close(s.r)
				}
				if s.w != nil {
					close(s.w)
				}
				close(s.done)
				s.l.delStream(s.serial)
				return
			case libvirt.MessageStatusOK:
				s.err = io.EOF
				if s.r != nil {
					close(s.r)
				}
				if s.w != nil {
					close(s.w)
				}
				close(s.done)
				s.l.delStream(s.serial)
				return
			case libvirt.MessageStatusContinue:
				if len(msg.Payload) > 0 {
					s.err = nil
					s.r <- msg.Payload
				} else {
					s.err = io.EOF
					close(s.r)
					s.r = nil
					close(s.w)
					s.w = nil
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
	<-s.done
	return nil
}

// Finish shutdown stream.
func (s *Stream) Finish() error {
	_, err := s.l.send(s.procedure, s.serial, libvirt.MessageTypeStream, libvirt.RemoteProgram, libvirt.MessageStatusOK, nil)
	if err != nil {
		return err
	}
	<-s.done
	return nil
}

// Read reads from stream
func (s *Stream) Read(p []byte) (int, error) {
	var n int
	if s.rbuf.Len() == 0 {
		buf := <-s.r
		_, _ = s.rbuf.Write(buf)
	}
	n, _ = s.rbuf.Read(p)
	return n, s.err
}

// Write writes to stream
func (s *Stream) Write(p []byte) (int, error) {
	if len(p) > 0 {
		c, _ := s.wbuf.Write(p)
		defer s.wbuf.Reset()
		_, err := s.l.send(s.procedure, s.serial, libvirt.MessageTypeStream, libvirt.RemoteProgram, libvirt.MessageStatusContinue, s.wbuf)
		return c, err
	}
	return 0, nil
}

// Close closes stream, to implement io.Closer interface
func (s *Stream) Close() error {
	return s.Finish()
}
