package libvirt

import "fmt"

//go:generate go run gen.go
//go:generate goimports -w gen_lxc_protocol.go gen_qemu_protocol.go gen_remote_protocol.go gen_virkeepaliveprotocol.go gen_virnetprotocol.go
//go:generate stringer -type=RemoteProcedure,MessageStatus,MessageType -output internal_string.go
//go:generate stringer -type=StoragePoolsFlags,StorageXmlFlags -output storagepool_string.go
//go:generate stringer -type=StorageVolumeInfoFlags,StorageVolumeDownloadFlags,StorageVolumeUploadFlags,StorageVolumeCreateFlags,StorageVolumeDeleteFlags,StorageVolumeResizeFlags,StorageVolumeWipeAlgFlags -output storagevolume_string.go
//go:generate stringer -type=DomainSnapshotCreateFlags,DomainSnapshotDeleteFlags,DomainSnapshotListFlags,DomainSnapshotRevertFlags -output domainsnapshot_string.go
//go:generate stringer -type=EventIDTypes,DomainConsoleFlags,DomainStatsTypes,ListDomainsFlags,GetDomainsStatsFlags,DomainAffectFlags,DomainDeviceModifyFlags,DomainBlockResizeFlags,DomainXMLFlags,DomainCreateFlags,DomainRebootFlags,DomainShutdownFlags,DomainMigrateFlags,DomainUndefineFlags,DomainDefineXMLFlags,DomainDestroyFlags -output domain_string.go
//go:generate stringer -type=TypedParamTypes  -output=consts_string.go

type UUID [16]byte

type RemoteAuthType uint32

// ErrUnsupported is returned if a procedure is not supported by libvirt
var ErrUnsupported = fmt.Errorf("unsupported procedure requested")

func LookupMsgTypeByProc(proc RemoteProcedure) interface{} {
	var msg interface{} = nil

	switch proc {
	case RemoteProcDomainEventCallbackLifecycle:
		msg = RemoteDomainEventCallbackLifecycleMsg{}
	case RemoteProcDomainEventCallbackReboot:
		msg = RemoteDomainEventCallbackRebootMsg{}
	case RemoteProcDomainEventCallbackRtcChange:
		msg = RemoteDomainEventCallbackRtcChangeMsg{}
	case RemoteProcDomainEventCallbackWatchdog:
		msg = RemoteDomainEventCallbackWatchdogMsg{}
	case RemoteProcDomainEventCallbackIoError:
		msg = RemoteDomainEventCallbackIoErrorMsg{}
	case RemoteProcDomainEventCallbackGraphics:
		msg = RemoteDomainEventCallbackGraphicsMsg{}
	case RemoteProcDomainEventCallbackIoErrorReason:
		msg = RemoteDomainEventCallbackIoErrorReasonMsg{}
	case RemoteProcDomainEventCallbackControlError:
		msg = RemoteDomainEventCallbackControlErrorMsg{}
	case RemoteProcDomainEventCallbackBlockJob:
		msg = RemoteDomainEventCallbackBlockJobMsg{}
	case RemoteProcDomainEventCallbackDiskChange:
		msg = RemoteDomainEventCallbackDiskChangeMsg{}
	case RemoteProcDomainEventCallbackTrayChange:
		msg = RemoteDomainEventCallbackTrayChangeMsg{}
	case RemoteProcDomainEventCallbackPmwakeup:
		msg = RemoteDomainEventCallbackPmwakeupMsg{}
	case RemoteProcDomainEventCallbackPmsuspend:
		msg = RemoteDomainEventCallbackPmsuspendMsg{}
	case RemoteProcDomainEventCallbackBalloonChange:
		msg = RemoteDomainEventCallbackBalloonChangeMsg{}
	case RemoteProcDomainEventCallbackPmsuspendDisk:
		msg = RemoteDomainEventCallbackPmsuspendDiskMsg{}
	case RemoteProcDomainEventCallbackDeviceRemoved:
		msg = RemoteDomainEventCallbackDeviceRemovedMsg{}
	case RemoteProcDomainEventCallbackTunable:
		msg = RemoteDomainEventCallbackTunableMsg{}
	case RemoteProcDomainEventCallbackAgentLifecycle:
		msg = RemoteDomainEventCallbackAgentLifecycleMsg{}
	case RemoteProcDomainEventCallbackDeviceAdded:
		msg = RemoteDomainEventCallbackDeviceAddedMsg{}
	case RemoteProcDomainEventCallbackMigrationIteration:
		msg = RemoteDomainEventCallbackMigrationIterationMsg{}
	case RemoteProcDomainEventCallbackJobCompleted:
		msg = RemoteDomainEventCallbackJobCompletedMsg{}
	case RemoteProcDomainEventCallbackDeviceRemovalFailed:
		msg = RemoteDomainEventCallbackDeviceRemovalFailedMsg{}
	case RemoteProcDomainEventCallbackMetadataChange:
		msg = RemoteDomainEventCallbackMetadataChangeMsg{}
	default:
		msg = nil
	}
	return msg
}

