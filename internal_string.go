// Code generated by "stringer -type=RemoteProcedure,MessageStatus,MessageType -output internal_string.go"; DO NOT EDIT.

package libvirt

import "fmt"

const _RemoteProcedure_name = "RemoteProcConnectOpenRemoteProcConnectCloseRemoteProcConnectGetTypeRemoteProcConnectGetVersionRemoteProcConnectGetMaxVcpusRemoteProcNodeGetInfoRemoteProcConnectGetCapabilitiesRemoteProcDomainAttachDeviceRemoteProcDomainCreateRemoteProcDomainCreateXmlRemoteProcDomainDefineXmlRemoteProcDomainDestroyRemoteProcDomainDetachDeviceRemoteProcDomainGetXmlDescRemoteProcDomainGetAutostartRemoteProcDomainGetInfoRemoteProcDomainGetMaxMemoryRemoteProcDomainGetMaxVcpusRemoteProcDomainGetOsTypeRemoteProcDomainGetVcpusRemoteProcConnectListDefinedDomainsRemoteProcDomainLookupByIdRemoteProcDomainLookupByNameRemoteProcDomainLookupByUuidRemoteProcConnectNumOfDefinedDomainsRemoteProcDomainPinVcpuRemoteProcDomainRebootRemoteProcDomainResumeRemoteProcDomainSetAutostartRemoteProcDomainSetMaxMemoryRemoteProcDomainSetMemoryRemoteProcDomainSetVcpusRemoteProcDomainShutdownRemoteProcDomainSuspendRemoteProcDomainUndefineRemoteProcConnectListDefinedNetworksRemoteProcConnectListDomainsRemoteProcConnectListNetworksRemoteProcNetworkCreateRemoteProcNetworkCreateXmlRemoteProcNetworkDefineXmlRemoteProcNetworkDestroyRemoteProcNetworkGetXmlDescRemoteProcNetworkGetAutostartRemoteProcNetworkGetBridgeNameRemoteProcNetworkLookupByNameRemoteProcNetworkLookupByUuidRemoteProcNetworkSetAutostartRemoteProcNetworkUndefineRemoteProcConnectNumOfDefinedNetworksRemoteProcConnectNumOfDomainsRemoteProcConnectNumOfNetworksRemoteProcDomainCoreDumpRemoteProcDomainRestoreRemoteProcDomainSaveRemoteProcDomainGetSchedulerTypeRemoteProcDomainGetSchedulerParametersRemoteProcDomainSetSchedulerParametersRemoteProcConnectGetHostnameRemoteProcConnectSupportsFeatureRemoteProcDomainMigratePrepareRemoteProcDomainMigratePerformRemoteProcDomainMigrateFinishRemoteProcDomainBlockStatsRemoteProcDomainInterfaceStatsRemoteProcAuthListRemoteProcAuthSaslInitRemoteProcAuthSaslStartRemoteProcAuthSaslStepRemoteProcAuthPolkitRemoteProcConnectNumOfStoragePoolsRemoteProcConnectListStoragePoolsRemoteProcConnectNumOfDefinedStoragePoolsRemoteProcConnectListDefinedStoragePoolsRemoteProcConnectFindStoragePoolSourcesRemoteProcStoragePoolCreateXmlRemoteProcStoragePoolDefineXmlRemoteProcStoragePoolCreateRemoteProcStoragePoolBuildRemoteProcStoragePoolDestroyRemoteProcStoragePoolDeleteRemoteProcStoragePoolUndefineRemoteProcStoragePoolRefreshRemoteProcStoragePoolLookupByNameRemoteProcStoragePoolLookupByUuidRemoteProcStoragePoolLookupByVolumeRemoteProcStoragePoolGetInfoRemoteProcStoragePoolGetXmlDescRemoteProcStoragePoolGetAutostartRemoteProcStoragePoolSetAutostartRemoteProcStoragePoolNumOfVolumesRemoteProcStoragePoolListVolumesRemoteProcStorageVolCreateXmlRemoteProcStorageVolDeleteRemoteProcStorageVolLookupByNameRemoteProcStorageVolLookupByKeyRemoteProcStorageVolLookupByPathRemoteProcStorageVolGetInfoRemoteProcStorageVolGetXmlDescRemoteProcStorageVolGetPathRemoteProcNodeGetCellsFreeMemoryRemoteProcNodeGetFreeMemoryRemoteProcDomainBlockPeekRemoteProcDomainMemoryPeekRemoteProcConnectDomainEventRegisterRemoteProcConnectDomainEventDeregisterRemoteProcDomainEventLifecycleRemoteProcDomainMigratePrepare2RemoteProcDomainMigrateFinish2RemoteProcConnectGetUriRemoteProcNodeNumOfDevicesRemoteProcNodeListDevicesRemoteProcNodeDeviceLookupByNameRemoteProcNodeDeviceGetXmlDescRemoteProcNodeDeviceGetParentRemoteProcNodeDeviceNumOfCapsRemoteProcNodeDeviceListCapsRemoteProcNodeDeviceDettachRemoteProcNodeDeviceReAttachRemoteProcNodeDeviceResetRemoteProcDomainGetSecurityLabelRemoteProcNodeGetSecurityModelRemoteProcNodeDeviceCreateXmlRemoteProcNodeDeviceDestroyRemoteProcStorageVolCreateXmlFromRemoteProcConnectNumOfInterfacesRemoteProcConnectListInterfacesRemoteProcInterfaceLookupByNameRemoteProcInterfaceLookupByMacStringRemoteProcInterfaceGetXmlDescRemoteProcInterfaceDefineXmlRemoteProcInterfaceUndefineRemoteProcInterfaceCreateRemoteProcInterfaceDestroyRemoteProcConnectDomainXmlFromNativeRemoteProcConnectDomainXmlToNativeRemoteProcConnectNumOfDefinedInterfacesRemoteProcConnectListDefinedInterfacesRemoteProcConnectNumOfSecretsRemoteProcConnectListSecretsRemoteProcSecretLookupByUuidRemoteProcSecretDefineXmlRemoteProcSecretGetXmlDescRemoteProcSecretSetValueRemoteProcSecretGetValueRemoteProcSecretUndefineRemoteProcSecretLookupByUsageRemoteProcDomainMigratePrepareTunnelRemoteProcConnectIsSecureRemoteProcDomainIsActiveRemoteProcDomainIsPersistentRemoteProcNetworkIsActiveRemoteProcNetworkIsPersistentRemoteProcStoragePoolIsActiveRemoteProcStoragePoolIsPersistentRemoteProcInterfaceIsActiveRemoteProcConnectGetLibVersionRemoteProcConnectCompareCpuRemoteProcDomainMemoryStatsRemoteProcDomainAttachDeviceFlagsRemoteProcDomainDetachDeviceFlagsRemoteProcConnectBaselineCpuRemoteProcDomainGetJobInfoRemoteProcDomainAbortJobRemoteProcStorageVolWipeRemoteProcDomainMigrateSetMaxDowntimeRemoteProcConnectDomainEventRegisterAnyRemoteProcConnectDomainEventDeregisterAnyRemoteProcDomainEventRebootRemoteProcDomainEventRtcChangeRemoteProcDomainEventWatchdogRemoteProcDomainEventIoErrorRemoteProcDomainEventGraphicsRemoteProcDomainUpdateDeviceFlagsRemoteProcNwfilterLookupByNameRemoteProcNwfilterLookupByUuidRemoteProcNwfilterGetXmlDescRemoteProcConnectNumOfNwfiltersRemoteProcConnectListNwfiltersRemoteProcNwfilterDefineXmlRemoteProcNwfilterUndefineRemoteProcDomainManagedSaveRemoteProcDomainHasManagedSaveImageRemoteProcDomainManagedSaveRemoveRemoteProcDomainSnapshotCreateXmlRemoteProcDomainSnapshotGetXmlDescRemoteProcDomainSnapshotNumRemoteProcDomainSnapshotListNamesRemoteProcDomainSnapshotLookupByNameRemoteProcDomainHasCurrentSnapshotRemoteProcDomainSnapshotCurrentRemoteProcDomainRevertToSnapshotRemoteProcDomainSnapshotDeleteRemoteProcDomainGetBlockInfoRemoteProcDomainEventIoErrorReasonRemoteProcDomainCreateWithFlagsRemoteProcDomainSetMemoryParametersRemoteProcDomainGetMemoryParametersRemoteProcDomainSetVcpusFlagsRemoteProcDomainGetVcpusFlagsRemoteProcDomainOpenConsoleRemoteProcDomainIsUpdatedRemoteProcConnectGetSysinfoRemoteProcDomainSetMemoryFlagsRemoteProcDomainSetBlkioParametersRemoteProcDomainGetBlkioParametersRemoteProcDomainMigrateSetMaxSpeedRemoteProcStorageVolUploadRemoteProcStorageVolDownloadRemoteProcDomainInjectNmiRemoteProcDomainScreenshotRemoteProcDomainGetStateRemoteProcDomainMigrateBegin3RemoteProcDomainMigratePrepare3RemoteProcDomainMigratePrepareTunnel3RemoteProcDomainMigratePerform3RemoteProcDomainMigrateFinish3RemoteProcDomainMigrateConfirm3RemoteProcDomainSetSchedulerParametersFlagsRemoteProcInterfaceChangeBeginRemoteProcInterfaceChangeCommitRemoteProcInterfaceChangeRollbackRemoteProcDomainGetSchedulerParametersFlagsRemoteProcDomainEventControlErrorRemoteProcDomainPinVcpuFlagsRemoteProcDomainSendKeyRemoteProcNodeGetCpuStatsRemoteProcNodeGetMemoryStatsRemoteProcDomainGetControlInfoRemoteProcDomainGetVcpuPinInfoRemoteProcDomainUndefineFlagsRemoteProcDomainSaveFlagsRemoteProcDomainRestoreFlagsRemoteProcDomainDestroyFlagsRemoteProcDomainSaveImageGetXmlDescRemoteProcDomainSaveImageDefineXmlRemoteProcDomainBlockJobAbortRemoteProcDomainGetBlockJobInfoRemoteProcDomainBlockJobSetSpeedRemoteProcDomainBlockPullRemoteProcDomainEventBlockJobRemoteProcDomainMigrateGetMaxSpeedRemoteProcDomainBlockStatsFlagsRemoteProcDomainSnapshotGetParentRemoteProcDomainResetRemoteProcDomainSnapshotNumChildrenRemoteProcDomainSnapshotListChildrenNamesRemoteProcDomainEventDiskChangeRemoteProcDomainOpenGraphicsRemoteProcNodeSuspendForDurationRemoteProcDomainBlockResizeRemoteProcDomainSetBlockIoTuneRemoteProcDomainGetBlockIoTuneRemoteProcDomainSetNumaParametersRemoteProcDomainGetNumaParametersRemoteProcDomainSetInterfaceParametersRemoteProcDomainGetInterfaceParametersRemoteProcDomainShutdownFlagsRemoteProcStorageVolWipePatternRemoteProcStorageVolResizeRemoteProcDomainPmSuspendForDurationRemoteProcDomainGetCpuStatsRemoteProcDomainGetDiskErrorsRemoteProcDomainSetMetadataRemoteProcDomainGetMetadataRemoteProcDomainBlockRebaseRemoteProcDomainPmWakeupRemoteProcDomainEventTrayChangeRemoteProcDomainEventPmwakeupRemoteProcDomainEventPmsuspendRemoteProcDomainSnapshotIsCurrentRemoteProcDomainSnapshotHasMetadataRemoteProcConnectListAllDomainsRemoteProcDomainListAllSnapshotsRemoteProcDomainSnapshotListAllChildrenRemoteProcDomainEventBalloonChangeRemoteProcDomainGetHostnameRemoteProcDomainGetSecurityLabelListRemoteProcDomainPinEmulatorRemoteProcDomainGetEmulatorPinInfoRemoteProcConnectListAllStoragePoolsRemoteProcStoragePoolListAllVolumesRemoteProcConnectListAllNetworksRemoteProcConnectListAllInterfacesRemoteProcConnectListAllNodeDevicesRemoteProcConnectListAllNwfiltersRemoteProcConnectListAllSecretsRemoteProcNodeSetMemoryParametersRemoteProcNodeGetMemoryParametersRemoteProcDomainBlockCommitRemoteProcNetworkUpdateRemoteProcDomainEventPmsuspendDiskRemoteProcNodeGetCpuMapRemoteProcDomainFstrimRemoteProcDomainSendProcessSignalRemoteProcDomainOpenChannelRemoteProcNodeDeviceLookupScsiHostByWwnRemoteProcDomainGetJobStatsRemoteProcDomainMigrateGetCompressionCacheRemoteProcDomainMigrateSetCompressionCacheRemoteProcNodeDeviceDetachFlagsRemoteProcDomainMigrateBegin3ParamsRemoteProcDomainMigratePrepare3ParamsRemoteProcDomainMigratePrepareTunnel3ParamsRemoteProcDomainMigratePerform3ParamsRemoteProcDomainMigrateFinish3ParamsRemoteProcDomainMigrateConfirm3ParamsRemoteProcDomainSetMemoryStatsPeriodRemoteProcDomainCreateXmlWithFilesRemoteProcDomainCreateWithFilesRemoteProcDomainEventDeviceRemovedRemoteProcConnectGetCpuModelNamesRemoteProcConnectNetworkEventRegisterAnyRemoteProcConnectNetworkEventDeregisterAnyRemoteProcNetworkEventLifecycleRemoteProcConnectDomainEventCallbackRegisterAnyRemoteProcConnectDomainEventCallbackDeregisterAnyRemoteProcDomainEventCallbackLifecycleRemoteProcDomainEventCallbackRebootRemoteProcDomainEventCallbackRtcChangeRemoteProcDomainEventCallbackWatchdogRemoteProcDomainEventCallbackIoErrorRemoteProcDomainEventCallbackGraphicsRemoteProcDomainEventCallbackIoErrorReasonRemoteProcDomainEventCallbackControlErrorRemoteProcDomainEventCallbackBlockJobRemoteProcDomainEventCallbackDiskChangeRemoteProcDomainEventCallbackTrayChangeRemoteProcDomainEventCallbackPmwakeupRemoteProcDomainEventCallbackPmsuspendRemoteProcDomainEventCallbackBalloonChangeRemoteProcDomainEventCallbackPmsuspendDiskRemoteProcDomainEventCallbackDeviceRemovedRemoteProcDomainCoreDumpWithFormatRemoteProcDomainFsfreezeRemoteProcDomainFsthawRemoteProcDomainGetTimeRemoteProcDomainSetTimeRemoteProcDomainEventBlockJob2RemoteProcNodeGetFreePagesRemoteProcNetworkGetDhcpLeasesRemoteProcConnectGetDomainCapabilitiesRemoteProcDomainOpenGraphicsFdRemoteProcConnectGetAllDomainStatsRemoteProcDomainBlockCopyRemoteProcDomainEventCallbackTunableRemoteProcNodeAllocPagesRemoteProcDomainEventCallbackAgentLifecycleRemoteProcDomainGetFsinfoRemoteProcDomainDefineXmlFlagsRemoteProcDomainGetIothreadInfoRemoteProcDomainPinIothreadRemoteProcDomainInterfaceAddressesRemoteProcDomainEventCallbackDeviceAddedRemoteProcDomainAddIothreadRemoteProcDomainDelIothreadRemoteProcDomainSetUserPasswordRemoteProcDomainRenameRemoteProcDomainEventCallbackMigrationIterationRemoteProcConnectRegisterCloseCallbackRemoteProcConnectUnregisterCloseCallbackRemoteProcConnectEventConnectionClosedRemoteProcDomainEventCallbackJobCompletedRemoteProcDomainMigrateStartPostCopyRemoteProcDomainGetPerfEventsRemoteProcDomainSetPerfEventsRemoteProcDomainEventCallbackDeviceRemovalFailedRemoteProcConnectStoragePoolEventRegisterAnyRemoteProcConnectStoragePoolEventDeregisterAnyRemoteProcStoragePoolEventLifecycleRemoteProcDomainGetGuestVcpusRemoteProcDomainSetGuestVcpusRemoteProcStoragePoolEventRefreshRemoteProcConnectNodeDeviceEventRegisterAnyRemoteProcConnectNodeDeviceEventDeregisterAnyRemoteProcNodeDeviceEventLifecycleRemoteProcNodeDeviceEventUpdateRemoteProcStorageVolGetInfoFlagsRemoteProcDomainEventCallbackMetadataChangeRemoteProcConnectSecretEventRegisterAnyRemoteProcConnectSecretEventDeregisterAnyRemoteProcSecretEventLifecycleRemoteProcSecretEventValueChangedRemoteProcDomainSetVcpuRemoteProcDomainEventBlockThreshold"

