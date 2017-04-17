package libvirt

import "fmt"

//go:generate go run gen.go
//go:generate goimports -w gen_lxc_protocol.go gen_qemu_protocol.go gen_remote_protocol.go gen_virkeepaliveprotocol.go gen_virnetprotocol.go
//go:generate stringer -type=RemoteProcedure
//go:generate stringer -type=MessageStatus
//go:generate stringer -type=MessageType

type UUID [16]byte

type RemoteAuthType uint32

type RemoteTypedParam struct {
	Type  int
	Value interface{}
}

// ErrUnsupported is returned if a procedure is not supported by libvirt
var ErrUnsupported = fmt.Errorf("unsupported procedure requested")

// MessageType type of message
type MessageType uint32

// request and response types
const (
	// MessageTypeCall is used when making calls to the remote server.
	MessageTypeCall MessageType = iota

	// MessageTypeReply indicates a server reply.
	MessageTypeReply

	// MessageTypeMessage is an asynchronous notification.
	MessageTypeMessage

	// MessageTypeStream represents a stream data packet.
	MessageTypeStream

	// MessageTypeCallWithFDs is used by a client to indicate the request has
	// arguments with file descriptors.
	MessageTypeCallWithFDs

	// MessageTypeReplyWithFDs is used by a server to indicate the request has
	// arguments with file descriptors.
	MessageTypeReplyWithFDs

	// MessageTypeStreamSkip represents a stream skip packet.
	MessageTypeStreamSkip
)

// MessageStatus status of message
type MessageStatus uint32

// request and response statuses
const (
	// MessageStatusOK is always set for method calls or events.
	// For replies it indicates successful completion of the method.
	// For streams it indicates confirmation of the end of file on the stream.
	MessageStatusOK MessageStatus = iota

	// MessageStatusError for replies indicates that the method call failed
	// and error information is being returned. For streams this indicates
	// that not all data was sent and the stream has aborted.
	MessageStatusError

	// MessageStatusContinue is only used for streams.
	// This indicates that further data packets will be following.
	MessageStatusContinue
)

const (
	SecurityModelBuflen = 256 + 1
	SecurityDoiBuflen   = 256 + 1
	SecurityLabelBuflen = 4096 + 1
)

type RemoteProcedure uint32

const (
	RemoteAuthNone RemoteAuthType = iota
	RemoteAuthSASL
	RemoteAuthPolkit
)
