package libvirt

const (
	LxcProgram         = 0x00068000
	LxcProtocolVersion = 1
)

const ()

type LxcDomainOpenNamespaceReq struct {
	Domain *RemoteDomain
	Flags  uint32
}