var _RemoteProcedure_index = [...]uint16{0, 21, 43, 67, 94, 122, 143, 175, 203, 225, 250, 275, 298, 326, 352, 380, 403, 431, 458, 483, 507, 542, 568, 596, 624, 660, 683, 705, 727, 755, 783, 808, 832, 856, 879, 903, 939, 967, 996, 1019, 1045, 1071, 1095, 1122, 1151, 1181, 1210, 1239, 1268, 1293, 1330, 1359, 1389, 1413, 1436, 1456, 1488, 1526, 1564, 1592, 1624, 1654, 1684, 1713, 1739, 1769, 1787, 1809, 1832, 1854, 1874, 1908, 1941, 1982, 2022, 2061, 2091, 2121, 2148, 2174, 2202, 2229, 2258, 2286, 2319, 2352, 2387, 2415, 2446, 2479, 2512, 2545, 2577, 2606, 2632, 2664, 2695, 2727, 2754, 2784, 2811, 2843, 2870, 2895, 2921, 2957, 2995, 3025, 3056, 3086, 3109, 3135, 3160, 3192, 3222, 3251, 3280, 3308, 3335, 3363, 3388, 3420, 3450, 3479, 3506, 3539, 3571, 3602, 3633, 3669, 3698, 3726, 3753, 3778, 3804, 3840, 3874, 3913, 3951, 3980, 4008, 4036, 4061, 4087, 4111, 4135, 4159, 4188, 4224, 4249, 4273, 4301, 4326, 4355, 4384, 4417, 4444, 4474, 4501, 4528, 4561, 4594, 4622, 4648, 4672, 4696, 4733, 4772, 4813, 4840, 4870, 4899, 4927, 4956, 4989, 5019, 5049, 5077, 5108, 5138, 5165, 5191, 5218, 5253, 5286, 5319, 5353, 5380, 5413, 5449, 5483, 5514, 5546, 5576, 5604, 5638, 5669, 5704, 5739, 5768, 5797, 5824, 5849, 5876, 5906, 5940, 5974, 6008, 6034, 6062, 6087, 6113, 6137, 6166, 6197, 6234, 6265, 6295, 6326, 6369, 6399, 6430, 6463, 6506, 6539, 6567, 6590, 6615, 6643, 6673, 6703, 6732, 6757, 6785, 6813, 6848, 6882, 6911, 6942, 6974, 6999, 7028, 7062, 7093, 7126, 7147, 7182, 7223, 7254, 7282, 7314, 7341, 7371, 7401, 7434, 7467, 7505, 7543, 7572, 7603, 7629, 7665, 7692, 7721, 7748, 7775, 7802, 7826, 7857, 7886, 7916, 7949, 7984, 8015, 8047, 8086, 8120, 8147, 8183, 8210, 8244, 8280, 8315, 8347, 8381, 8416, 8449, 8480, 8513, 8546, 8573, 8596, 8630, 8653, 8675, 8708, 8735, 8774, 8801, 8843, 8885, 8916, 8951, 8988, 9031, 9068, 9104, 9141, 9177, 9211, 9242, 9276, 9309, 9349, 9391, 9422, 9469, 9518, 9556, 9591, 9629, 9666, 9702, 9739, 9781, 9822, 9859, 9898, 9937, 9974, 10012, 10054, 10096, 10138, 10172, 10196, 10218, 10241, 10264, 10294, 10320, 10350, 10388, 10418, 10452, 10477, 10513, 10537, 10580, 10605, 10635, 10666, 10693, 10727, 10767, 10794, 10821, 10852, 10874, 10921, 10959, 10999, 11037, 11078, 11114, 11143, 11172, 11220, 11264, 11310, 11345, 11374, 11403, 11436, 11479, 11524, 11558, 11589, 11621, 11664, 11703, 11744, 11774, 11807, 11830, 11865}

