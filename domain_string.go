// Code generated by "stringer -type=EventIDTypes,DomainConsoleFlags,DomainStatsTypes,ListDomainsFlags,GetDomainsStatsFlags,DomainAffectFlags,DomainDeviceModifyFlags,DomainBlockResizeFlags,DomainXMLFlags,DomainCreateFlags,DomainRebootFlags,DomainShutdownFlags,DomainMigrateFlags,DomainUndefineFlags,DomainDefineXMLFlags,DomainDestroyFlags -output domain_string.go"; DO NOT EDIT.

package libvirt

import "fmt"

const _EventIDTypes_name = "DomainEventIDLifecycleDomainEventIDRebootDomainEventIDRtcChangeDomainEventIDWatchdogDomainEventIDIOErrorDomainEventIDGraphicsDomainEventIDIOErrorReasonDomainEventIDControlErrorDomainEventIDBlockJobDomainEventIDDiskChangeDomainEventIDTrayChangeDomainEventIDPmWakeupDomainEventIDPmSuspendDomainEventIDBalloonChangeDomainEventIDPmSuspendDiskDomainEventIDDeviceRemovedDomainEventIDBlockJob2DomainEventIDTunableDomainEventIDAgentLifecycleDomainEventIDDeviceAddedDomainEventIDMigrationIterationDomainEventIDJobCompletedDomainEventIDDeviceRemovalFailedDomainEventIDMetadataChangeDomainEventIDLast"

var _EventIDTypes_index = [...]uint16{0, 22, 41, 63, 84, 104, 125, 151, 176, 197, 220, 243, 264, 286, 312, 338, 364, 386, 406, 433, 457, 488, 513, 545, 572, 589}

func (i EventIDTypes) String() string {
	if i < 0 || i >= EventIDTypes(len(_EventIDTypes_index)-1) {
		return fmt.Sprintf("EventIDTypes(%d)", i)
	}
	return _EventIDTypes_name[_EventIDTypes_index[i]:_EventIDTypes_index[i+1]]
}

const _DomainConsoleFlags_name = "DomainConsoleFlagForceDomainConsoleFlagSafe"

var _DomainConsoleFlags_index = [...]uint8{0, 22, 43}

func (i DomainConsoleFlags) String() string {
	i -= 1
	if i >= DomainConsoleFlags(len(_DomainConsoleFlags_index)-1) {
		return fmt.Sprintf("DomainConsoleFlags(%d)", i+1)
	}
	return _DomainConsoleFlags_name[_DomainConsoleFlags_index[i]:_DomainConsoleFlags_index[i+1]]
}

const (
	_DomainStatsTypes_name_0 = "DomainStatsStateDomainStatsCpuTotal"
	_DomainStatsTypes_name_1 = "DomainStatsBalloon"
	_DomainStatsTypes_name_2 = "DomainStatsVcpu"
	_DomainStatsTypes_name_3 = "DomainStatsInterface"
	_DomainStatsTypes_name_4 = "DomainStatsBlock"
	_DomainStatsTypes_name_5 = "DomainStatsPerf"
)

var (
	_DomainStatsTypes_index_0 = [...]uint8{0, 16, 35}
	_DomainStatsTypes_index_1 = [...]uint8{0, 18}
	_DomainStatsTypes_index_2 = [...]uint8{0, 15}
	_DomainStatsTypes_index_3 = [...]uint8{0, 20}
	_DomainStatsTypes_index_4 = [...]uint8{0, 16}
	_DomainStatsTypes_index_5 = [...]uint8{0, 15}
)

func (i DomainStatsTypes) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DomainStatsTypes_name_0[_DomainStatsTypes_index_0[i]:_DomainStatsTypes_index_0[i+1]]
	case i == 4:
		return _DomainStatsTypes_name_1
	case i == 8:
		return _DomainStatsTypes_name_2
	case i == 16:
		return _DomainStatsTypes_name_3
	case i == 32:
		return _DomainStatsTypes_name_4
	case i == 64:
		return _DomainStatsTypes_name_5
	default:
		return fmt.Sprintf("DomainStatsTypes(%d)", i)
	}
}

const _ListDomainsFlags_name = "ListDomainsActiveListDomainsInactiveListDomainsPersistentListDomainsTransientListDomainsRunningListDomainsPausedListDomainsShutoffListDomainsOtherListDomainsManagedsaveListDomainsNoManagedsaveListDomainsAutostartListDomainsNoAutostartListDomainsHasSnapshotListDomainsNoSnapshot"

