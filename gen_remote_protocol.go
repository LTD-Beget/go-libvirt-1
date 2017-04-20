package libvirt

const (
	RemoteStringMax                      = 4194304
	RemoteDomainListMax                  = 16384
	RemoteCpumapMax                      = 2048
	RemoteVcpuinfoMax                    = 16384
	RemoteCpumapsMax                     = 8388608
	RemoteIothreadInfoMax                = 16384
	RemoteMigrateCookieMax               = 4194304
	RemoteNetworkListMax                 = 16384
	RemoteInterfaceListMax               = 16384
	RemoteStoragePoolListMax             = 4096
	RemoteStorageVolListMax              = 16384
	RemoteNodeDeviceListMax              = 65536
	RemoteNodeDeviceCapsListMax          = 65536
	RemoteNwfilterListMax                = 1024
	RemoteDomainSchedulerParametersMax   = 16
	RemoteDomainBlkioParametersMax       = 16
	RemoteDomainMemoryParametersMax      = 16
	RemoteDomainBlockIoTuneParametersMax = 32
	RemoteDomainNumaParametersMax        = 16
	RemoteDomainPerfEventsMax            = 64
	RemoteDomainBlockCopyParametersMax   = 16
	RemoteNodeCpuStatsMax                = 16
	RemoteNodeMemoryStatsMax             = 16
	RemoteDomainBlockStatsParametersMax  = 16
	RemoteNodeMaxCells                   = 1024
	RemoteAuthSaslDataMax                = 65536
	RemoteAuthTypeListMax                = 20
	RemoteDomainMemoryStatsMax           = 1024
	RemoteDomainSnapshotListMax          = 1024
	RemoteDomainBlockPeekBufferMax       = 4194304
	RemoteDomainMemoryPeekBufferMax      = 4194304
	RemoteSecurityLabelListMax           = 64
	RemoteSecurityModelMax               = SecurityModelBuflen
	RemoteSecurityLabelMax               = SecurityLabelBuflen
	RemoteSecurityDoiMax                 = SecurityDoiBuflen
	RemoteSecretValueMax                 = 65536
	RemoteSecretListMax                  = 16384
	RemoteCpuBaselineMax                 = 256
	RemoteDomainSendKeyMax               = 16
	RemoteDomainInterfaceParametersMax   = 16
	RemoteDomainGetCpuStatsNcpusMax      = 128
	RemoteDomainGetCpuStatsMax           = 2048
	RemoteDomainDiskErrorsMax            = 256
	RemoteNodeMemoryParametersMax        = 64
	RemoteDomainMigrateParamListMax      = 64
	RemoteDomainJobStatsMax              = 64
	RemoteConnectCpuModelsMax            = 8192
	RemoteDomainFsfreezeMountpointsMax   = 256
	RemoteNetworkDhcpLeasesMax           = 65536
	RemoteConnectGetAllDomainStatsMax    = 4096
	RemoteDomainEventTunableMax          = 2048
	RemoteDomainFsinfoMax                = 256
	RemoteDomainFsinfoDisksMax           = 256
	RemoteDomainInterfaceMax             = 2048
	RemoteDomainIpAddrMax                = 2048
	RemoteDomainGuestVcpuParamsMax       = 64
	RemoteDomainEventGraphicsIdentityMax = 20
	RemoteProgram                        = 0x20008086
	RemoteProtocolVersion                = 1
)