func (i RemoteProcedure) String() string {
	i -= 1
	if i >= RemoteProcedure(len(_RemoteProcedure_index)-1) {
		return fmt.Sprintf("RemoteProcedure(%d)", i+1)
	}
	return _RemoteProcedure_name[_RemoteProcedure_index[i]:_RemoteProcedure_index[i+1]]
}

const _MessageStatus_name = "MessageStatusOKMessageStatusErrorMessageStatusContinue"

var _MessageStatus_index = [...]uint8{0, 15, 33, 54}

func (i MessageStatus) String() string {
	if i >= MessageStatus(len(_MessageStatus_index)-1) {
		return fmt.Sprintf("MessageStatus(%d)", i)
	}
	return _MessageStatus_name[_MessageStatus_index[i]:_MessageStatus_index[i+1]]
}

const _MessageType_name = "MessageTypeCallMessageTypeReplyMessageTypeMessageMessageTypeStreamMessageTypeCallWithFDsMessageTypeReplyWithFDsMessageTypeStreamSkip"

var _MessageType_index = [...]uint8{0, 15, 31, 49, 66, 88, 111, 132}

func (i MessageType) String() string {
	if i >= MessageType(len(_MessageType_index)-1) {
		return fmt.Sprintf("MessageType(%d)", i)
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}