var _ListDomainsFlags_map = map[ListDomainsFlags]string{
	1:    _ListDomainsFlags_name[0:17],
	2:    _ListDomainsFlags_name[17:36],
	4:    _ListDomainsFlags_name[36:57],
	8:    _ListDomainsFlags_name[57:77],
	16:   _ListDomainsFlags_name[77:95],
	32:   _ListDomainsFlags_name[95:112],
	64:   _ListDomainsFlags_name[112:130],
	128:  _ListDomainsFlags_name[130:146],
	256:  _ListDomainsFlags_name[146:168],
	512:  _ListDomainsFlags_name[168:192],
	1024: _ListDomainsFlags_name[192:212],
	2048: _ListDomainsFlags_name[212:234],
	4096: _ListDomainsFlags_name[234:256],
	8192: _ListDomainsFlags_name[256:277],
}

func (i ListDomainsFlags) String() string {
	if str, ok := _ListDomainsFlags_map[i]; ok {
		return str
	}
	return fmt.Sprintf("ListDomainsFlags(%d)", i)
}

const (
	_GetDomainsStatsFlags_name_0 = "GetDomainsStatsActiveGetDomainsStatsInactive"
	_GetDomainsStatsFlags_name_1 = "ListDomainsStatsPersistent"
	_GetDomainsStatsFlags_name_2 = "GetDomainsStatsTransient"
	_GetDomainsStatsFlags_name_3 = "GetDomainsStatsRunning"
	_GetDomainsStatsFlags_name_4 = "ListDomainsStatsPaused"
	_GetDomainsStatsFlags_name_5 = "GetDomainsStatsShutoff"
	_GetDomainsStatsFlags_name_6 = "GetDomainsStatsOther"
	_GetDomainsStatsFlags_name_7 = "GetDomainsStatsBacking"
	_GetDomainsStatsFlags_name_8 = "GetDomainsStatsEnforce"
)

var (
	_GetDomainsStatsFlags_index_0 = [...]uint8{0, 21, 44}
	_GetDomainsStatsFlags_index_1 = [...]uint8{0, 26}
	_GetDomainsStatsFlags_index_2 = [...]uint8{0, 24}
	_GetDomainsStatsFlags_index_3 = [...]uint8{0, 22}
	_GetDomainsStatsFlags_index_4 = [...]uint8{0, 22}
	_GetDomainsStatsFlags_index_5 = [...]uint8{0, 22}
	_GetDomainsStatsFlags_index_6 = [...]uint8{0, 20}
	_GetDomainsStatsFlags_index_7 = [...]uint8{0, 22}
	_GetDomainsStatsFlags_index_8 = [...]uint8{0, 22}
)

func (i GetDomainsStatsFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _GetDomainsStatsFlags_name_0[_GetDomainsStatsFlags_index_0[i]:_GetDomainsStatsFlags_index_0[i+1]]
	case i == 4:
		return _GetDomainsStatsFlags_name_1
	case i == 8:
		return _GetDomainsStatsFlags_name_2
	case i == 16:
		return _GetDomainsStatsFlags_name_3
	case i == 32:
		return _GetDomainsStatsFlags_name_4
	case i == 64:
		return _GetDomainsStatsFlags_name_5
	case i == 128:
		return _GetDomainsStatsFlags_name_6
	case i == 1073741824:
		return _GetDomainsStatsFlags_name_7
	case i == 2147483648:
		return _GetDomainsStatsFlags_name_8
	default:
		return fmt.Sprintf("GetDomainsStatsFlags(%d)", i)
	}
}

const _DomainAffectFlags_name = "DomainAffectFlagCurrentDomainAffectFlagLiveDomainAffectFlagConfig"

var _DomainAffectFlags_index = [...]uint8{0, 23, 43, 65}

func (i DomainAffectFlags) String() string {
	if i >= DomainAffectFlags(len(_DomainAffectFlags_index)-1) {
		return fmt.Sprintf("DomainAffectFlags(%d)", i)
	}
	return _DomainAffectFlags_name[_DomainAffectFlags_index[i]:_DomainAffectFlags_index[i+1]]
}

const _DomainDeviceModifyFlags_name = "DomainDeviceModifyFlagConfigDomainDeviceModifyFlagCurrentDomainDeviceModifyFlagLiveDomainDeviceModifyFlagForce"

var _DomainDeviceModifyFlags_index = [...]uint8{0, 28, 57, 83, 110}

func (i DomainDeviceModifyFlags) String() string {
	if i >= DomainDeviceModifyFlags(len(_DomainDeviceModifyFlags_index)-1) {
		return fmt.Sprintf("DomainDeviceModifyFlags(%d)", i)
	}
	return _DomainDeviceModifyFlags_name[_DomainDeviceModifyFlags_index[i]:_DomainDeviceModifyFlags_index[i+1]]
}