const (
	RemoteProcConnectOpen                             RemoteProcedure = 1
	RemoteProcConnectClose                            RemoteProcedure = 2
	RemoteProcConnectGetType                          RemoteProcedure = 3
	RemoteProcConnectGetVersion                       RemoteProcedure = 4
	RemoteProcConnectGetMaxVcpus                      RemoteProcedure = 5
	RemoteProcNodeGetInfo                             RemoteProcedure = 6
	RemoteProcConnectGetCapabilities                  RemoteProcedure = 7
	RemoteProcDomainAttachDevice                      RemoteProcedure = 8
	RemoteProcDomainCreate                            RemoteProcedure = 9
	RemoteProcDomainCreateXml                         RemoteProcedure = 10
	RemoteProcDomainDefineXml                         RemoteProcedure = 11
	RemoteProcDomainDestroy                           RemoteProcedure = 12
	RemoteProcDomainDetachDevice                      RemoteProcedure = 13
	RemoteProcDomainGetXmlDesc                        RemoteProcedure = 14
	RemoteProcDomainGetAutostart                      RemoteProcedure = 15
	RemoteProcDomainGetInfo                           RemoteProcedure = 16
	RemoteProcDomainGetMaxMemory                      RemoteProcedure = 17
	RemoteProcDomainGetMaxVcpus                       RemoteProcedure = 18
	RemoteProcDomainGetOsType                         RemoteProcedure = 19
	RemoteProcDomainGetVcpus                          RemoteProcedure = 20
	RemoteProcConnectListDefinedDomains               RemoteProcedure = 21
	RemoteProcDomainLookupById                        RemoteProcedure = 22
	RemoteProcDomainLookupByName                      RemoteProcedure = 23
	RemoteProcDomainLookupByUuid                      RemoteProcedure = 24
	RemoteProcConnectNumOfDefinedDomains              RemoteProcedure = 25
	RemoteProcDomainPinVcpu                           RemoteProcedure = 26
	RemoteProcDomainReboot                            RemoteProcedure = 27
	RemoteProcDomainResume                            RemoteProcedure = 28
	RemoteProcDomainSetAutostart                      RemoteProcedure = 29
	RemoteProcDomainSetMaxMemory                      RemoteProcedure = 30
	RemoteProcDomainSetMemory                         RemoteProcedure = 31
	RemoteProcDomainSetVcpus                          RemoteProcedure = 32
	RemoteProcDomainShutdown                          RemoteProcedure = 33
	RemoteProcDomainSuspend                           RemoteProcedure = 34
	RemoteProcDomainUndefine                          RemoteProcedure = 35
	RemoteProcConnectListDefinedNetworks              RemoteProcedure = 36
	RemoteProcConnectListDomains                      RemoteProcedure = 37
	RemoteProcConnectListNetworks                     RemoteProcedure = 38
	RemoteProcNetworkCreate                           RemoteProcedure = 39
	RemoteProcNetworkCreateXml                        RemoteProcedure = 40
	RemoteProcNetworkDefineXml                        RemoteProcedure = 41
	RemoteProcNetworkDestroy                          RemoteProcedure = 42
	RemoteProcNetworkGetXmlDesc                       RemoteProcedure = 43
	RemoteProcNetworkGetAutostart                     RemoteProcedure = 44
	RemoteProcNetworkGetBridgeName                    RemoteProcedure = 45
	RemoteProcNetworkLookupByName                     RemoteProcedure = 46
	RemoteProcNetworkLookupByUuid                     RemoteProcedure = 47
	RemoteProcNetworkSetAutostart                     RemoteProcedure = 48
	RemoteProcNetworkUndefine                         RemoteProcedure = 49
	RemoteProcConnectNumOfDefinedNetworks             RemoteProcedure = 50
	RemoteProcConnectNumOfDomains                     RemoteProcedure = 51
	RemoteProcConnectNumOfNetworks                    RemoteProcedure = 52
	RemoteProcDomainCoreDump                          RemoteProcedure = 53
	RemoteProcDomainRestore                           RemoteProcedure = 54
	RemoteProcDomainSave                              RemoteProcedure = 55
	RemoteProcDomainGetSchedulerType                  RemoteProcedure = 56
	RemoteProcDomainGetSchedulerParameters            RemoteProcedure = 57
	RemoteProcDomainSetSchedulerParameters            RemoteProcedure = 58
	RemoteProcConnectGetHostname                      RemoteProcedure = 59
	RemoteProcConnectSupportsFeature                  RemoteProcedure = 60
	RemoteProcDomainMigratePrepare                    RemoteProcedure = 61
	RemoteProcDomainMigratePerform                    RemoteProcedure = 62
	RemoteProcDomainMigrateFinish                     RemoteProcedure = 63
	RemoteProcDomainBlockStats                        RemoteProcedure = 64
	RemoteProcDomainInterfaceStats                    RemoteProcedure = 65
	RemoteProcAuthList                                RemoteProcedure = 66
	RemoteProcAuthSaslInit                            RemoteProcedure = 67
	RemoteProcAuthSaslStart                           RemoteProcedure = 68
	RemoteProcAuthSaslStep                            RemoteProcedure = 69
	RemoteProcAuthPolkit                              RemoteProcedure = 70
	RemoteProcConnectNumOfStoragePools                RemoteProcedure = 71
	RemoteProcConnectListStoragePools                 RemoteProcedure = 72
	RemoteProcConnectNumOfDefinedStoragePools         RemoteProcedure = 73
	RemoteProcConnectListDefinedStoragePools          RemoteProcedure = 74
	RemoteProcConnectFindStoragePoolSources           RemoteProcedure = 75
	RemoteProcStoragePoolCreateXml                    RemoteProcedure = 76
	RemoteProcStoragePoolDefineXml                    RemoteProcedure = 77
	RemoteProcStoragePoolCreate                       RemoteProcedure = 78
	RemoteProcStoragePoolBuild                        RemoteProcedure = 79
	RemoteProcStoragePoolDestroy                      RemoteProcedure = 80
	RemoteProcStoragePoolDelete                       RemoteProcedure = 81
	RemoteProcStoragePoolUndefine                     RemoteProcedure = 82
	RemoteProcStoragePoolRefresh                      RemoteProcedure = 83
	RemoteProcStoragePoolLookupByName                 RemoteProcedure = 84
	RemoteProcStoragePoolLookupByUuid                 RemoteProcedure = 85
	RemoteProcStoragePoolLookupByVolume               RemoteProcedure = 86
	RemoteProcStoragePoolGetInfo                      RemoteProcedure = 87
	RemoteProcStoragePoolGetXmlDesc                   RemoteProcedure = 88
	RemoteProcStoragePoolGetAutostart                 RemoteProcedure = 89
	RemoteProcStoragePoolSetAutostart                 RemoteProcedure = 90
	RemoteProcStoragePoolNumOfVolumes                 RemoteProcedure = 91
	RemoteProcStoragePoolListVolumes                  RemoteProcedure = 92
	RemoteProcStorageVolCreateXml                     RemoteProcedure = 93
	RemoteProcStorageVolDelete                        RemoteProcedure = 94
	RemoteProcStorageVolLookupByName                  RemoteProcedure = 95
	RemoteProcStorageVolLookupByKey                   RemoteProcedure = 96
	RemoteProcStorageVolLookupByPath                  RemoteProcedure = 97
	RemoteProcStorageVolGetInfo                       RemoteProcedure = 98
	RemoteProcStorageVolGetXmlDesc                    RemoteProcedure = 99
	RemoteProcStorageVolGetPath                       RemoteProcedure = 100
	RemoteProcNodeGetCellsFreeMemory                  RemoteProcedure = 101
	RemoteProcNodeGetFreeMemory                       RemoteProcedure = 102
	RemoteProcDomainBlockPeek                         RemoteProcedure = 103
	RemoteProcDomainMemoryPeek                        RemoteProcedure = 104
	RemoteProcConnectDomainEventRegister              RemoteProcedure = 105
	RemoteProcConnectDomainEventDeregister            RemoteProcedure = 106
	RemoteProcDomainEventLifecycle                    RemoteProcedure = 107
	RemoteProcDomainMigratePrepare2                   RemoteProcedure = 108
	RemoteProcDomainMigrateFinish2                    RemoteProcedure = 109
	RemoteProcConnectGetUri                           RemoteProcedure = 110
	RemoteProcNodeNumOfDevices                        RemoteProcedure = 111
	RemoteProcNodeListDevices                         RemoteProcedure = 112
	RemoteProcNodeDeviceLookupByName                  RemoteProcedure = 113
	RemoteProcNodeDeviceGetXmlDesc                    RemoteProcedure = 114
	RemoteProcNodeDeviceGetParent                     RemoteProcedure = 115
	RemoteProcNodeDeviceNumOfCaps                     RemoteProcedure = 116
	RemoteProcNodeDeviceListCaps                      RemoteProcedure = 117
	RemoteProcNodeDeviceDettach                       RemoteProcedure = 118
	RemoteProcNodeDeviceReAttach                      RemoteProcedure = 119
	RemoteProcNodeDeviceReset                         RemoteProcedure = 120
	RemoteProcDomainGetSecurityLabel                  RemoteProcedure = 121
	RemoteProcNodeGetSecurityModel                    RemoteProcedure = 122
	RemoteProcNodeDeviceCreateXml                     RemoteProcedure = 123
	RemoteProcNodeDeviceDestroy                       RemoteProcedure = 124
	RemoteProcStorageVolCreateXmlFrom                 RemoteProcedure = 125
	RemoteProcConnectNumOfInterfaces                  RemoteProcedure = 126
	RemoteProcConnectListInterfaces                   RemoteProcedure = 127
	RemoteProcInterfaceLookupByName                   RemoteProcedure = 128
	RemoteProcInterfaceLookupByMacString              RemoteProcedure = 129
	RemoteProcInterfaceGetXmlDesc                     RemoteProcedure = 130
	RemoteProcInterfaceDefineXml                      RemoteProcedure = 131
	RemoteProcInterfaceUndefine                       RemoteProcedure = 132
	RemoteProcInterfaceCreate                         RemoteProcedure = 133
	RemoteProcInterfaceDestroy                        RemoteProcedure = 134
	RemoteProcConnectDomainXmlFromNative              RemoteProcedure = 135
	RemoteProcConnectDomainXmlToNative                RemoteProcedure = 136
	RemoteProcConnectNumOfDefinedInterfaces           RemoteProcedure = 137
	RemoteProcConnectListDefinedInterfaces            RemoteProcedure = 138
	RemoteProcConnectNumOfSecrets                     RemoteProcedure = 139
	RemoteProcConnectListSecrets                      RemoteProcedure = 140
	RemoteProcSecretLookupByUuid                      RemoteProcedure = 141
	RemoteProcSecretDefineXml                         RemoteProcedure = 142
	RemoteProcSecretGetXmlDesc                        RemoteProcedure = 143
	RemoteProcSecretSetValue                          RemoteProcedure = 144
	RemoteProcSecretGetValue                          RemoteProcedure = 145
	RemoteProcSecretUndefine                          RemoteProcedure = 146
	RemoteProcSecretLookupByUsage                     RemoteProcedure = 147
	RemoteProcDomainMigratePrepareTunnel              RemoteProcedure = 148
	RemoteProcConnectIsSecure                         RemoteProcedure = 149
	RemoteProcDomainIsActive                          RemoteProcedure = 150
	RemoteProcDomainIsPersistent                      RemoteProcedure = 151
	RemoteProcNetworkIsActive                         RemoteProcedure = 152
	RemoteProcNetworkIsPersistent                     RemoteProcedure = 153
	RemoteProcStoragePoolIsActive                     RemoteProcedure = 154
	RemoteProcStoragePoolIsPersistent                 RemoteProcedure = 155
	RemoteProcInterfaceIsActive                       RemoteProcedure = 156
	RemoteProcConnectGetLibVersion                    RemoteProcedure = 157
	RemoteProcConnectCompareCpu                       RemoteProcedure = 158
	RemoteProcDomainMemoryStats                       RemoteProcedure = 159
	RemoteProcDomainAttachDeviceFlags                 RemoteProcedure = 160
	RemoteProcDomainDetachDeviceFlags                 RemoteProcedure = 161
	RemoteProcConnectBaselineCpu                      RemoteProcedure = 162
	RemoteProcDomainGetJobInfo                        RemoteProcedure = 163
	RemoteProcDomainAbortJob                          RemoteProcedure = 164
	RemoteProcStorageVolWipe                          RemoteProcedure = 165
	RemoteProcDomainMigrateSetMaxDowntime             RemoteProcedure = 166
	RemoteProcConnectDomainEventRegisterAny           RemoteProcedure = 167
	RemoteProcConnectDomainEventDeregisterAny         RemoteProcedure = 168
	RemoteProcDomainEventReboot                       RemoteProcedure = 169
	RemoteProcDomainEventRtcChange                    RemoteProcedure = 170
	RemoteProcDomainEventWatchdog                     RemoteProcedure = 171
	RemoteProcDomainEventIoError                      RemoteProcedure = 172
	RemoteProcDomainEventGraphics                     RemoteProcedure = 173
	RemoteProcDomainUpdateDeviceFlags                 RemoteProcedure = 174
	RemoteProcNwfilterLookupByName                    RemoteProcedure = 175
	RemoteProcNwfilterLookupByUuid                    RemoteProcedure = 176
	RemoteProcNwfilterGetXmlDesc                      RemoteProcedure = 177
	RemoteProcConnectNumOfNwfilters                   RemoteProcedure = 178
	RemoteProcConnectListNwfilters                    RemoteProcedure = 179
	RemoteProcNwfilterDefineXml                       RemoteProcedure = 180
	RemoteProcNwfilterUndefine                        RemoteProcedure = 181
	RemoteProcDomainManagedSave                       RemoteProcedure = 182
	RemoteProcDomainHasManagedSaveImage               RemoteProcedure = 183
	RemoteProcDomainManagedSaveRemove                 RemoteProcedure = 184
	RemoteProcDomainSnapshotCreateXml                 RemoteProcedure = 185
	RemoteProcDomainSnapshotGetXmlDesc                RemoteProcedure = 186
	RemoteProcDomainSnapshotNum                       RemoteProcedure = 187
	RemoteProcDomainSnapshotListNames                 RemoteProcedure = 188
	RemoteProcDomainSnapshotLookupByName              RemoteProcedure = 189
	RemoteProcDomainHasCurrentSnapshot                RemoteProcedure = 190
	RemoteProcDomainSnapshotCurrent                   RemoteProcedure = 191
	RemoteProcDomainRevertToSnapshot                  RemoteProcedure = 192
	RemoteProcDomainSnapshotDelete                    RemoteProcedure = 193
	RemoteProcDomainGetBlockInfo                      RemoteProcedure = 194
	RemoteProcDomainEventIoErrorReason                RemoteProcedure = 195
	RemoteProcDomainCreateWithFlags                   RemoteProcedure = 196
	RemoteProcDomainSetMemoryParameters               RemoteProcedure = 197
	RemoteProcDomainGetMemoryParameters               RemoteProcedure = 198
	RemoteProcDomainSetVcpusFlags                     RemoteProcedure = 199
	RemoteProcDomainGetVcpusFlags                     RemoteProcedure = 200
	RemoteProcDomainOpenConsole                       RemoteProcedure = 201
	RemoteProcDomainIsUpdated                         RemoteProcedure = 202
	RemoteProcConnectGetSysinfo                       RemoteProcedure = 203
	RemoteProcDomainSetMemoryFlags                    RemoteProcedure = 204
	RemoteProcDomainSetBlkioParameters                RemoteProcedure = 205
	RemoteProcDomainGetBlkioParameters                RemoteProcedure = 206
	RemoteProcDomainMigrateSetMaxSpeed                RemoteProcedure = 207
	RemoteProcStorageVolUpload                        RemoteProcedure = 208
	RemoteProcStorageVolDownload                      RemoteProcedure = 209
	RemoteProcDomainInjectNmi                         RemoteProcedure = 210
	RemoteProcDomainScreenshot                        RemoteProcedure = 211
	RemoteProcDomainGetState                          RemoteProcedure = 212
	RemoteProcDomainMigrateBegin3                     RemoteProcedure = 213
	RemoteProcDomainMigratePrepare3                   RemoteProcedure = 214
	RemoteProcDomainMigratePrepareTunnel3             RemoteProcedure = 215
	RemoteProcDomainMigratePerform3                   RemoteProcedure = 216
	RemoteProcDomainMigrateFinish3                    RemoteProcedure = 217
	RemoteProcDomainMigrateConfirm3                   RemoteProcedure = 218
	RemoteProcDomainSetSchedulerParametersFlags       RemoteProcedure = 219
	RemoteProcInterfaceChangeBegin                    RemoteProcedure = 220
	RemoteProcInterfaceChangeCommit                   RemoteProcedure = 221
	RemoteProcInterfaceChangeRollback                 RemoteProcedure = 222
	RemoteProcDomainGetSchedulerParametersFlags       RemoteProcedure = 223
	RemoteProcDomainEventControlError                 RemoteProcedure = 224
	RemoteProcDomainPinVcpuFlags                      RemoteProcedure = 225
	RemoteProcDomainSendKey                           RemoteProcedure = 226
	RemoteProcNodeGetCpuStats                         RemoteProcedure = 227
	RemoteProcNodeGetMemoryStats                      RemoteProcedure = 228
	RemoteProcDomainGetControlInfo                    RemoteProcedure = 229
	RemoteProcDomainGetVcpuPinInfo                    RemoteProcedure = 230
	RemoteProcDomainUndefineFlags                     RemoteProcedure = 231
	RemoteProcDomainSaveFlags                         RemoteProcedure = 232
	RemoteProcDomainRestoreFlags                      RemoteProcedure = 233
	RemoteProcDomainDestroyFlags                      RemoteProcedure = 234
	RemoteProcDomainSaveImageGetXmlDesc               RemoteProcedure = 235
	RemoteProcDomainSaveImageDefineXml                RemoteProcedure = 236
	RemoteProcDomainBlockJobAbort                     RemoteProcedure = 237
	RemoteProcDomainGetBlockJobInfo                   RemoteProcedure = 238
	RemoteProcDomainBlockJobSetSpeed                  RemoteProcedure = 239
	RemoteProcDomainBlockPull                         RemoteProcedure = 240
	RemoteProcDomainEventBlockJob                     RemoteProcedure = 241
	RemoteProcDomainMigrateGetMaxSpeed                RemoteProcedure = 242
	RemoteProcDomainBlockStatsFlags                   RemoteProcedure = 243
	RemoteProcDomainSnapshotGetParent                 RemoteProcedure = 244
	RemoteProcDomainReset                             RemoteProcedure = 245
	RemoteProcDomainSnapshotNumChildren               RemoteProcedure = 246
	RemoteProcDomainSnapshotListChildrenNames         RemoteProcedure = 247
	RemoteProcDomainEventDiskChange                   RemoteProcedure = 248
	RemoteProcDomainOpenGraphics                      RemoteProcedure = 249
	RemoteProcNodeSuspendForDuration                  RemoteProcedure = 250
	RemoteProcDomainBlockResize                       RemoteProcedure = 251
	RemoteProcDomainSetBlockIoTune                    RemoteProcedure = 252
	RemoteProcDomainGetBlockIoTune                    RemoteProcedure = 253
	RemoteProcDomainSetNumaParameters                 RemoteProcedure = 254
	RemoteProcDomainGetNumaParameters                 RemoteProcedure = 255
	RemoteProcDomainSetInterfaceParameters            RemoteProcedure = 256
	RemoteProcDomainGetInterfaceParameters            RemoteProcedure = 257
	RemoteProcDomainShutdownFlags                     RemoteProcedure = 258
	RemoteProcStorageVolWipePattern                   RemoteProcedure = 259
	RemoteProcStorageVolResize                        RemoteProcedure = 260
	RemoteProcDomainPmSuspendForDuration              RemoteProcedure = 261
	RemoteProcDomainGetCpuStats                       RemoteProcedure = 262
	RemoteProcDomainGetDiskErrors                     RemoteProcedure = 263
	RemoteProcDomainSetMetadata                       RemoteProcedure = 264
	RemoteProcDomainGetMetadata                       RemoteProcedure = 265
	RemoteProcDomainBlockRebase                       RemoteProcedure = 266
	RemoteProcDomainPmWakeup                          RemoteProcedure = 267
	RemoteProcDomainEventTrayChange                   RemoteProcedure = 268
	RemoteProcDomainEventPmwakeup                     RemoteProcedure = 269
	RemoteProcDomainEventPmsuspend                    RemoteProcedure = 270
	RemoteProcDomainSnapshotIsCurrent                 RemoteProcedure = 271
	RemoteProcDomainSnapshotHasMetadata               RemoteProcedure = 272
	RemoteProcConnectListAllDomains                   RemoteProcedure = 273
	RemoteProcDomainListAllSnapshots                  RemoteProcedure = 274
	RemoteProcDomainSnapshotListAllChildren           RemoteProcedure = 275
	RemoteProcDomainEventBalloonChange                RemoteProcedure = 276
	RemoteProcDomainGetHostname                       RemoteProcedure = 277
	RemoteProcDomainGetSecurityLabelList              RemoteProcedure = 278
	RemoteProcDomainPinEmulator                       RemoteProcedure = 279
	RemoteProcDomainGetEmulatorPinInfo                RemoteProcedure = 280
	RemoteProcConnectListAllStoragePools              RemoteProcedure = 281
	RemoteProcStoragePoolListAllVolumes               RemoteProcedure = 282
	RemoteProcConnectListAllNetworks                  RemoteProcedure = 283
	RemoteProcConnectListAllInterfaces                RemoteProcedure = 284
	RemoteProcConnectListAllNodeDevices               RemoteProcedure = 285
	RemoteProcConnectListAllNwfilters                 RemoteProcedure = 286
	RemoteProcConnectListAllSecrets                   RemoteProcedure = 287
	RemoteProcNodeSetMemoryParameters                 RemoteProcedure = 288
	RemoteProcNodeGetMemoryParameters                 RemoteProcedure = 289
	RemoteProcDomainBlockCommit                       RemoteProcedure = 290
	RemoteProcNetworkUpdate                           RemoteProcedure = 291
	RemoteProcDomainEventPmsuspendDisk                RemoteProcedure = 292
	RemoteProcNodeGetCpuMap                           RemoteProcedure = 293
	RemoteProcDomainFstrim                            RemoteProcedure = 294
	RemoteProcDomainSendProcessSignal                 RemoteProcedure = 295
	RemoteProcDomainOpenChannel                       RemoteProcedure = 296
	RemoteProcNodeDeviceLookupScsiHostByWwn           RemoteProcedure = 297
	RemoteProcDomainGetJobStats                       RemoteProcedure = 298
	RemoteProcDomainMigrateGetCompressionCache        RemoteProcedure = 299
	RemoteProcDomainMigrateSetCompressionCache        RemoteProcedure = 300
	RemoteProcNodeDeviceDetachFlags                   RemoteProcedure = 301
	RemoteProcDomainMigrateBegin3Params               RemoteProcedure = 302
	RemoteProcDomainMigratePrepare3Params             RemoteProcedure = 303
	RemoteProcDomainMigratePrepareTunnel3Params       RemoteProcedure = 304
	RemoteProcDomainMigratePerform3Params             RemoteProcedure = 305
	RemoteProcDomainMigrateFinish3Params              RemoteProcedure = 306
	RemoteProcDomainMigrateConfirm3Params             RemoteProcedure = 307
	RemoteProcDomainSetMemoryStatsPeriod              RemoteProcedure = 308
	RemoteProcDomainCreateXmlWithFiles                RemoteProcedure = 309
	RemoteProcDomainCreateWithFiles                   RemoteProcedure = 310
	RemoteProcDomainEventDeviceRemoved                RemoteProcedure = 311
	RemoteProcConnectGetCpuModelNames                 RemoteProcedure = 312
	RemoteProcConnectNetworkEventRegisterAny          RemoteProcedure = 313
	RemoteProcConnectNetworkEventDeregisterAny        RemoteProcedure = 314
	RemoteProcNetworkEventLifecycle                   RemoteProcedure = 315
	RemoteProcConnectDomainEventCallbackRegisterAny   RemoteProcedure = 316
	RemoteProcConnectDomainEventCallbackDeregisterAny RemoteProcedure = 317
	RemoteProcDomainEventCallbackLifecycle            RemoteProcedure = 318
	RemoteProcDomainEventCallbackReboot               RemoteProcedure = 319
	RemoteProcDomainEventCallbackRtcChange            RemoteProcedure = 320
	RemoteProcDomainEventCallbackWatchdog             RemoteProcedure = 321
	RemoteProcDomainEventCallbackIoError              RemoteProcedure = 322
	RemoteProcDomainEventCallbackGraphics             RemoteProcedure = 323
	RemoteProcDomainEventCallbackIoErrorReason        RemoteProcedure = 324
	RemoteProcDomainEventCallbackControlError         RemoteProcedure = 325
	RemoteProcDomainEventCallbackBlockJob             RemoteProcedure = 326
	RemoteProcDomainEventCallbackDiskChange           RemoteProcedure = 327
	RemoteProcDomainEventCallbackTrayChange           RemoteProcedure = 328
	RemoteProcDomainEventCallbackPmwakeup             RemoteProcedure = 329
	RemoteProcDomainEventCallbackPmsuspend            RemoteProcedure = 330
	RemoteProcDomainEventCallbackBalloonChange        RemoteProcedure = 331
	RemoteProcDomainEventCallbackPmsuspendDisk        RemoteProcedure = 332
	RemoteProcDomainEventCallbackDeviceRemoved        RemoteProcedure = 333
	RemoteProcDomainCoreDumpWithFormat                RemoteProcedure = 334
	RemoteProcDomainFsfreeze                          RemoteProcedure = 335
	RemoteProcDomainFsthaw                            RemoteProcedure = 336
	RemoteProcDomainGetTime                           RemoteProcedure = 337
	RemoteProcDomainSetTime                           RemoteProcedure = 338
	RemoteProcDomainEventBlockJob2                    RemoteProcedure = 339
	RemoteProcNodeGetFreePages                        RemoteProcedure = 340
	RemoteProcNetworkGetDhcpLeases                    RemoteProcedure = 341
	RemoteProcConnectGetDomainCapabilities            RemoteProcedure = 342
	RemoteProcDomainOpenGraphicsFd                    RemoteProcedure = 343
	RemoteProcConnectGetAllDomainStats                RemoteProcedure = 344
	RemoteProcDomainBlockCopy                         RemoteProcedure = 345
	RemoteProcDomainEventCallbackTunable              RemoteProcedure = 346
	RemoteProcNodeAllocPages                          RemoteProcedure = 347
	RemoteProcDomainEventCallbackAgentLifecycle       RemoteProcedure = 348
	RemoteProcDomainGetFsinfo                         RemoteProcedure = 349
	RemoteProcDomainDefineXmlFlags                    RemoteProcedure = 350
	RemoteProcDomainGetIothreadInfo                   RemoteProcedure = 351
	RemoteProcDomainPinIothread                       RemoteProcedure = 352
	RemoteProcDomainInterfaceAddresses                RemoteProcedure = 353
	RemoteProcDomainEventCallbackDeviceAdded          RemoteProcedure = 354
	RemoteProcDomainAddIothread                       RemoteProcedure = 355
	RemoteProcDomainDelIothread                       RemoteProcedure = 356
	RemoteProcDomainSetUserPassword                   RemoteProcedure = 357
	RemoteProcDomainRename                            RemoteProcedure = 358
	RemoteProcDomainEventCallbackMigrationIteration   RemoteProcedure = 359
	RemoteProcConnectRegisterCloseCallback            RemoteProcedure = 360
	RemoteProcConnectUnregisterCloseCallback          RemoteProcedure = 361
	RemoteProcConnectEventConnectionClosed            RemoteProcedure = 362
	RemoteProcDomainEventCallbackJobCompleted         RemoteProcedure = 363
	RemoteProcDomainMigrateStartPostCopy              RemoteProcedure = 364
	RemoteProcDomainGetPerfEvents                     RemoteProcedure = 365
	RemoteProcDomainSetPerfEvents                     RemoteProcedure = 366
	RemoteProcDomainEventCallbackDeviceRemovalFailed  RemoteProcedure = 367
	RemoteProcConnectStoragePoolEventRegisterAny      RemoteProcedure = 368
	RemoteProcConnectStoragePoolEventDeregisterAny    RemoteProcedure = 369
	RemoteProcStoragePoolEventLifecycle               RemoteProcedure = 370
	RemoteProcDomainGetGuestVcpus                     RemoteProcedure = 371
	RemoteProcDomainSetGuestVcpus                     RemoteProcedure = 372
	RemoteProcStoragePoolEventRefresh                 RemoteProcedure = 373
	RemoteProcConnectNodeDeviceEventRegisterAny       RemoteProcedure = 374
	RemoteProcConnectNodeDeviceEventDeregisterAny     RemoteProcedure = 375
	RemoteProcNodeDeviceEventLifecycle                RemoteProcedure = 376
	RemoteProcNodeDeviceEventUpdate                   RemoteProcedure = 377
	RemoteProcStorageVolGetInfoFlags                  RemoteProcedure = 378
	RemoteProcDomainEventCallbackMetadataChange       RemoteProcedure = 379
	RemoteProcConnectSecretEventRegisterAny           RemoteProcedure = 380
	RemoteProcConnectSecretEventDeregisterAny         RemoteProcedure = 381
	RemoteProcSecretEventLifecycle                    RemoteProcedure = 382
	RemoteProcSecretEventValueChanged                 RemoteProcedure = 383
	RemoteProcDomainSetVcpu                           RemoteProcedure = 384
	RemoteProcDomainEventBlockThreshold               RemoteProcedure = 385
	RemoteProcDomainSetBlockThreshold                 RemoteProcedure = 38
)