// EventTypes
type EventIDTypes int

const (
	// DomainEventIDLifecycle lifecycle events
	DomainEventIDLifecycle EventIDTypes = iota

	// DomainEventIDReboot reboot events
	DomainEventIDReboot

	// DomainEventIDRtcChange
	DomainEventIDRtcChange

	// DomainEventIDWatchdog
	DomainEventIDWatchdog

	// DomainEventIDIOError
	DomainEventIDIOError

	// DomainEventIDGraphics
	DomainEventIDGraphics

	// DomainEventIDIOErrorReason
	DomainEventIDIOErrorReason

	// DomainEventIDControlError
	DomainEventIDControlError

	// DomainEventIDBlockJob
	DomainEventIDBlockJob

	// DomainEventIDDiskChange
	DomainEventIDDiskChange

	// DomainEventIDTrayChange
	DomainEventIDTrayChange

	// DomainEventIDPmWakeup
	DomainEventIDPmWakeup

	// DomainEventIDPmSuspend
	DomainEventIDPmSuspend

	// DomainEventIDBalloonChange
	DomainEventIDBalloonChange

	// DomainEventIDPmSuspendDisk
	DomainEventIDPmSuspendDisk

	// DomainEventIDDeviceRemoved
	DomainEventIDDeviceRemoved

	// DomainEventIDBlockJob2
	DomainEventIDBlockJob2

	// DomainEventIDTunable
	DomainEventIDTunable

	// DomainEventIDAgentLifecycle
	DomainEventIDAgentLifecycle

	// DomainEventIDDeviceAdded
	DomainEventIDDeviceAdded

	// DomainEventIDMigrationIteration
	DomainEventIDMigrationIteration

	// DomainEventIDJobCompleted
	DomainEventIDJobCompleted

	// DomainEventIDDeviceRemovalFailed
	DomainEventIDDeviceRemovalFailed

	// DomainEventIDMetadataChange
	DomainEventIDMetadataChange

	// DomainEventIDLast
	DomainEventIDLast
)

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

// DomainDeviceModifyFlags specifies flags for device modify
type DomainDeviceModifyFlags uint32

const (
	// DomainDeviceModifyFlagConfig affect only config.
	DomainDeviceModifyFlagConfig DomainDeviceModifyFlags = iota

	// DomainDeviceModifyFlagCurrent affect current domain.
	DomainDeviceModifyFlagCurrent

	// DomainDeviceModifyFlagLive affect live running domain.
	DomainDeviceModifyFlagLive

	// DomainDeviceModifyFlagForce force remove device
	DomainDeviceModifyFlagForce
)

// TypedParamTypes
type TypedParamTypes uint32

const (
	_ TypedParamTypes = iota
	TypedParamTypeINT
	TypedParamTypeUINT
	TypedParamTypeLLONG
	TypedParamTypeULLONG
	TypedParamTypeDOUBLE
	TypedParamTypeBOOLEAN
	TypedParamTypeSTRING
	TypedParamTypeLAST
)

type RemoteTypedParamValue []byte

/*
type RemoteTypedParamValue struct {
	i  int
	ui uint
	l  int64
	ul uint64
	b  bool
	d  float64
	s  string
	sl []string
}
*/

// DomainAffectFlags specifies flags for domain modify
type DomainAffectFlags uint32

const (
	// DomainAffectFlagCurrent affect current domain.
	DomainAffectFlagCurrent DomainAffectFlags = iota

	// DomainAffectFlagLive affect live running domain.
	DomainAffectFlagLive

	// DomainAffectFlagConfig affect only config.
	DomainAffectFlagConfig
)

// DomainConsoleFlags
type DomainConsoleFlags uint32