const _DomainBlockResizeFlags_name = "DomainBlockResizeFlagNoneDomainBlockResizeFlagBytes"

var _DomainBlockResizeFlags_index = [...]uint8{0, 25, 51}

func (i DomainBlockResizeFlags) String() string {
	if i >= DomainBlockResizeFlags(len(_DomainBlockResizeFlags_index)-1) {
		return fmt.Sprintf("DomainBlockResizeFlags(%d)", i)
	}
	return _DomainBlockResizeFlags_name[_DomainBlockResizeFlags_index[i]:_DomainBlockResizeFlags_index[i+1]]
}

const (
	_DomainXMLFlags_name_0 = "DomainXMLFlagSecureDomainXMLFlagInactive"
	_DomainXMLFlags_name_1 = "DomainXMLFlagUpdateCPU"
	_DomainXMLFlags_name_2 = "DomainXMLFlagMigratable"
)

var (
	_DomainXMLFlags_index_0 = [...]uint8{0, 19, 40}
	_DomainXMLFlags_index_1 = [...]uint8{0, 22}
	_DomainXMLFlags_index_2 = [...]uint8{0, 23}
)

func (i DomainXMLFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DomainXMLFlags_name_0[_DomainXMLFlags_index_0[i]:_DomainXMLFlags_index_0[i+1]]
	case i == 4:
		return _DomainXMLFlags_name_1
	case i == 8:
		return _DomainXMLFlags_name_2
	default:
		return fmt.Sprintf("DomainXMLFlags(%d)", i)
	}
}

const (
	_DomainCreateFlags_name_0 = "DomainCreateFlagNoneDomainCreateFlagPausedDomainCreateFlagAutoDestroy"
	_DomainCreateFlags_name_1 = "DomainCreateFlagBypassCache"
	_DomainCreateFlags_name_2 = "DomainCreateFlagStartForceBoot"
	_DomainCreateFlags_name_3 = "DomainCreateFlagStartValidate"
)

var (
	_DomainCreateFlags_index_0 = [...]uint8{0, 20, 42, 69}
	_DomainCreateFlags_index_1 = [...]uint8{0, 27}
	_DomainCreateFlags_index_2 = [...]uint8{0, 30}
	_DomainCreateFlags_index_3 = [...]uint8{0, 29}
)

func (i DomainCreateFlags) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _DomainCreateFlags_name_0[_DomainCreateFlags_index_0[i]:_DomainCreateFlags_index_0[i+1]]
	case i == 4:
		return _DomainCreateFlags_name_1
	case i == 8:
		return _DomainCreateFlags_name_2
	case i == 16:
		return _DomainCreateFlags_name_3
	default:
		return fmt.Sprintf("DomainCreateFlags(%d)", i)
	}
}

const (
	_DomainRebootFlags_name_0 = "DomainRebootFlagDefaultDomainRebootFlagACPIDomainRebootFlagGuestAgent"
	_DomainRebootFlags_name_1 = "DomainRebootFlagInitctl"
	_DomainRebootFlags_name_2 = "DomainRebootFlagSignal"
	_DomainRebootFlags_name_3 = "DomainRebootFlagParavirt"
)

var (
	_DomainRebootFlags_index_0 = [...]uint8{0, 23, 43, 69}
	_DomainRebootFlags_index_1 = [...]uint8{0, 23}
	_DomainRebootFlags_index_2 = [...]uint8{0, 22}
	_DomainRebootFlags_index_3 = [...]uint8{0, 24}
)

func (i DomainRebootFlags) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _DomainRebootFlags_name_0[_DomainRebootFlags_index_0[i]:_DomainRebootFlags_index_0[i+1]]
	case i == 4:
		return _DomainRebootFlags_name_1
	case i == 8:
		return _DomainRebootFlags_name_2
	case i == 16:
		return _DomainRebootFlags_name_3
	default:
		return fmt.Sprintf("DomainRebootFlags(%d)", i)
	}
}

const (
	_DomainShutdownFlags_name_0 = "DomainShutdownFlagDefaultDomainShutdownFlagACPIDomainShutdownFlagGuestAgent"
	_DomainShutdownFlags_name_1 = "DomainShutdownFlagInitctl"
	_DomainShutdownFlags_name_2 = "DomainShutdownFlagSignal"
	_DomainShutdownFlags_name_3 = "DomainShutdownFlagParavirt"
)