type RemoteDomain struct {
	Name string
	UUID UUID
	ID   int
}

type RemoteNetwork struct {
	Name string
	UUID UUID
}

type RemoteNwFilter struct {
	Name string
	UUID UUID
}

type RemoteInterface struct {
	Name string
	Mac  string
}

type RemoteStoragePool struct {
	Name string
	UUID UUID
}

type RemoteStorageVolume struct {
	Pool string
	Name string
	Key  string
}

type RemoteNodeDevice struct {
	Name string
}

type RemoteSecret struct {
	UUID      UUID
	UsageType int
	UsageID   string
}

type RemoteDomainSnapshot struct {
	Name   string
	Domain *RemoteDomain
}

type RemoteError struct {
	Code     int
	DomainID int
	Message  string
	Level    int
	Domain   *RemoteDomain
	Str1     string
	Str2     string
	Str3     string
	Int1     int
	Int2     int
	Network  *RemoteNetwork
}

type RemoteVcpuInfo struct {
	Number  uint32
	State   int
	CpuTime uint64
	Cpu     int
}

type RemoteTypedParam struct {
	Field string
}

type RemoteNodeGetCpuStats struct {
	Field string
	Value uint64
}

type RemoteNodeGetMemoryStats struct {
	Field string
	Value uint64
}

type RemoteDomainDiskError struct {
	Disk  string
	Error int
}

type RemoteConnectOpenReq struct {
	Name  string
	Flags uint32
}

type RemoteConnectSupportsFeatureReq struct {
	Feature int
}

type RemoteConnectSupportsFeatureRes struct {
	Supported int
}

type RemoteConnectGetTypeRes struct {
	Type string
}

type RemoteConnectGetVersionRes struct {
	HvVer uint64
}

type RemoteConnectGetLibVersionRes struct {
	LibVer uint64
}

