package libvirt

import "fmt"

//go:generate go run gen.go
//go:generate goimports -w gen_lxc_protocol.go gen_qemu_protocol.go gen_remote_protocol.go gen_virkeepaliveprotocol.go gen_virnetprotocol.go
//go:generate stringer -type=RemoteProcedure,MessageStatus,MessageType -output internal_string.go
//go:generate stringer -type=StoragePoolsFlags -output storagepool_string.go
//go:generate stringer -type=StorageVolumeDownloadFlags,StorageVolumeUploadFlags,StorageVolumeCreateFlags,StorageVolumeDeleteFlags,StorageVolumeResizeFlags -output storagevolume_string.go
//go:generate stringer -type=DomainSnapshotCreateFlags,DomainSnapshotDeleteFlags,DomainSnapshotListFlags,DomainSnapshotRevertFlags -output domainsnapshot_string.go

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

// StoragePoolsFlags specifies storage pools to list.
type StoragePoolsFlags uint32

// These flags come in groups; if all bits from a group are 0,
// then that group is not used to filter results.
const (
	StoragePoolsFlagInactive StoragePoolsFlags = 1 << iota
	StoragePoolsFlagActive

	StoragePoolsFlagPersistent
	StoragePoolsFlagTransient

	StoragePoolsFlagAutostart
	StoragePoolsFlagNoAutostart

	// pools by type
	StoragePoolsFlagDir
	StoragePoolsFlagFS
	StoragePoolsFlagNETFS
	StoragePoolsFlagLogical
	StoragePoolsFlagDisk
	StoragePoolsFlagISCSI
	StoragePoolsFlagSCSI
	StoragePoolsFlagMPATH
	StoragePoolsFlagRBD
	StoragePoolsFlagSheepdog
	StoragePoolsFlagGluster
	StoragePoolsFlagZFS
)

// StorageVolumeDownloadFlags specifies options when performing a volume download.
type StorageVolumeDownloadFlags uint32

const (
	// StorageVolumeDownloadFlagNone default
	StorageVolumeDownloadFlagNone StorageVolumeDownloadFlags = 0

	// StorageVolumeDownloadFlagSparseStream use sparse stream.
	StorageVolumeDownloadFlagSparseStream StorageVolumeDownloadFlags = 1 << iota
)

// StorageVolumeUploadFlags specifies options when performing a volume upload.
type StorageVolumeUploadFlags uint32

const (
	// StorageVolumeUploadFlagNone default
	StorageVolumeUploadFlagNone StorageVolumeUploadFlags = 0

	// StorageVolumeUploadFlagSparseStream use sparse stream.
	StorageVolumeUploadFlagSparseStream StorageVolumeUploadFlags = 1 << iota
)

// StorageVolumeCreateFlags specifies options when performing a volume creation.
type StorageVolumeCreateFlags uint32

const (
	StorageVolumeCreateFlagNone StorageVolumeCreateFlags = iota
	// StorageVolumeCreateFlagPreallocMetadata preallocates metadata
	StorageVolumeCreateFlagPreallocMetadata

	// StorageVolumeCreateFlagReflink use btrfs light copy
	StorageVolumeCreateFlagReflink
)

// StorageVolumeDeleteFlags specifies options when performing a volume deletion.
type StorageVolumeDeleteFlags uint32

const (
	// StorageVolumeDeleteFlagNormal delete metadata only (fast)
	StorageVolumeDeleteFlagNormal StorageVolumeDeleteFlags = iota

	// StorageVolumeDeleteFlagZeroes clear all data to zeros (slow)
	StorageVolumeDeleteFlagZeroes

	// StorageVolumeDeleteFlagWithSnapshots force removal of volume, even if in use
	StorageVolumeDeleteFlagWithSnapshots
)

// StorageVolumeResizeFlags specifies options when performing a volume deletion.
type StorageVolumeResizeFlags uint32

const (
	StorageVolumeResizeFlagNone StorageVolumeResizeFlags = 1 << iota
	// StorageVolumeResizeFlagAllocate force allocation of new size.
	StorageVolumeResizeFlagAllocate

	// StorageVolumeResizeFlagDelta size is relative to current.
	StorageVolumeResizeFlagDelta

	// StorageVolumeResizeFlagShrink allow decrease in capacity.
	StorageVolumeResizeFlagShrink
)

