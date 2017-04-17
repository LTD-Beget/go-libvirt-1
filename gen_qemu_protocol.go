package libvirt

const (
	QemuProgram         = 0x20008087
	QemuProtocolVersion = 1
)

const ()

type QemuDomainMonitorCommandReq struct {
	Domain *RemoteDomain
	Cmd    string
	Flags  uint32
}

type QemuDomainMonitorCommandRes struct {
	Result string
}

type QemuDomainAttachReq struct {
	PidValue uint32
	Flags    uint32
}

type QemuDomainAttachRes struct {
	Domain *RemoteDomain
}

type QemuDomainAgentCommandReq struct {
	Domain  *RemoteDomain
	Cmd     string
	Timeout int
	Flags   uint32
}

type QemuDomainAgentCommandRes struct {
	Result string
}

type QemuConnectDomainMonitorEventRegisterReq struct {
	Domain *RemoteDomain
	Event  string
	Flags  uint32
}

type QemuConnectDomainMonitorEventRegisterRes struct {
	CallbackID int
}

type QemuConnectDomainMonitorEventDeregisterReq struct {
	CallbackID int
}

type QemuDomainMonitorEventMsg struct {
	CallbackID int
	Domain     *RemoteDomain
	Event      string
	Seconds    int64
	Micros     uint32
	Details    string
}