type RemoteConnectGetHostnameRes struct {
	Hostname string
}

type RemoteConnectGetSysinfoReq struct {
	Flags uint32
}

type RemoteConnectGetSysinfoRes struct {
	Sysinfo string
}

type RemoteConnectGetUriRes struct {
	Uri string
}

type RemoteConnectGetMaxVcpusReq struct {
	Type string
}

type RemoteConnectGetMaxVcpusRes struct {
	MaxVcpus int
}

type RemoteNodeGetInfoRes struct {
	Model   [32]byte
	Memory  uint64
	Cpus    int
	Mhz     int
	Nodes   int
	Sockets int
	Cores   int
	Threads int
}

type RemoteConnectGetCapabilitiesRes struct {
	Capabilities string
}

type RemoteConnectGetDomainCapabilitiesReq struct {
	Emulatorbin string
	Arch        string
	Machine     string
	Virttype    string
	Flags       uint32
}

type RemoteConnectGetDomainCapabilitiesRes struct {
	Capabilities string
}

type RemoteNodeGetCpuStatsReq struct {
	CpuNum  int
	Nparams int
	Flags   uint32
}

type RemoteNodeGetCpuStatsRes struct {
	Params  []RemoteNodeGetCpuStats
	Nparams int
}

type RemoteNodeGetMemoryStatsReq struct {
	Nparams int
	CellNum int
	Flags   uint32
}

type RemoteNodeGetMemoryStatsRes struct {
	Params  []RemoteNodeGetMemoryStats
	Nparams int
}

type RemoteNodeGetCellsFreeMemoryReq struct {
	StartCell int
	Maxcells  int
}

type RemoteNodeGetCellsFreeMemoryRes struct {
	Cells []uint64
}

type RemoteNodeGetFreeMemoryRes struct {
	FreeMem uint64
}

type RemoteDomainGetSchedulerTypeReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetSchedulerTypeRes struct {
	Type    string
	Nparams int
}

type RemoteDomainGetSchedulerParametersReq struct {
	Domain  *RemoteDomain
	Nparams int
}

type RemoteDomainGetSchedulerParametersRes struct {
}

type RemoteDomainGetSchedulerParametersFlagsReq struct {
	Domain  *RemoteDomain
	Nparams int
	Flags   uint32
}

type RemoteDomainGetSchedulerParametersFlagsRes struct {
}

type RemoteDomainSetSchedulerParametersReq struct {
	Domain *RemoteDomain
}

type RemoteDomainSetSchedulerParametersFlagsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainSetBlkioParametersReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetBlkioParametersReq struct {
	Domain  *RemoteDomain
	Nparams int
	Flags   uint32
}

type RemoteDomainGetBlkioParametersRes struct {
	Nparams int
}

type RemoteDomainSetMemoryParametersReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetMemoryParametersReq struct {
	Domain  *RemoteDomain
	Nparams int
	Flags   uint32
}

type RemoteDomainGetMemoryParametersRes struct {
	Nparams int
}

type RemoteDomainBlockResizeReq struct {
	Domain *RemoteDomain
	Disk   string
	Size   uint64
	Flags  uint32
}

type RemoteDomainSetNumaParametersReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetNumaParametersReq struct {
	Domain  *RemoteDomain
	Nparams int
	Flags   uint32
}

type RemoteDomainGetNumaParametersRes struct {
	Nparams int
}

type RemoteDomainSetPerfEventsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetPerfEventsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetPerfEventsRes struct {
}

type RemoteDomainBlockStatsReq struct {
	Domain *RemoteDomain
	Path   string
}

type RemoteDomainBlockStatsRes struct {
	RdReq   int64
	RdBytes int64
	WrReq   int64
	WrBytes int64
	Errs    int64
}

type RemoteDomainBlockStatsFlagsReq struct {
	Domain  *RemoteDomain
	Path    string
	Nparams int
	Flags   uint32
}

type RemoteDomainBlockStatsFlagsRes struct {
	Nparams int
}

type RemoteDomainInterfaceStatsReq struct {
	Domain *RemoteDomain
	Path   string
}

type RemoteDomainInterfaceStatsRes struct {
	RxBytes   int64
	RxPackets int64
	RxErrs    int64
	RxDrop    int64
	TxBytes   int64
	TxPackets int64
	TxErrs    int64
	TxDrop    int64
}

type RemoteDomainSetInterfaceParametersReq struct {
	Domain *RemoteDomain
	Device string
	Flags  uint32
}

type RemoteDomainGetInterfaceParametersReq struct {
	Domain  *RemoteDomain
	Device  string
	Nparams int
	Flags   uint32
}

type RemoteDomainGetInterfaceParametersRes struct {
	Nparams int
}

type RemoteDomainMemoryStatsReq struct {
	Domain   *RemoteDomain
	MaxStats uint32
	Flags    uint32
}

type RemoteDomainMemoryStat struct {
	Tag int
	Val uint64
}

type RemoteDomainMemoryStatsRes struct {
	Stats []RemoteDomainMemoryStat
}

type RemoteDomainBlockPeekReq struct {
	Domain *RemoteDomain
	Path   string
	Offset uint64
	Size   uint32
	Flags  uint32
}

type RemoteDomainBlockPeekRes struct {
	Buffer []byte
}

type RemoteDomainMemoryPeekReq struct {
	Domain *RemoteDomain
	Offset uint64
	Size   uint32
	Flags  uint32
}

type RemoteDomainMemoryPeekRes struct {
	Buffer []byte
}

type RemoteDomainGetBlockInfoReq struct {
	Domain *RemoteDomain
	Path   string
	Flags  uint32
}

type RemoteDomainGetBlockInfoRes struct {
	Allocation uint64
	Capacity   uint64
	Physical   uint64
}

type RemoteConnectListDomainsReq struct {
	Maxids int
}

type RemoteConnectListDomainsRes struct {
	Ids []int
}

type RemoteConnectNumOfDomainsRes struct {
	Num int
}

type RemoteDomainCreateXmlReq struct {
	XML   string
	Flags uint32
}

type RemoteDomainCreateXmlRes struct {
	Domain *RemoteDomain
}

type RemoteDomainCreateXmlWithFilesReq struct {
	XML   string
	Flags uint32
}

type RemoteDomainCreateXmlWithFilesRes struct {
	Domain *RemoteDomain
}

type RemoteDomainLookupByIdReq struct {
	Id int
}

type RemoteDomainLookupByIdRes struct {
	Domain *RemoteDomain
}

type RemoteDomainLookupByUuidReq struct {
	UUID UUID
}

type RemoteDomainLookupByUuidRes struct {
	Domain *RemoteDomain
}

type RemoteDomainLookupByNameReq struct {
	Name string
}

type RemoteDomainLookupByNameRes struct {
	Domain *RemoteDomain
}

type RemoteDomainSuspendReq struct {
	Domain *RemoteDomain
}

type RemoteDomainResumeReq struct {
	Domain *RemoteDomain
}

type RemoteDomainPmSuspendForDurationReq struct {
	Domain   *RemoteDomain
	Target   uint32
	Duration uint64
	Flags    uint32
}

type RemoteDomainPmWakeupReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainShutdownReq struct {
	Domain *RemoteDomain
}

type RemoteDomainRebootReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainResetReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainDestroyReq struct {
	Domain *RemoteDomain
}

type RemoteDomainDestroyFlagsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetOsTypeReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetOsTypeRes struct {
	Type string
}

type RemoteDomainGetMaxMemoryReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetMaxMemoryRes struct {
	Memory uint64
}

type RemoteDomainSetMaxMemoryReq struct {
	Domain *RemoteDomain
	Memory uint64
}

type RemoteDomainSetMemoryReq struct {
	Domain *RemoteDomain
	Memory uint64
}

type RemoteDomainSetMemoryFlagsReq struct {
	Domain *RemoteDomain
	Memory uint64
	Flags  uint32
}

type RemoteDomainSetMemoryStatsPeriodReq struct {
	Domain *RemoteDomain
	Period int
	Flags  uint32
}

type RemoteDomainGetInfoReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetInfoRes struct {
	State     byte
	MaxMem    uint64
	Memory    uint64
	NrVirtCpu uint8
	CpuTime   uint64
}

type RemoteDomainSaveReq struct {
	Domain *RemoteDomain
	To     string
}

type RemoteDomainSaveFlagsReq struct {
	Domain *RemoteDomain
	To     string
	XML    string
	Flags  uint32
}

type RemoteDomainRestoreReq struct {
	From string
}

type RemoteDomainRestoreFlagsReq struct {
	From  string
	XML   string
	Flags uint32
}

type RemoteDomainSaveImageGetXmlDescReq struct {
	File  string
	Flags uint32
}

type RemoteDomainSaveImageGetXmlDescRes struct {
	Xml string
}

type RemoteDomainSaveImageDefineXmlReq struct {
	File  string
	XML   string
	Flags uint32
}

type RemoteDomainCoreDumpReq struct {
	Domain *RemoteDomain
	To     string
	Flags  uint32
}

type RemoteDomainCoreDumpWithFormatReq struct {
	Domain     *RemoteDomain
	To         string
	Dumpformat uint32
	Flags      uint32
}

type RemoteDomainScreenshotReq struct {
	Domain *RemoteDomain
	Screen uint32
	Flags  uint32
}

type RemoteDomainScreenshotRes struct {
	Mime string
}

type RemoteDomainGetXmlDescReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetXmlDescRes struct {
	Xml string
}

type RemoteDomainMigratePrepareReq struct {
	UriIn    string
	Flags    uint64
	Dname    string
	Resource uint64
}

type RemoteDomainMigratePrepareRes struct {
	Cookie []byte
	UriOut string
}

type RemoteDomainMigratePerformReq struct {
	Domain   *RemoteDomain
	Cookie   []byte
	Uri      string
	Flags    uint64
	Dname    string
	Resource uint64
}

type RemoteDomainMigrateFinishReq struct {
	Dname  string
	Cookie []byte
	Uri    string
	Flags  uint64
}

type RemoteDomainMigrateFinishRes struct {
	Domain *RemoteDomain
}

type RemoteDomainMigratePrepare2Req struct {
	UriIn    string
	Flags    uint64
	Dname    string
	Resource uint64
	XML      string
}

type RemoteDomainMigratePrepare2Res struct {
	Cookie []byte
	UriOut string
}

type RemoteDomainMigrateFinish2Req struct {
	Dname   string
	Cookie  []byte
	Uri     string
	Flags   uint64
	Retcode int
}

type RemoteDomainMigrateFinish2Res struct {
	Domain *RemoteDomain
}

type RemoteConnectListDefinedDomainsReq struct {
	Maxnames int
}

type RemoteConnectListDefinedDomainsRes struct {
	Names []string
}

type RemoteConnectNumOfDefinedDomainsRes struct {
	Num int
}

type RemoteDomainCreateReq struct {
	Domain *RemoteDomain
}

type RemoteDomainCreateWithFlagsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainCreateWithFlagsRes struct {
	Domain *RemoteDomain
}

type RemoteDomainCreateWithFilesReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainCreateWithFilesRes struct {
	Domain *RemoteDomain
}

type RemoteDomainDefineXmlReq struct {
	Xml string
}

type RemoteDomainDefineXmlRes struct {
	Domain *RemoteDomain
}

type RemoteDomainDefineXmlFlagsReq struct {
	Xml   string
	Flags uint32
}

type RemoteDomainDefineXmlFlagsRes struct {
	Domain *RemoteDomain
}

type RemoteDomainUndefineReq struct {
	Domain *RemoteDomain
}

type RemoteDomainUndefineFlagsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainInjectNmiReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainSendKeyReq struct {
	Domain   *RemoteDomain
	Codeset  uint32
	Holdtime uint32
	Keycodes []uint32
	Flags    uint32
}