const (
	_ DomainConsoleFlags = iota

	// DomainConsoleFlagForce
	DomainConsoleFlagForce

	// DomainConsoleFlagSafe
	DomainConsoleFlagSafe
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

// DomainStatsTypes
type DomainStatsTypes uint32

const (
	// DomainStatsState return domain state
	DomainStatsState DomainStatsTypes = 1 << iota

	// DomainStatsCpuTotal return domain CPU info
	DomainStatsCpuTotal

	// DomainStatsBalloon return domain balloon info
	DomainStatsBalloon

	// DomainStatsVcpu return domain virtual CPU info
	DomainStatsVcpu

	// DomainStatsInterface return domain interfaces info
	DomainStatsInterface

	// DomainStatsBlock return domain block info
	DomainStatsBlock

	// DomainStatsPerf return domain perf event info
	DomainStatsPerf
)

// GetDomainsStatsFlags
type GetDomainsStatsFlags uint32

const (
	// GetDomainsStatsActive
	GetDomainsStatsActive GetDomainsStatsFlags = 1 << iota

	// GetDomainsStatsInactive
	GetDomainsStatsInactive

	// ListDomainsStatsPersistent
	ListDomainsStatsPersistent

	// GetDomainsStatsTransient
	GetDomainsStatsTransient

	// GetDomainsStatsRunning
	GetDomainsStatsRunning

	// ListDomainsPaused
	ListDomainsStatsPaused

	// GetDomainsStatsShutoff
	GetDomainsStatsShutoff

	// GetDomainsStatsOther
	GetDomainsStatsOther

	// GetDomainsStatsBacking
	GetDomainsStatsBacking GetDomainsStatsFlags = 1073741824

	// GetDomainsStatsEnforce
	GetDomainsStatsEnforce GetDomainsStatsFlags = 2147483648
)

// ListDomainsFlags
type ListDomainsFlags uint32

const (
	// ListDomainsActive
	ListDomainsActive ListDomainsFlags = 1 << iota

	// ListDomainsInactive
	ListDomainsInactive

	// ListDomainsPersistent
	ListDomainsPersistent

	// ListDomainsTransient
	ListDomainsTransient

	// ListDomainsRunning
	ListDomainsRunning

	// ListDomainsPaused
	ListDomainsPaused

	// ListDomainsShutoff
	ListDomainsShutoff

	// ListDomainsOther
	ListDomainsOther

	// ListDomainsManagedsave
	ListDomainsManagedsave

	// ListDomainsNoManagedsave
	ListDomainsNoManagedsave

	// ListDomainsAutostart
	ListDomainsAutostart

	// ListDomainsNoAutostart
	ListDomainsNoAutostart

	// ListDomainsHasSnapshot
	ListDomainsHasSnapshot

	// ListDomainsNoSnapshot
	ListDomainsNoSnapshot
)

// StorageXmlFlags
type StorageXmlFlags uint32

const (
	// StorageXmlInactive dump inactive pool/volume information
	StorageXmlInactive StorageXmlFlags = 1
)

// StorageVolumeInfoFlags specifies options when performing a volume info.
type StorageVolumeInfoFlags uint32

const (
	// StorageVolumeInfoFlagAllocation
	StorageVolumeInfoFlagAllocation StorageVolumeInfoFlags = iota

	// StorageVolumeInfoFlagPhysical
	StorageVolumeInfoFlagPhysical
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

// DomainBlockResizeFlags specifies options for block resize.
type DomainBlockResizeFlags uint32

const (
	DomainBlockResizeFlagNone DomainBlockResizeFlags = iota
	// DomainBlockResizeFlagBytes specify size in bytes for BlockResize
	DomainBlockResizeFlagBytes
)

// DomainXMLFlags specifies options for dumping a domain's XML.
type DomainXMLFlags uint32

const (
	// DomainXMLFlagSecure dumps XML with sensitive information included.
	DomainXMLFlagSecure DomainXMLFlags = 1 << iota

	// DomainXMLFlagInactive dumps XML with inactive domain information.
	DomainXMLFlagInactive

	// DomainXMLFlagUpdateCPU dumps XML with guest CPU requirements according to the host CPU.
	DomainXMLFlagUpdateCPU

	// DomainXMLFlagMigratable dumps XML suitable for migration.
	DomainXMLFlagMigratable
)

// DomainCreateFlags specifies options when performing a domain creation.
type DomainCreateFlags uint32

const (
	// DomainCreateFlagNone is the default behavior.
	DomainCreateFlagNone DomainCreateFlags = 0

	// DomainCreateFlagPaused creates paused domain.
	DomainCreateFlagPaused DomainCreateFlags = 1 << (iota - 1)

	// DomainCreateFlagAutoDestroy destoy domain after libvirt connection closed.
	DomainCreateFlagAutoDestroy

	// DomainCreateFlagBypassCache avoid file system cache pollution.
	DomainCreateFlagBypassCache

	// DomainCreateFlagStartForceBoot boot, discarding any managed save
	DomainCreateFlagStartForceBoot

	// DomainCreateFlagStartValidate validate the XML document against schema
	DomainCreateFlagStartValidate
)

// DomainRebootFlags specifies options when performing a reboot.
type DomainRebootFlags uint32

const (
	// DomainRebootFlagDefault use hypervisor choice.
	DomainRebootFlagDefault DomainRebootFlags = 0

	// DomainRebootFlagACPI send ACPI event.
	DomainRebootFlagACPI DomainRebootFlags = 1 << (iota - 1)

	// DomainRebootFlagGuestAgent use guest agent.
	DomainRebootFlagGuestAgent

	// DomainRebootFlagInitctl use initctl.
	DomainRebootFlagInitctl

	// DomainRebootFlagSignal send a signal.
	DomainRebootFlagSignal

	// DomainRebootFlagParavirt use paravirt guest control.
	DomainRebootFlagParavirt
)

// DomainShutdownFlags specifies options when performing a shutdown.
type DomainShutdownFlags uint32

const (
	// DomainShutdownFlagDefault use hypervisor choice.
	DomainShutdownFlagDefault DomainShutdownFlags = 0

	// DomainShutdownFlagACPI send ACPI event.
	DomainShutdownFlagACPI DomainShutdownFlags = 1 << (iota - 1)

	// DomainShutdownFlagGuestAgent use guest agent.
	DomainShutdownFlagGuestAgent

	// DomainShutdownFlagInitctl use initctl.
	DomainShutdownFlagInitctl

	// DomainShutdownFlagSignal send a signal.
	DomainShutdownFlagSignal

	// DomainShutdownFlagParavirt use paravirt guest control.
	DomainShutdownFlagParavirt
)

// DomainMigrateFlags specifies options when performing a migration.
type DomainMigrateFlags uint64

const (
	// DomainMigrateFlagLive performs a zero-downtime live migration.
	DomainMigrateFlagLive DomainMigrateFlags = 1 << iota

	// DomainMigrateFlagPeerToPeer creates a direct source to destination control channel.
	DomainMigrateFlagPeerToPeer

	// DomainMigrateFlagTunneled tunnels migration data over the libvirtd connection.
	DomainMigrateFlagTunneled

	// DomainMigrateFlagPersistDestination will persist the VM on the destination host.
	DomainMigrateFlagPersistDestination

	// DomainMigrateFlagUndefineSource undefines the VM on the source host.
	DomainMigrateFlagUndefineSource

	// DomainMigrateFlagPaused will pause the remote side VM.
	DomainMigrateFlagPaused

	// DomainMigrateFlagNonSharedDisk migrate non-shared storage with full disk copy.
	DomainMigrateFlagNonSharedDisk

	// DomainMigrateFlagNonSharedIncremental migrate non-shared storage with incremental copy.
	DomainMigrateFlagNonSharedIncremental

	// DomainMigrateFlagChangeProtection prevents any changes to the domain configuration through the whole migration process.
	DomainMigrateFlagChangeProtection

	// DomainMigrateFlagUnsafe will force a migration even when it is considered unsafe.
	DomainMigrateFlagUnsafe

	// DomainMigrateFlagOffline is used to perform an offline migration.
	DomainMigrateFlagOffline

	// DomainMigrateFlagCompressed compresses data during migration.
	DomainMigrateFlagCompressed

	// DomainMigrateFlagAbortOnError will abort a migration on I/O errors encountered during migration.
	DomainMigrateFlagAbortOnError

	// DomainMigrateFlagAutoConverge forces convergence.
	DomainMigrateFlagAutoConverge

	// DomainMigrateFlagRDMAPinAll enables RDMA memory pinning.
	DomainMigrateFlagRDMAPinAll

	// DomainMigrateFlagPostcopy enable post-copy migration.
	DomainMigrateFlagPostcopy
)

// DomainUndefineFlags specifies options available when undefining a domain.
type DomainUndefineFlags uint32

const (
	// DomainUndefineFlagManagedSave removes all domain managed save data.
	DomainUndefineFlagManagedSave DomainUndefineFlags = 1 << iota

	// DomainUndefineFlagSnapshotsMetadata removes all domain snapshot metadata.
	DomainUndefineFlagSnapshotsMetadata

	// DomainUndefineFlagNVRAM removes all domain NVRAM files.
	DomainUndefineFlagNVRAM
)

// DomainDefineXMLFlags specifies options available when defining a domain.
type DomainDefineXMLFlags uint32

const (
	// DefineValidate validates the XML document against schema
	DefineValidate DomainDefineXMLFlags = 1
)

// DomainDestroyFlags specifies options available when destroying a domain.
type DomainDestroyFlags uint32

const (
	// DestroyFlagDefault default behavior, forcefully terminate the domain.
	DestroyFlagDefault DomainDestroyFlags = 1 << iota

	// DestroyFlagGraceful only sends a SIGTERM no SIGKILL.
	DestroyFlagGraceful
)

// StorageVolumeWipeAlgorithmFlags flags for different wiping algorithms
type StorageVolumeWipeAlgFlags uint32

const (
	// StorageVolumeWipeAlgZero 1-pass, all zeroes
	StorageVolumeWipeAlgZero StorageVolumeWipeAlgFlags = iota

	// StorageVolumeWipeAlgNNSA 4-pass NNSA Policy Letter NAP-14.1-C (XVI-8)
	StorageVolumeWipeAlgNNSA

	// StorageVolumeWipeAlgDOD 4-pass DoD 5220.22-M section 8-306 procedure
	StorageVolumeWipeAlgDOD

	// StorageVolumeWipeAlgBSI 9-pass method recommended by the German Center of Security in Information Technologies
	StorageVolumeWipeAlgBSI

	// StorageVolumeWipeAlgGUTMANN The canonical 35-pass sequence
	StorageVolumeWipeAlgGUTMANN

	// StorageVolumeWipeAlgSCHNEIER 7-pass method described by Bruce Schneier in "Applied Cryptography" (1996)
	StorageVolumeWipeAlgSCHNEIER

	// StorageVolumeWipeAlgPFITZNER7 7-pass random
	StorageVolumeWipeAlgPFITZNER7

	// StorageVolumeWipeAlgPFITZNER33 33-pass random
	StorageVolumeWipeAlgPFITZNER33

	// StorageVolumeWipeAlgRandom 1-pass random
	StorageVolumeWipeAlgRandom

	// StorageVolumeWipeAlgTrim 1-pass, trim all data on the volume by using TRIM or DISCARD
	StorageVolumeWipeAlgTrim

	// StorageVolumeWipeAlgLast It reflects the last algorithm supported by this version of the libvirt API
	StorageVolumeWipeAlgLast
)

// DomainState specifies state of the domain
type DomainState uint32

const (
	// DomainStateNoState No state
	DomainStateNoState = iota
	// DomainStateRunning The domain is running
	DomainStateRunning
	// DomainStateBlocked The domain is blocked on resource
	DomainStateBlocked
	// DomainStatePaused The domain is paused by user
	DomainStatePaused
	// DomainStateShutdown The domain is being shut down
	DomainStateShutdown
	// DomainStateShutoff The domain is shut off
	DomainStateShutoff
	// DomainStateCrashed The domain is crashed
	DomainStateCrashed
	// DomainStatePMSuspended The domain is suspended by guest power management
	DomainStatePMSuspended
	// DomainStateLast This value will increase over time as new events are added to the libvirt
	// API. It reflects the last state supported by this version of the libvirt API.
	DomainStateLast
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

// MessageHeader is a libvirt rpc packet header
type MessageHeader struct {
	// Program identifier
	Program uint32

	// Program version
	Version uint32

	// Remote procedure identifier
	Procedure RemoteProcedure

	// Call type, e.g., Reply
	Type MessageType

	// Call serial number
	Serial uint32

	// Request status, e.g., StatusOK
	Status MessageStatus
}

// Packet represents a RPC request or response.
type Packet struct {
	// Size of packet, in bytes, including length.
	// Length + Header + Payload
	Length uint32
	Header MessageHeader
}

type Event struct {
	CallbackID uint32
	Msg        interface{}
}

type Message struct {
	Header  MessageHeader
	Payload []byte
}

func NewMessage(hdr *MessageHeader, payload []byte) Message {
	return Message{Payload: payload, Header: *hdr}
}