// DomainSnapshotCreateFlags specifies options when performing a snapshot creation.
type DomainSnapshotCreateFlags uint32

const (
	// DomainSnapshotCreateFlagRedefine restore or alter metadata.
	DomainSnapshotCreateFlagRedefine DomainSnapshotCreateFlags = 1 << iota

	// DomainSnapshotCreateFlagCreateCurrent with redefine, make snapshot current.
	DomainSnapshotCreateFlagCreateCurrent

	// DomainSnapshotCreateFlagNoMetadata make snapshot without remembering it.
	DomainSnapshotCreateFlagNoMetadata

	// DomainSnapshotCreateFlagCreateHalt stop running guest after snapshot.
	DomainSnapshotCreateFlagCreateHalt

	// DomainSnapshotCreateFlagCreateDiskOnly disk snapshot, not system checkpoint.
	DomainSnapshotCreateFlagCreateDiskOnly

	// DomainSnapshotCreateFlagCreateReuseExt reuse any existing external files.
	DomainSnapshotCreateFlagCreateReuseExt

	// DomainSnapshotCreateFlagCreateQuiesce use guest agent to quiesce all mounted file systems within the domain.
	DomainSnapshotCreateFlagCreateQuiesce

	// DomainSnapshotCreateFlagCreateAtomic atomically avoid partial changes.
	DomainSnapshotCreateFlagCreateAtomic

	// DomainSnapshotCreateFlagCreateLive create the snapshot while the guest is running.
	DomainSnapshotCreateFlagCreateLive
)

// DomainSnapshotDeleteFlags specifies options when performing a snapshot deletion.
type DomainSnapshotDeleteFlags uint32

const (
	// DomainSnapshotDeleteFlagChildren also delete children.
	DomainSnapshotDeleteFlagChildren DomainSnapshotDeleteFlags = 1 << iota

	// DomainSnapshotDeleteFlagMetadataOnly delete just metadata.
	DomainSnapshotDeleteFlagMetadataOnly

	// DomainSnapshotDeleteFlagChildrenOnly delete just children.
	DomainSnapshotDeleteFlagChildrenOnly
)

type DomainSnapshotListFlags uint32

const (
	// DomainSnapshotListDescendants list all descendants, not just children, when listing a snapshot.
	DomainSnapshotListFlagDescendants DomainSnapshotListFlags = 1 << iota

	// DomainSnapshotListRoots filter by snapshots with no parents, when listing a domain.
	DomainSnapshotListFlagRoots

	// DomainSnapshotListMetadata filter by snapshots which have metadata.
	DomainSnapshotListFlagMetadata

	// DomainSnapshotListLeaves filter by snapshots with no children.
	DomainSnapshotListFlagLeaves

	// DomainSnapshotListNoLeaves filter by snapshots that have children.
	DomainSnapshotListFlagNoLeaves

	// DomainSnapshotListNoMetadata filter by snapshots with no metadata.
	DomainSnapshotListFlagNoMetadata

	// DomainSnapshotListInactive filter by snapshots taken while guest was shut off.
	DomainSnapshotListFlagInactive

	// DomainSnapshotListAactive filter by snapshots taken while guest was active, and with memory state.
	DomainSnapshotListFlagAactive

	// DomainSnapshotListDiskOnly filter by snapshots taken while guest was active, but without memory state.
	DomainSnapshotListFlagDiskOnly

	// DomainSnapshotListInternal filter by snapshots stored internal to disk images.
	DomainSnapshotListFlagInternal

	// DomainSnapshotListExternal filter by snapshots that use files external to disk images.
	DomainSnapshotListFlagExternal
)

type DomainSnapshotRevertFlags uint32

const (
	// DomainSnapshotRevertFlagRunning run after revert.
	DomainSnapshotRevertFlagRunning DomainSnapshotRevertFlags = 1 << iota

	// DomainSnapshotRevertFlagPaused pause after revert.
	DomainSnapshotRevertFlagPaused

	// DomainSnapshotRevertFlagForce allow risky reverts.
	DomainSnapshotRevertFlagForce
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