type RemoteDomainSendProcessSignalReq struct {
	Domain   *RemoteDomain
	PidValue int64
	Signum   uint32
	Flags    uint32
}

type RemoteDomainSetVcpusReq struct {
	Domain *RemoteDomain
	Nvcpus uint32
}

type RemoteDomainSetVcpusFlagsReq struct {
	Domain *RemoteDomain
	Nvcpus uint32
	Flags  uint32
}

type RemoteDomainGetVcpusFlagsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetVcpusFlagsRes struct {
	Num int
}

type RemoteDomainPinVcpuReq struct {
	Domain *RemoteDomain
	Vcpu   uint32
	Cpumap []byte
}

type RemoteDomainPinVcpuFlagsReq struct {
	Domain *RemoteDomain
	Vcpu   uint32
	Cpumap []byte
	Flags  uint32
}

type RemoteDomainGetVcpuPinInfoReq struct {
	Domain   *RemoteDomain
	Ncpumaps int
	Maplen   int
	Flags    uint32
}

type RemoteDomainGetVcpuPinInfoRes struct {
	Cpumaps []byte
	Num     int
}

type RemoteDomainPinEmulatorReq struct {
	Domain *RemoteDomain
	Cpumap []byte
	Flags  uint32
}

type RemoteDomainGetEmulatorPinInfoReq struct {
	Domain *RemoteDomain
	Maplen int
	Flags  uint32
}

type RemoteDomainGetEmulatorPinInfoRes struct {
	Cpumaps []byte
	Ret     int
}

type RemoteDomainGetVcpusReq struct {
	Domain  *RemoteDomain
	Maxinfo int
	Maplen  int
}

type RemoteDomainGetVcpusRes struct {
	Info    []RemoteVcpuInfo
	Cpumaps []byte
}

type RemoteDomainGetMaxVcpusReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetMaxVcpusRes struct {
	Num int
}

type RemoteDomainIothreadInfo struct {
	IothreadID uint32
	Cpumap     []byte
}

type RemoteDomainGetIothreadInfoReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetIothreadInfoRes struct {
	Info []RemoteDomainIothreadInfo
	Ret  uint32
}

type RemoteDomainPinIothreadReq struct {
	Domain      *RemoteDomain
	IothreadsId uint32
	Cpumap      []byte
	Flags       uint32
}

type RemoteDomainAddIothreadReq struct {
	Domain     *RemoteDomain
	IothreadId uint32
	Flags      uint32
}

type RemoteDomainDelIothreadReq struct {
	Domain     *RemoteDomain
	IothreadId uint32
	Flags      uint32
}

type RemoteDomainGetSecurityLabelReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetSecurityLabelRes struct {
	Label     []byte
	Enforcing int
}

type RemoteDomainGetSecurityLabelListReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetSecurityLabelListRes struct {
	Labels []RemoteDomainGetSecurityLabelRes
	Ret    int
}

type RemoteNodeGetSecurityModelRes struct {
	Model []byte
	Doi   []byte
}

type RemoteDomainAttachDeviceReq struct {
	Domain *RemoteDomain
	Xml    string
}

type RemoteDomainAttachDeviceFlagsReq struct {
	Domain *RemoteDomain
	Xml    string
	Flags  uint32
}

type RemoteDomainDetachDeviceReq struct {
	Domain *RemoteDomain
	Xml    string
}

type RemoteDomainDetachDeviceFlagsReq struct {
	Domain *RemoteDomain
	Xml    string
	Flags  uint32
}

type RemoteDomainUpdateDeviceFlagsReq struct {
	Domain *RemoteDomain
	Xml    string
	Flags  uint32
}

type RemoteDomainGetAutostartReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetAutostartRes struct {
	Autostart int
}

type RemoteDomainSetAutostartReq struct {
	Domain    *RemoteDomain
	Autostart int
}

type RemoteDomainSetMetadataReq struct {
	Domain   *RemoteDomain
	Type     int
	Metadata string
	Key      string
	Uri      string
	Flags    uint32
}

type RemoteDomainGetMetadataReq struct {
	Domain *RemoteDomain
	Type   int
	Uri    string
	Flags  uint32
}

type RemoteDomainGetMetadataRes struct {
	Metadata string
}

type RemoteDomainBlockJobAbortReq struct {
	Domain *RemoteDomain
	Path   string
	Flags  uint32
}

type RemoteDomainGetBlockJobInfoReq struct {
	Domain *RemoteDomain
	Path   string
	Flags  uint32
}

type RemoteDomainGetBlockJobInfoRes struct {
	Found     int
	Type      int
	Bandwidth uint64
	Cur       uint64
	End       uint64
}

type RemoteDomainBlockJobSetSpeedReq struct {
	Domain    *RemoteDomain
	Path      string
	Bandwidth uint64
	Flags     uint32
}

type RemoteDomainBlockPullReq struct {
	Domain    *RemoteDomain
	Path      string
	Bandwidth uint64
	Flags     uint32
}

type RemoteDomainBlockRebaseReq struct {
	Domain    *RemoteDomain
	Path      string
	Base      string
	Bandwidth uint64
	Flags     uint32
}

type RemoteDomainBlockCopyReq struct {
	Domain  *RemoteDomain
	Path    string
	Destxml string
	Flags   uint32
}

type RemoteDomainBlockCommitReq struct {
	Domain    *RemoteDomain
	Disk      string
	Base      string
	Top       string
	Bandwidth uint64
	Flags     uint32
}

type RemoteDomainSetBlockIoTuneReq struct {
	Domain *RemoteDomain
	Disk   string
	Flags  uint32
}

type RemoteDomainGetBlockIoTuneReq struct {
	Domain  *RemoteDomain
	Disk    string
	Nparams int
	Flags   uint32
}

type RemoteDomainGetBlockIoTuneRes struct {
	Nparams int
}

type RemoteDomainGetCpuStatsReq struct {
	Domain   *RemoteDomain
	Nparams  uint32
	StartCpu int
	Ncpus    uint32
	Flags    uint32
}

type RemoteDomainGetCpuStatsRes struct {
	Nparams int
}

type RemoteDomainGetHostnameReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetHostnameRes struct {
	Hostname string
}

type RemoteConnectNumOfNetworksRes struct {
	Num int
}

type RemoteConnectListNetworksReq struct {
	Maxnames int
}

type RemoteConnectListNetworksRes struct {
	Names []string
}

type RemoteConnectNumOfDefinedNetworksRes struct {
	Num int
}

type RemoteConnectListDefinedNetworksReq struct {
	Maxnames int
}

type RemoteConnectListDefinedNetworksRes struct {
	Names []string
}

type RemoteNetworkLookupByUuidReq struct {
	UUID UUID
}

type RemoteNetworkLookupByUuidRes struct {
	Network *RemoteNetwork
}

type RemoteNetworkLookupByNameReq struct {
	Name string
}

type RemoteNetworkLookupByNameRes struct {
	Network *RemoteNetwork
}

type RemoteNetworkCreateXmlReq struct {
	Xml string
}

type RemoteNetworkCreateXmlRes struct {
	Network *RemoteNetwork
}

type RemoteNetworkDefineXmlReq struct {
	Xml string
}

type RemoteNetworkDefineXmlRes struct {
	Network *RemoteNetwork
}

type RemoteNetworkUndefineReq struct {
	Network *RemoteNetwork
}

type RemoteNetworkUpdateReq struct {
	Network     *RemoteNetwork
	Command     uint32
	Section     uint32
	ParentIndex int
	Xml         string
	Flags       uint32
}

type RemoteNetworkCreateReq struct {
	Network *RemoteNetwork
}

type RemoteNetworkDestroyReq struct {
	Network *RemoteNetwork
}

type RemoteNetworkGetXmlDescReq struct {
	Network *RemoteNetwork
	Flags   uint32
}

type RemoteNetworkGetXmlDescRes struct {
	Xml string
}

type RemoteNetworkGetBridgeNameReq struct {
	Network *RemoteNetwork
}

type RemoteNetworkGetBridgeNameRes struct {
	Name string
}

type RemoteNetworkGetAutostartReq struct {
	Network *RemoteNetwork
}

type RemoteNetworkGetAutostartRes struct {
	Autostart int
}

type RemoteNetworkSetAutostartReq struct {
	Network   *RemoteNetwork
	Autostart int
}

type RemoteConnectNumOfNwfiltersRes struct {
	Num int
}

type RemoteConnectListNwfiltersReq struct {
	Maxnames int
}

type RemoteConnectListNwfiltersRes struct {
	Names []string
}

type RemoteNwfilterLookupByUuidReq struct {
	UUID UUID
}

type RemoteNwfilterLookupByUuidRes struct {
	Nwfilter *RemoteNwFilter
}

type RemoteNwfilterLookupByNameReq struct {
	Name string
}

type RemoteNwfilterLookupByNameRes struct {
	Nwfilter *RemoteNwFilter
}

type RemoteNwfilterDefineXmlReq struct {
	Xml string
}

type RemoteNwfilterDefineXmlRes struct {
	Nwfilter *RemoteNwFilter
}

type RemoteNwfilterUndefineReq struct {
	Nwfilter *RemoteNwFilter
}

type RemoteNwfilterGetXmlDescReq struct {
	Nwfilter *RemoteNwFilter
	Flags    uint32
}

type RemoteNwfilterGetXmlDescRes struct {
	Xml string
}

type RemoteConnectNumOfInterfacesRes struct {
	Num int
}

type RemoteConnectListInterfacesReq struct {
	Maxnames int
}

type RemoteConnectListInterfacesRes struct {
	Names []string
}

type RemoteConnectNumOfDefinedInterfacesRes struct {
	Num int
}

type RemoteConnectListDefinedInterfacesReq struct {
	Maxnames int
}

type RemoteConnectListDefinedInterfacesRes struct {
	Names []string
}

type RemoteInterfaceLookupByNameReq struct {
	Name string
}

type RemoteInterfaceLookupByNameRes struct {
	Iface *RemoteInterface
}

type RemoteInterfaceLookupByMacStringReq struct {
	Mac string
}

type RemoteInterfaceLookupByMacStringRes struct {
	Iface *RemoteInterface
}

type RemoteInterfaceGetXmlDescReq struct {
	Iface *RemoteInterface
	Flags uint32
}

type RemoteInterfaceGetXmlDescRes struct {
	Xml string
}

type RemoteInterfaceDefineXmlReq struct {
	Xml   string
	Flags uint32
}

type RemoteInterfaceDefineXmlRes struct {
	Iface *RemoteInterface
}

type RemoteInterfaceUndefineReq struct {
	Iface *RemoteInterface
}

type RemoteInterfaceCreateReq struct {
	Iface *RemoteInterface
	Flags uint32
}

type RemoteInterfaceDestroyReq struct {
	Iface *RemoteInterface
	Flags uint32
}

type RemoteInterfaceChangeBeginReq struct {
	Flags uint32
}

type RemoteInterfaceChangeCommitReq struct {
	Flags uint32
}

type RemoteInterfaceChangeRollbackReq struct {
	Flags uint32
}

type RemoteAuthListRes struct {
	Types []RemoteAuthType
}

type RemoteAuthSaslInitRes struct {
	Mechlist string
}

type RemoteAuthSaslStartReq struct {
	Mech string
	Nil  int
	Data []byte
}

type RemoteAuthSaslStartRes struct {
	Complete int
	Nil      int
	Data     []byte
}

type RemoteAuthSaslStepReq struct {
	Nil  int
	Data []byte
}

type RemoteAuthSaslStepRes struct {
	Complete int
	Nil      int
	Data     []byte
}

type RemoteAuthPolkitRes struct {
	Complete int
}

type RemoteConnectNumOfStoragePoolsRes struct {
	Num int
}

type RemoteConnectListStoragePoolsReq struct {
	Maxnames int
}