var (
	_DomainShutdownFlags_index_0 = [...]uint8{0, 25, 47, 75}
	_DomainShutdownFlags_index_1 = [...]uint8{0, 25}
	_DomainShutdownFlags_index_2 = [...]uint8{0, 24}
	_DomainShutdownFlags_index_3 = [...]uint8{0, 26}
)

func (i DomainShutdownFlags) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _DomainShutdownFlags_name_0[_DomainShutdownFlags_index_0[i]:_DomainShutdownFlags_index_0[i+1]]
	case i == 4:
		return _DomainShutdownFlags_name_1
	case i == 8:
		return _DomainShutdownFlags_name_2
	case i == 16:
		return _DomainShutdownFlags_name_3
	default:
		return fmt.Sprintf("DomainShutdownFlags(%d)", i)
	}
}

const _DomainMigrateFlags_name = "DomainMigrateFlagLiveDomainMigrateFlagPeerToPeerDomainMigrateFlagTunneledDomainMigrateFlagPersistDestinationDomainMigrateFlagUndefineSourceDomainMigrateFlagPausedDomainMigrateFlagNonSharedDiskDomainMigrateFlagNonSharedIncrementalDomainMigrateFlagChangeProtectionDomainMigrateFlagUnsafeDomainMigrateFlagOfflineDomainMigrateFlagCompressedDomainMigrateFlagAbortOnErrorDomainMigrateFlagAutoConvergeDomainMigrateFlagRDMAPinAllDomainMigrateFlagPostcopy"

var _DomainMigrateFlags_map = map[DomainMigrateFlags]string{
	1:     _DomainMigrateFlags_name[0:21],
	2:     _DomainMigrateFlags_name[21:48],
	4:     _DomainMigrateFlags_name[48:73],
	8:     _DomainMigrateFlags_name[73:108],
	16:    _DomainMigrateFlags_name[108:139],
	32:    _DomainMigrateFlags_name[139:162],
	64:    _DomainMigrateFlags_name[162:192],
	128:   _DomainMigrateFlags_name[192:229],
	256:   _DomainMigrateFlags_name[229:262],
	512:   _DomainMigrateFlags_name[262:285],
	1024:  _DomainMigrateFlags_name[285:309],
	2048:  _DomainMigrateFlags_name[309:336],
	4096:  _DomainMigrateFlags_name[336:365],
	8192:  _DomainMigrateFlags_name[365:394],
	16384: _DomainMigrateFlags_name[394:421],
	32768: _DomainMigrateFlags_name[421:446],
}

func (i DomainMigrateFlags) String() string {
	if str, ok := _DomainMigrateFlags_map[i]; ok {
		return str
	}
	return fmt.Sprintf("DomainMigrateFlags(%d)", i)
}

const (
	_DomainUndefineFlags_name_0 = "DomainUndefineFlagManagedSaveDomainUndefineFlagSnapshotsMetadata"
	_DomainUndefineFlags_name_1 = "DomainUndefineFlagNVRAM"
)

var (
	_DomainUndefineFlags_index_0 = [...]uint8{0, 29, 64}
	_DomainUndefineFlags_index_1 = [...]uint8{0, 23}
)

func (i DomainUndefineFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DomainUndefineFlags_name_0[_DomainUndefineFlags_index_0[i]:_DomainUndefineFlags_index_0[i+1]]
	case i == 4:
		return _DomainUndefineFlags_name_1
	default:
		return fmt.Sprintf("DomainUndefineFlags(%d)", i)
	}
}

const _DomainDefineXMLFlags_name = "DefineValidate"

var _DomainDefineXMLFlags_index = [...]uint8{0, 14}

func (i DomainDefineXMLFlags) String() string {
	i -= 1
	if i >= DomainDefineXMLFlags(len(_DomainDefineXMLFlags_index)-1) {
		return fmt.Sprintf("DomainDefineXMLFlags(%d)", i+1)
	}
	return _DomainDefineXMLFlags_name[_DomainDefineXMLFlags_index[i]:_DomainDefineXMLFlags_index[i+1]]
}

const _DomainDestroyFlags_name = "DestroyFlagDefaultDestroyFlagGraceful"

var _DomainDestroyFlags_index = [...]uint8{0, 18, 37}

func (i DomainDestroyFlags) String() string {
	i -= 1
	if i >= DomainDestroyFlags(len(_DomainDestroyFlags_index)-1) {
		return fmt.Sprintf("DomainDestroyFlags(%d)", i+1)
	}
	return _DomainDestroyFlags_name[_DomainDestroyFlags_index[i]:_DomainDestroyFlags_index[i+1]]
}