type RemoteConnectListStoragePoolsRes struct {
	Names []string
}

type RemoteConnectNumOfDefinedStoragePoolsRes struct {
	Num int
}

type RemoteConnectListDefinedStoragePoolsReq struct {
	Maxnames int
}

type RemoteConnectListDefinedStoragePoolsRes struct {
	Names []string
}

type RemoteConnectFindStoragePoolSourcesReq struct {
	Type    string
	SrcSpec string
	Flags   uint32
}

type RemoteConnectFindStoragePoolSourcesRes struct {
	Xml string
}

type RemoteStoragePoolLookupByUuidReq struct {
	UUID UUID
}

type RemoteStoragePoolLookupByUuidRes struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolLookupByNameReq struct {
	Name string
}

type RemoteStoragePoolLookupByNameRes struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolLookupByVolumeReq struct {
	Vol *RemoteStorageVolume
}

type RemoteStoragePoolLookupByVolumeRes struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolCreateXmlReq struct {
	Xml   string
	Flags uint32
}

type RemoteStoragePoolCreateXmlRes struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolDefineXmlReq struct {
	Xml   string
	Flags uint32
}

type RemoteStoragePoolDefineXmlRes struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolBuildReq struct {
	Pool  *RemoteStoragePool
	Flags uint32
}

type RemoteStoragePoolUndefineReq struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolCreateReq struct {
	Pool  *RemoteStoragePool
	Flags uint32
}

type RemoteStoragePoolDestroyReq struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolDeleteReq struct {
	Pool  *RemoteStoragePool
	Flags uint32
}

type RemoteStoragePoolRefreshReq struct {
	Pool  *RemoteStoragePool
	Flags uint32
}

type RemoteStoragePoolGetXmlDescReq struct {
	Pool  *RemoteStoragePool
	Flags uint32
}

type RemoteStoragePoolGetXmlDescRes struct {
	Xml string
}

type RemoteStoragePoolGetInfoReq struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolGetInfoRes struct {
	State      byte
	Capacity   uint64
	Allocation uint64
	Available  uint64
}

type RemoteStoragePoolGetAutostartReq struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolGetAutostartRes struct {
	Autostart int
}

type RemoteStoragePoolSetAutostartReq struct {
	Pool      *RemoteStoragePool
	Autostart int
}

type RemoteStoragePoolNumOfVolumesReq struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolNumOfVolumesRes struct {
	Num int
}

type RemoteStoragePoolListVolumesReq struct {
	Pool     *RemoteStoragePool
	Maxnames int
}

type RemoteStoragePoolListVolumesRes struct {
	Names []string
}

type RemoteStorageVolLookupByNameReq struct {
	Pool *RemoteStoragePool
	Name string
}

type RemoteStorageVolLookupByNameRes struct {
	Vol *RemoteStorageVolume
}

type RemoteStorageVolLookupByKeyReq struct {
	Key string
}

type RemoteStorageVolLookupByKeyRes struct {
	Vol *RemoteStorageVolume
}

type RemoteStorageVolLookupByPathReq struct {
	Path string
}

type RemoteStorageVolLookupByPathRes struct {
	Vol *RemoteStorageVolume
}

type RemoteStorageVolCreateXmlReq struct {
	Pool  *RemoteStoragePool
	Xml   string
	Flags uint32
}

type RemoteStorageVolCreateXmlRes struct {
	Vol *RemoteStorageVolume
}

type RemoteStorageVolCreateXmlFromReq struct {
	Pool     *RemoteStoragePool
	Xml      string
	Clonevol *RemoteStorageVolume
	Flags    uint32
}

type RemoteStorageVolCreateXmlFromRes struct {
	Vol *RemoteStorageVolume
}

type RemoteStorageVolDeleteReq struct {
	Vol   *RemoteStorageVolume
	Flags uint32
}

type RemoteStorageVolWipeReq struct {
	Vol   *RemoteStorageVolume
	Flags uint32
}

type RemoteStorageVolWipePatternReq struct {
	Vol       *RemoteStorageVolume
	Algorithm uint32
	Flags     uint32
}

type RemoteStorageVolGetXmlDescReq struct {
	Vol   *RemoteStorageVolume
	Flags uint32
}

type RemoteStorageVolGetXmlDescRes struct {
	Xml string
}

type RemoteStorageVolGetInfoReq struct {
	Vol *RemoteStorageVolume
}

type RemoteStorageVolGetInfoRes struct {
	Type       byte
	Capacity   uint64
	Allocation uint64
}

type RemoteStorageVolGetInfoFlagsReq struct {
	Vol   *RemoteStorageVolume
	Flags uint32
}

type RemoteStorageVolGetInfoFlagsRes struct {
	Type       byte
	Capacity   uint64
	Allocation uint64
}

type RemoteStorageVolGetPathReq struct {
	Vol *RemoteStorageVolume
}

type RemoteStorageVolGetPathRes struct {
	Name string
}

type RemoteStorageVolResizeReq struct {
	Vol      *RemoteStorageVolume
	Capacity uint64
	Flags    uint32
}

type RemoteNodeNumOfDevicesReq struct {
	Cap   string
	Flags uint32
}

type RemoteNodeNumOfDevicesRes struct {
	Num int
}

type RemoteNodeListDevicesReq struct {
	Cap      string
	Maxnames int
	Flags    uint32
}

type RemoteNodeListDevicesRes struct {
	Names []string
}

type RemoteNodeDeviceLookupByNameReq struct {
	Name string
}

type RemoteNodeDeviceLookupByNameRes struct {
	Device *RemoteNodeDevice
}

type RemoteNodeDeviceLookupScsiHostByWwnReq struct {
	Wwnn  string
	Wwpn  string
	Flags uint32
}

type RemoteNodeDeviceLookupScsiHostByWwnRes struct {
	Device *RemoteNodeDevice
}

type RemoteNodeDeviceGetXmlDescReq struct {
	Name  string
	Flags uint32
}

type RemoteNodeDeviceGetXmlDescRes struct {
	Xml string
}

type RemoteNodeDeviceGetParentReq struct {
	Name string
}

type RemoteNodeDeviceGetParentRes struct {
	Parent string
}

type RemoteNodeDeviceNumOfCapsReq struct {
	Name string
}

type RemoteNodeDeviceNumOfCapsRes struct {
	Num int
}

type RemoteNodeDeviceListCapsReq struct {
	Name     string
	Maxnames int
}

type RemoteNodeDeviceListCapsRes struct {
	Names []string
}

type RemoteNodeDeviceDettachReq struct {
	Name string
}

type RemoteNodeDeviceDetachFlagsReq struct {
	Name       string
	DriverName string
	Flags      uint32
}

type RemoteNodeDeviceReAttachReq struct {
	Name string
}

type RemoteNodeDeviceResetReq struct {
	Name string
}

type RemoteNodeDeviceCreateXmlReq struct {
	XML   string
	Flags uint32
}

type RemoteNodeDeviceCreateXmlRes struct {
	Device *RemoteNodeDevice
}

type RemoteNodeDeviceDestroyReq struct {
	Name string
}

type RemoteConnectDomainEventRegisterRes struct {
	CbRegistered int
}

type RemoteConnectDomainEventDeregisterRes struct {
	CbRegistered int
}

type RemoteDomainEventLifecycleMsg struct {
	Domain *RemoteDomain
	Event  int
	Detail int
}

type RemoteDomainEventCallbackLifecycleMsg struct {
	CallbackID int
	Msg        RemoteDomainEventLifecycleMsg
}

type RemoteConnectDomainXmlFromNativeReq struct {
	NativeFormat string
	NativeConfig string
	Flags        uint32
}

type RemoteConnectDomainXmlFromNativeRes struct {
	DomainXml string
}

type RemoteConnectDomainXmlToNativeReq struct {
	NativeFormat string
	DomainXml    string
	Flags        uint32
}

type RemoteConnectDomainXmlToNativeRes struct {
	NativeConfig string
}

type RemoteConnectNumOfSecretsRes struct {
	Num int
}

type RemoteConnectListSecretsReq struct {
	Maxuuids int
}

type RemoteConnectListSecretsRes struct {
	Uuids []string
}

type RemoteSecretLookupByUuidReq struct {
	UUID UUID
}

type RemoteSecretLookupByUuidRes struct {
	Secret *RemoteSecret
}

type RemoteSecretDefineXmlReq struct {
	Xml   string
	Flags uint32
}

type RemoteSecretDefineXmlRes struct {
	Secret *RemoteSecret
}

type RemoteSecretGetXmlDescReq struct {
	Secret *RemoteSecret
	Flags  uint32
}

type RemoteSecretGetXmlDescRes struct {
	Xml string
}

type RemoteSecretSetValueReq struct {
	Secret *RemoteSecret
	Value  []byte
	Flags  uint32
}

type RemoteSecretGetValueReq struct {
	Secret *RemoteSecret
	Flags  uint32
}

type RemoteSecretGetValueRes struct {
	Value []byte
}

type RemoteSecretUndefineReq struct {
	Secret *RemoteSecret
}

type RemoteSecretLookupByUsageReq struct {
	UsageType int
	UsageID   string
}

type RemoteSecretLookupByUsageRes struct {
	Secret *RemoteSecret
}

type RemoteDomainMigratePrepareTunnelReq struct {
	Flags    uint64
	Dname    string
	Resource uint64
	XML      string
}

type RemoteConnectIsSecureRes struct {
	Secure int
}

type RemoteDomainIsActiveReq struct {
	Domain *RemoteDomain
}

type RemoteDomainIsActiveRes struct {
	Active int
}

type RemoteDomainIsPersistentReq struct {
	Domain *RemoteDomain
}

type RemoteDomainIsPersistentRes struct {
	Persistent int
}

type RemoteDomainIsUpdatedReq struct {
	Domain *RemoteDomain
}

type RemoteDomainIsUpdatedRes struct {
	Updated int
}

type RemoteNetworkIsActiveReq struct {
	Network *RemoteNetwork
}

type RemoteNetworkIsActiveRes struct {
	Active int
}

type RemoteNetworkIsPersistentReq struct {
	Network *RemoteNetwork
}

type RemoteNetworkIsPersistentRes struct {
	Persistent int
}

type RemoteStoragePoolIsActiveReq struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolIsActiveRes struct {
	Active int
}

type RemoteStoragePoolIsPersistentReq struct {
	Pool *RemoteStoragePool
}

type RemoteStoragePoolIsPersistentRes struct {
	Persistent int
}

type RemoteInterfaceIsActiveReq struct {
	Iface *RemoteInterface
}

type RemoteInterfaceIsActiveRes struct {
	Active int
}

type RemoteConnectCompareCpuReq struct {
	Xml   string
	Flags uint32
}

type RemoteConnectCompareCpuRes struct {
	Result int
}

type RemoteConnectBaselineCpuReq struct {
	XmlCPUs []string
	Flags   uint32
}

type RemoteConnectBaselineCpuRes struct {
	Cpu string
}

type RemoteDomainGetJobInfoReq struct {
	Domain *RemoteDomain
}

type RemoteDomainGetJobInfoRes struct {
	Type int

	TimeElapsed   uint64
	TimeRemaining uint64

	DataTotal     uint64
	DataProcessed uint64
	DataRemaining uint64

	MemTotal     uint64
	MemProcessed uint64
	MemRemaining uint64

	FileTotal     uint64
	FileProcessed uint64
	FileRemaining uint64
}

type RemoteDomainGetJobStatsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetJobStatsRes struct {
	Type int
}

type RemoteDomainAbortJobReq struct {
	Domain *RemoteDomain
}

type RemoteDomainMigrateSetMaxDowntimeReq struct {
	Domain   *RemoteDomain
	Downtime uint64
	Flags    uint32
}

type RemoteDomainMigrateGetCompressionCacheReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainMigrateGetCompressionCacheRes struct {
	CacheSize uint64
}

type RemoteDomainMigrateSetCompressionCacheReq struct {
	Domain    *RemoteDomain
	CacheSize uint64
	Flags     uint32
}

type RemoteDomainMigrateSetMaxSpeedReq struct {
	Domain    *RemoteDomain
	Bandwidth uint64
	Flags     uint32
}

type RemoteDomainMigrateGetMaxSpeedReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainMigrateGetMaxSpeedRes struct {
	Bandwidth uint64
}

type RemoteConnectDomainEventRegisterAnyReq struct {
	EventID int
}

type RemoteConnectDomainEventDeregisterAnyReq struct {
	EventID int
}

type RemoteConnectDomainEventCallbackRegisterAnyReq struct {
	EventID int
	Domain  *RemoteDomain
}

type RemoteConnectDomainEventCallbackRegisterAnyRes struct {
	CallbackID int
}

type RemoteConnectDomainEventCallbackDeregisterAnyReq struct {
	CallbackID int
}

type RemoteDomainEventRebootMsg struct {
	Domain *RemoteDomain
}

type RemoteDomainEventCallbackRebootMsg struct {
	CallbackID int
	Msg        RemoteDomainEventRebootMsg
}

type RemoteDomainEventRtcChangeMsg struct {
	Domain *RemoteDomain
	Offset int64
}

type RemoteDomainEventCallbackRtcChangeMsg struct {
	CallbackID int
	Msg        RemoteDomainEventRtcChangeMsg
}

type RemoteDomainEventWatchdogMsg struct {
	Domain *RemoteDomain
	Action int
}

type RemoteDomainEventCallbackWatchdogMsg struct {
	CallbackID int
	Msg        RemoteDomainEventWatchdogMsg
}

type RemoteDomainEventIoErrorMsg struct {
	Domain   *RemoteDomain
	SrcPath  string
	DevAlias string
	Action   int
}

type RemoteDomainEventCallbackIoErrorMsg struct {
	CallbackID int
	Msg        RemoteDomainEventIoErrorMsg
}

type RemoteDomainEventIoErrorReasonMsg struct {
	Domain   *RemoteDomain
	SrcPath  string
	DevAlias string
	Action   int
	Reason   string
}

type RemoteDomainEventCallbackIoErrorReasonMsg struct {
	CallbackID int
	Msg        RemoteDomainEventIoErrorReasonMsg
}

type RemoteDomainEventGraphicsAddress struct {
	Family  int
	Node    string
	Service string
}

type RemoteDomainEventGraphicsIdentity struct {
	Type string
	Name string
}

type RemoteDomainEventGraphicsMsg struct {
	Domain     *RemoteDomain
	Phase      int
	Local      RemoteDomainEventGraphicsAddress
	Remote     RemoteDomainEventGraphicsAddress
	AuthScheme string
	Subject    []RemoteDomainEventGraphicsIdentity
}

type RemoteDomainEventCallbackGraphicsMsg struct {
	CallbackID int
	Msg        RemoteDomainEventGraphicsMsg
}

type RemoteDomainEventBlockJobMsg struct {
	Domain *RemoteDomain
	Path   string
	Type   int
	Status int
}

type RemoteDomainEventCallbackBlockJobMsg struct {
	CallbackID int
	Msg        RemoteDomainEventBlockJobMsg
}

type RemoteDomainEventDiskChangeMsg struct {
	Domain     *RemoteDomain
	OldSrcPath string
	NewSrcPath string
	DevAlias   string
	Reason     int
}

type RemoteDomainEventCallbackDiskChangeMsg struct {
	CallbackID int
	Msg        RemoteDomainEventDiskChangeMsg
}

type RemoteDomainEventTrayChangeMsg struct {
	Domain   *RemoteDomain
	DevAlias string
	Reason   int
}

type RemoteDomainEventCallbackTrayChangeMsg struct {
	CallbackID int
	Msg        RemoteDomainEventTrayChangeMsg
}

type RemoteDomainEventPmwakeupMsg struct {
	Domain *RemoteDomain
}

type RemoteDomainEventCallbackPmwakeupMsg struct {
	CallbackID int
	Reason     int
	Msg        RemoteDomainEventPmwakeupMsg
}

type RemoteDomainEventPmsuspendMsg struct {
	Domain *RemoteDomain
}

type RemoteDomainEventCallbackPmsuspendMsg struct {
	CallbackID int
	Reason     int
	Msg        RemoteDomainEventPmsuspendMsg
}

type RemoteDomainEventBalloonChangeMsg struct {
	Domain *RemoteDomain
	Actual uint64
}

type RemoteDomainEventCallbackBalloonChangeMsg struct {
	CallbackID int
	Msg        RemoteDomainEventBalloonChangeMsg
}

type RemoteDomainEventPmsuspendDiskMsg struct {
	Domain *RemoteDomain
}

type RemoteDomainEventCallbackPmsuspendDiskMsg struct {
	CallbackID int
	Reason     int
	Msg        RemoteDomainEventPmsuspendDiskMsg
}

type RemoteDomainManagedSaveReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainHasManagedSaveImageReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainHasManagedSaveImageRes struct {
	Result int
}

type RemoteDomainManagedSaveRemoveReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainSnapshotCreateXmlReq struct {
	Domain *RemoteDomain
	XML    string
	Flags  uint32
}

type RemoteDomainSnapshotCreateXmlRes struct {
	Snapshot *RemoteDomainSnapshot
}

type RemoteDomainSnapshotGetXmlDescReq struct {
	Snapshot *RemoteDomainSnapshot
	Flags    uint32
}

type RemoteDomainSnapshotGetXmlDescRes struct {
	Xml string
}

type RemoteDomainSnapshotNumReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainSnapshotNumRes struct {
	Num int
}

type RemoteDomainSnapshotListNamesReq struct {
	Domain   *RemoteDomain
	Maxnames int
	Flags    uint32
}

type RemoteDomainSnapshotListNamesRes struct {
	Names []string
}

type RemoteDomainListAllSnapshotsReq struct {
	Domain      *RemoteDomain
	NeedResults int
	Flags       uint32
}

type RemoteDomainListAllSnapshotsRes struct {
	Snapshots []*RemoteDomainSnapshot
	Ret       int
}

type RemoteDomainSnapshotNumChildrenReq struct {
	Snapshot *RemoteDomainSnapshot
	Flags    uint32
}

type RemoteDomainSnapshotNumChildrenRes struct {
	Num int
}

type RemoteDomainSnapshotListChildrenNamesReq struct {
	Snapshot *RemoteDomainSnapshot
	Maxnames int
	Flags    uint32
}

type RemoteDomainSnapshotListChildrenNamesRes struct {
	Names []string
}

type RemoteDomainSnapshotListAllChildrenReq struct {
	Snapshot    *RemoteDomainSnapshot
	NeedResults int
	Flags       uint32
}

type RemoteDomainSnapshotListAllChildrenRes struct {
	Snapshots []*RemoteDomainSnapshot
	Ret       int
}

type RemoteDomainSnapshotLookupByNameReq struct {
	Domain *RemoteDomain
	Name   string
	Flags  uint32
}

type RemoteDomainSnapshotLookupByNameRes struct {
	Snapshot *RemoteDomainSnapshot
}

type RemoteDomainHasCurrentSnapshotReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainHasCurrentSnapshotRes struct {
	Result int
}

type RemoteDomainSnapshotGetParentReq struct {
	Snapshot *RemoteDomainSnapshot
	Flags    uint32
}

type RemoteDomainSnapshotGetParentRes struct {
	Snapshot *RemoteDomainSnapshot
}

type RemoteDomainSnapshotCurrentReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainSnapshotCurrentRes struct {
	Snapshot *RemoteDomainSnapshot
}

type RemoteDomainSnapshotIsCurrentReq struct {
	Snapshot *RemoteDomainSnapshot
	Flags    uint32
}

type RemoteDomainSnapshotIsCurrentRes struct {
	Current int
}

type RemoteDomainSnapshotHasMetadataReq struct {
	Snapshot *RemoteDomainSnapshot
	Flags    uint32
}

type RemoteDomainSnapshotHasMetadataRes struct {
	Metadata int
}

type RemoteDomainRevertToSnapshotReq struct {
	Snapshot *RemoteDomainSnapshot
	Flags    uint32
}

type RemoteDomainSnapshotDeleteReq struct {
	Snapshot *RemoteDomainSnapshot
	Flags    uint32
}

type RemoteDomainOpenConsoleReq struct {
	Domain  *RemoteDomain
	DevName string
	Flags   uint32
}

type RemoteDomainOpenChannelReq struct {
	Domain *RemoteDomain
	Name   string
	Flags  uint32
}

type RemoteStorageVolUploadReq struct {
	Vol    *RemoteStorageVolume
	Offset uint64
	Length uint64
	Flags  uint32
}

type RemoteStorageVolDownloadReq struct {
	Vol    *RemoteStorageVolume
	Offset uint64
	Length uint64
	Flags  uint32
}

type RemoteDomainGetStateReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetStateRes struct {
	State  int
	Reason int
}

type RemoteDomainMigrateBegin3Req struct {
	Domain   *RemoteDomain
	Xmlin    string
	Flags    uint64
	Dname    string
	Resource uint64
}

type RemoteDomainMigrateBegin3Res struct {
	CookieOut []byte
	Xml       string
}

type RemoteDomainMigratePrepare3Req struct {
	CookieIn []byte
	UriIn    string
	Flags    uint64
	Dname    string
	Resource uint64
	XML      string
}

type RemoteDomainMigratePrepare3Res struct {
	CookieOut []byte
	UriOut    string
}

type RemoteDomainMigratePrepareTunnel3Req struct {
	CookieIn []byte
	Flags    uint64
	Dname    string
	Resource uint64
	XML      string
}

type RemoteDomainMigratePrepareTunnel3Res struct {
	CookieOut []byte
}

type RemoteDomainMigratePerform3Req struct {
	Domain   *RemoteDomain
	Xmlin    string
	CookieIn []byte
	Dconnuri string
	Uri      string
	Flags    uint64
	Dname    string
	Resource uint64
}

type RemoteDomainMigratePerform3Res struct {
	CookieOut []byte
}

type RemoteDomainMigrateFinish3Req struct {
	Dname     string
	CookieIn  []byte
	Dconnuri  string
	Uri       string
	Flags     uint64
	Cancelled int
}

type RemoteDomainMigrateFinish3Res struct {
	Domain    *RemoteDomain
	CookieOut []byte
}

type RemoteDomainMigrateConfirm3Req struct {
	Domain    *RemoteDomain
	CookieIn  []byte
	Flags     uint64
	Cancelled int
}

type RemoteDomainEventControlErrorMsg struct {
	Domain *RemoteDomain
}

type RemoteDomainEventCallbackControlErrorMsg struct {
	CallbackID int
	Msg        RemoteDomainEventControlErrorMsg
}

type RemoteDomainGetControlInfoReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetControlInfoRes struct {
	State     uint32
	Details   uint32
	StateTime uint64
}

type RemoteDomainOpenGraphicsReq struct {
	Domain *RemoteDomain
	Idx    uint32
	Flags  uint32
}

type RemoteDomainOpenGraphicsFdReq struct {
	Domain *RemoteDomain
	Idx    uint32
	Flags  uint32
}

type RemoteNodeSuspendForDurationReq struct {
	Target   uint32
	Duration uint64
	Flags    uint32
}

type RemoteDomainShutdownFlagsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetDiskErrorsReq struct {
	Domain    *RemoteDomain
	Maxerrors uint32
	Flags     uint32
}

type RemoteDomainGetDiskErrorsRes struct {
	Errors  []RemoteDomainDiskError
	Nerrors int
}

type RemoteConnectListAllDomainsReq struct {
	NeedResults int
	Flags       uint32
}

type RemoteConnectListAllDomainsRes struct {
	Domains []*RemoteDomain
	Ret     uint32
}

type RemoteConnectListAllStoragePoolsReq struct {
	NeedResults int
	Flags       uint32
}

type RemoteConnectListAllStoragePoolsRes struct {
	Pools []*RemoteStoragePool
	Ret   uint32
}

type RemoteStoragePoolListAllVolumesReq struct {
	Pool        *RemoteStoragePool
	NeedResults int
	Flags       uint32
}

type RemoteStoragePoolListAllVolumesRes struct {
	Vols []*RemoteStorageVolume
	Ret  uint32
}

type RemoteConnectListAllNetworksReq struct {
	NeedResults int
	Flags       uint32
}

type RemoteConnectListAllNetworksRes struct {
	Nets []*RemoteNetwork
	Ret  uint32
}

type RemoteConnectListAllInterfacesReq struct {
	NeedResults int
	Flags       uint32
}

type RemoteConnectListAllInterfacesRes struct {
	Ifaces []*RemoteInterface
	Ret    uint32
}

type RemoteConnectListAllNodeDevicesReq struct {
	NeedResults int
	Flags       uint32
}

type RemoteConnectListAllNodeDevicesRes struct {
	Devices []*RemoteNodeDevice
	Ret     uint32
}

type RemoteConnectListAllNwfiltersReq struct {
	NeedResults int
	Flags       uint32
}

type RemoteConnectListAllNwfiltersRes struct {
	Filters []*RemoteNwFilter
	Ret     uint32
}

type RemoteConnectListAllSecretsReq struct {
	NeedResults int
	Flags       uint32
}

type RemoteConnectListAllSecretsRes struct {
	Secrets []*RemoteSecret
	Ret     uint32
}

type RemoteNodeSetMemoryParametersReq struct {
	Flags uint32
}

type RemoteNodeGetMemoryParametersReq struct {
	Nparams int
	Flags   uint32
}

type RemoteNodeGetMemoryParametersRes struct {
	Nparams int
}

type RemoteNodeGetCpuMapReq struct {
	NeedMap    int
	NeedOnline int
	Flags      uint32
}

type RemoteNodeGetCpuMapRes struct {
	Cpumap []byte
	Online uint32
	Ret    int
}

type RemoteDomainFstrimReq struct {
	Domain     *RemoteDomain
	MountPoint string
	Minimum    uint64
	Flags      uint32
}

type RemoteDomainGetTimeReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetTimeRes struct {
	Seconds  int64
	Nseconds uint32
}

type RemoteDomainSetTimeReq struct {
	Domain   *RemoteDomain
	Seconds  int64
	Nseconds uint32
	Flags    uint32
}

type RemoteDomainMigrateBegin3ParamsReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainMigrateBegin3ParamsRes struct {
	CookieOut []byte
	Xml       string
}

type RemoteDomainMigratePrepare3ParamsReq struct {
	CookieIn []byte
	Flags    uint32
}

type RemoteDomainMigratePrepare3ParamsRes struct {
	CookieOut []byte
	UriOut    string
}

type RemoteDomainMigratePrepareTunnel3ParamsReq struct {
	CookieIn []byte
	Flags    uint32
}

type RemoteDomainMigratePrepareTunnel3ParamsRes struct {
	CookieOut []byte
}

type RemoteDomainMigratePerform3ParamsReq struct {
	Domain   *RemoteDomain
	Dconnuri string
	CookieIn []byte
	Flags    uint32
}

type RemoteDomainMigratePerform3ParamsRes struct {
	CookieOut []byte
}

type RemoteDomainMigrateFinish3ParamsReq struct {
	CookieIn  []byte
	Flags     uint32
	Cancelled int
}

type RemoteDomainMigrateFinish3ParamsRes struct {
	Domain    *RemoteDomain
	CookieOut []byte
}

type RemoteDomainMigrateConfirm3ParamsReq struct {
	Domain    *RemoteDomain
	CookieIn  []byte
	Flags     uint32
	Cancelled int
}

type RemoteDomainEventDeviceRemovedMsg struct {
	Domain   *RemoteDomain
	DevAlias string
}

type RemoteDomainEventCallbackDeviceRemovedMsg struct {
	CallbackID int
	Msg        RemoteDomainEventDeviceRemovedMsg
}

type RemoteDomainEventBlockJob2Msg struct {
	CallbackID int
	Domain     *RemoteDomain
	Dst        string
	Type       int
	Status     int
}

type RemoteDomainEventBlockThresholdMsg struct {
	CallbackID int
	Domain     *RemoteDomain
	Device     string
	Path       string
	Threshold  uint64
	Excess     uint64
}

type RemoteDomainEventCallbackTunableMsg struct {
	CallbackID int
	Domain     *RemoteDomain
}

type RemoteDomainEventCallbackDeviceAddedMsg struct {
	CallbackID int
	Domain     *RemoteDomain
	DevAlias   string
}

type RemoteConnectEventConnectionClosedMsg struct {
	Reason int
}

type RemoteConnectGetCpuModelNamesReq struct {
	Arch        string
	NeedResults int
	Flags       uint32
}

type RemoteConnectGetCpuModelNamesRes struct {
	Models []string
	Ret    int
}

type RemoteConnectNetworkEventRegisterAnyReq struct {
	EventID int
	Network *RemoteNetwork
}

type RemoteConnectNetworkEventRegisterAnyRes struct {
	CallbackID int
}

type RemoteConnectNetworkEventDeregisterAnyReq struct {
	CallbackID int
}

type RemoteNetworkEventLifecycleMsg struct {
	CallbackID int
	Network    *RemoteNetwork
	Event      int
	Detail     int
}

type RemoteConnectStoragePoolEventRegisterAnyReq struct {
	EventID int
	Pool    *RemoteStoragePool
}

type RemoteConnectStoragePoolEventRegisterAnyRes struct {
	CallbackID int
}

type RemoteConnectStoragePoolEventDeregisterAnyReq struct {
	CallbackID int
}

type RemoteStoragePoolEventLifecycleMsg struct {
	CallbackID int
	Pool       *RemoteStoragePool
	Event      int
	Detail     int
}

type RemoteStoragePoolEventRefreshMsg struct {
	CallbackID int
	Pool       *RemoteStoragePool
}

type RemoteConnectNodeDeviceEventRegisterAnyReq struct {
	EventID int
	Device  *RemoteNodeDevice
}

type RemoteConnectNodeDeviceEventRegisterAnyRes struct {
	CallbackID int
}

type RemoteConnectNodeDeviceEventDeregisterAnyReq struct {
	CallbackID int
}

type RemoteNodeDeviceEventLifecycleMsg struct {
	CallbackID int
	Device     *RemoteNodeDevice
	Event      int
	Detail     int
}

type RemoteNodeDeviceEventUpdateMsg struct {
	CallbackID int
	Device     *RemoteNodeDevice
}

type RemoteDomainFsfreezeReq struct {
	Domain      *RemoteDomain
	Mountpoints []string
	Flags       uint32
}

type RemoteDomainFsfreezeRes struct {
	Filesystems int
}

type RemoteDomainFsthawReq struct {
	Domain      *RemoteDomain
	Mountpoints []string
	Flags       uint32
}

type RemoteDomainFsthawRes struct {
	Filesystems int
}

type RemoteNodeGetFreePagesReq struct {
	Pages     []uint32
	StartCell int
	CellCount uint32
	Flags     uint32
}

type RemoteNodeGetFreePagesRes struct {
	Counts []uint64
}

type RemoteNodeAllocPagesReq struct {
	PageSizes  []uint32
	PageCounts []uint64
	StartCell  int
	CellCount  uint32
	Flags      uint32
}

type RemoteNodeAllocPagesRes struct {
	Ret int
}

type RemoteNetworkDhcpLease struct {
	Iface      string
	Expirytime int64
	Type       int
	Mac        string
	Iaid       string
	Ipaddr     string
	Prefix     uint32
	Hostname   string
	Clientid   string
}

type RemoteNetworkGetDhcpLeasesReq struct {
	Network     *RemoteNetwork
	Mac         string
	NeedResults int
	Flags       uint32
}

type RemoteNetworkGetDhcpLeasesRes struct {
	Leases []RemoteNetworkDhcpLease
	Ret    uint32
}

type RemoteDomainStatsRecord struct {
	Domain *RemoteDomain
}

type RemoteConnectGetAllDomainStatsReq struct {
	Domains []*RemoteDomain
	Stats   uint32
	Flags   uint32
}

type RemoteDomainEventCallbackAgentLifecycleMsg struct {
	CallbackID int
	Domain     *RemoteDomain

	State  int
	Reason int
}

type RemoteConnectGetAllDomainStatsRes struct {
	RetStats []RemoteDomainStatsRecord
}

type RemoteDomainFsinfo struct {
	Mountpoint string
	Name       string
	Fstype     string
	DevAliases []string
}

type RemoteDomainGetFsinfoReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetFsinfoRes struct {
	Info []RemoteDomainFsinfo
	Ret  uint32
}

type RemoteDomainIpAddr struct {
	Type   int
	Addr   string
	Prefix uint32
}

type RemoteDomainInterface struct {
	Name   string
	Hwaddr string
	Addrs  []RemoteDomainIpAddr
}

type RemoteDomainInterfaceAddressesReq struct {
	Domain *RemoteDomain
	Source uint32
	Flags  uint32
}

type RemoteDomainInterfaceAddressesRes struct {
	Ifaces []*RemoteInterface
}

type RemoteDomainSetUserPasswordReq struct {
	Domain   *RemoteDomain
	User     string
	Password string
	Flags    uint32
}

type RemoteDomainRenameReq struct {
	Domain  *RemoteDomain
	NewName string
	Flags   uint32
}

type RemoteDomainRenameRes struct {
	Retcode int
}

type RemoteDomainEventCallbackMigrationIterationMsg struct {
	CallbackID int
	Domain     *RemoteDomain
	Iteration  int
}

type RemoteDomainEventCallbackJobCompletedMsg struct {
	CallbackID int
	Domain     *RemoteDomain
}

type RemoteDomainMigrateStartPostCopyReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainEventCallbackDeviceRemovalFailedMsg struct {
	CallbackID int
	Domain     *RemoteDomain
	DevAlias   string
}

type RemoteDomainGetGuestVcpusReq struct {
	Domain *RemoteDomain
	Flags  uint32
}

type RemoteDomainGetGuestVcpusRes struct {
}

type RemoteDomainSetGuestVcpusReq struct {
	Domain *RemoteDomain
	Cpumap string
	State  int
	Flags  uint32
}

type RemoteDomainSetVcpuReq struct {
	Domain *RemoteDomain
	Cpumap string
	State  int
	Flags  uint32
}

type RemoteDomainEventCallbackMetadataChangeMsg struct {
	CallbackID int
	Domain     *RemoteDomain
	Type       int
	Nsuri      string
}

type RemoteConnectSecretEventRegisterAnyReq struct {
	EventID int
	Secret  *RemoteSecret
}

type RemoteConnectSecretEventRegisterAnyRes struct {
	CallbackID int
}

type RemoteConnectSecretEventDeregisterAnyReq struct {
	CallbackID int
}

type RemoteSecretEventLifecycleMsg struct {
	CallbackID int
	Secret     *RemoteSecret
	Event      int
	Detail     int
}

type RemoteSecretEventValueChangedMsg struct {
	CallbackID int
	Secret     *RemoteSecret
}

type RemoteDomainSetBlockThresholdReq struct {
	Domain    *RemoteDomain
	Device    string
	Threshold uint64
	Flags     uint32
}
