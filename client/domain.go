package client

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"

	"github.com/davecgh/go-xdr/xdr2"
	libvirt "github.com/vtolstov/go-libvirt"
)

// Domain represents a domain as seen by libvirt.
type Domain struct {
	*libvirt.RemoteDomain
	l *Libvirt
}

// Domains returns a list of all domains managed by libvirt.
func (l *Libvirt) ListAllDomains() ([]*Domain, error) {
	req := libvirt.RemoteConnectListAllDomainsReq{
		NeedResults: 1,
		Flags:       3}
	res := libvirt.RemoteConnectListAllDomainsRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.send(libvirt.RemoteProcConnectListAllDomains, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	var domains []*Domain
	for _, d := range res.Domains {
		domains = append(domains, &Domain{RemoteDomain: d, l: l})
	}
	return domains, nil
}

// State returns state of the domain managed by libvirt.
func (d *Domain) State() (libvirt.DomainState, error) {
	req := libvirt.RemoteDomainGetStateReq{
		Domain: d.RemoteDomain,
		Flags:  0}
	res := libvirt.RemoteDomainGetStateRes{}

	buf, err := encode(&req)
	if err != nil {
		return libvirt.DomainStateNoState, err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainGetState, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return libvirt.DomainStateNoState, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return libvirt.DomainStateNoState, decodeError(r.Payload)
	}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return libvirt.DomainStateNoState, err
	}

	return libvirt.DomainState(res.State), nil
}

// Migrate synchronously migrates the domain specified by dom, e.g.,
// 'prod-lb-01', to the destination hypervisor specified by dest, e.g.,
// 'qemu+tcp://example.com/system'. The flags argument determines the
// type of migration and how it will be performed. For more information
// on available migration flags and their meaning, see MigrateFlag*.
func (d *Domain) Migrate(dst string, x string, flags libvirt.DomainMigrateFlags) error {
	_, err := url.Parse(dst)
	if err != nil {
		return err
	}

	req := libvirt.RemoteDomainMigratePerform3Req{
		Domain: d.RemoteDomain,
		Xmlin:  x,
		Uri:    dst,
		Flags:  uint64(flags),
	}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainMigratePerform3Params, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// BlockResize reesize a block device of domain while the domain is running.
func (d *Domain) BlockResize(disk string, size uint64, flags libvirt.DomainBlockResizeFlags) error {
	req := libvirt.RemoteDomainBlockResizeReq{
		Domain: d.RemoteDomain,
		Size:   size,
		Flags:  uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainBlockResize, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// MigrateSetMaxSpeed set the maximum migration bandwidth (in MiB/s) for a
// domain which is being migrated to another host. Specifying a negative value
// results in an essentially unlimited value being provided to the hypervisor.
func (d *Domain) MigrateSetMaxSpeed(bandwidth uint64, flags uint32) error {
	req := libvirt.RemoteDomainMigrateSetMaxSpeedReq{
		Domain:    d.RemoteDomain,
		Bandwidth: bandwidth,
		Flags:     uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}
	fmt.Printf("ZZZ %#+v\n", req)

	resp, err := d.l.send(libvirt.RemoteProcDomainMigrateSetMaxSpeed, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// MigrateSetMaxDowntime set the maximum downtime for a
// domain which is being migrated to another host. Specifying a negative value
// results in an essentially unlimited value being provided to the hypervisor.
func (d *Domain) MigrateSetMaxDowntime(downtime uint64, flags uint32) error {
	req := libvirt.RemoteDomainMigrateSetMaxDowntimeReq{
		Domain:   d.RemoteDomain,
		Downtime: downtime,
		Flags:    uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainMigrateSetMaxDowntime, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Undefine undefines the domain.
// The flags argument allows additional options to be specified such as
// cleaning up snapshot metadata. For more information on available
// flags, see DomainUndefineFlag*.
func (d *Domain) Undefine(flags libvirt.DomainUndefineFlags) error {
	req := libvirt.RemoteDomainUndefineFlagsReq{
		Domain: d.RemoteDomain,
		Flags:  uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainUndefineFlags, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Suspend suspends the domain.
func (d *Domain) Suspend() error {
	req := libvirt.RemoteDomainSuspendReq{Domain: d.RemoteDomain}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainSuspend, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Resume resume domain.
func (d *Domain) Resume() error {
	req := libvirt.RemoteDomainResumeReq{Domain: d.RemoteDomain}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainResume, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// SetAutostart set autostart for domain.
func (d *Domain) SetAutostart(autostart bool) error {
	req := libvirt.RemoteDomainSetAutostartReq{Domain: d.RemoteDomain}

	if autostart {
		req.Autostart = 1
	} else {
		req.Autostart = 0
	}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainSetAutostart, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Destroy destroys the domain.
// The flags argument allows additional options to be specified such as
// allowing a graceful shutdown with SIGTERM than SIGKILL.
// For more information on available flags, see DomainDestroyFlag*.
func (d *Domain) Destroy(flags libvirt.DomainDestroyFlags) error {
	req := libvirt.RemoteDomainDestroyFlagsReq{Domain: d.RemoteDomain, Flags: uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainDestroyFlags, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Reboot reboot the domain.
// The flags argument allows additional options to be specified.
// For more information on available flags, see DomainRebootFlags*.
func (d *Domain) Reboot(flags libvirt.DomainRebootFlags) error {
	req := libvirt.RemoteDomainRebootReq{Domain: d.RemoteDomain, Flags: uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainReboot, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Shutdown reboot the domain.
// The flags argument allows additional options to be specified.
// For more information on available flags, see DomainShutdownFlags*.
func (d *Domain) Shutdown(flags libvirt.DomainShutdownFlags) error {
	req := libvirt.RemoteDomainShutdownFlagsReq{Domain: d.RemoteDomain, Flags: uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainShutdownFlags, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// XML returns a domain's raw XML definition, akin to `virsh dumpxml <domain>`.
// See DomainXMLFlag* for optional flags.
func (d *Domain) XML(flags libvirt.DomainXMLFlags) (string, error) {
	req := libvirt.RemoteDomainGetXmlDescReq{Domain: d.RemoteDomain, Flags: uint32(flags)}
	res := libvirt.RemoteDomainGetXmlDescRes{}

	buf, err := encode(&req)
	if err != nil {
		return "", err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainGetXmlDesc, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return "", err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return "", decodeError(r.Payload)
	}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return "", err
	}

	return res.Xml, nil
}

// DefineXML defines a domain, but does not start it.
func (l *Libvirt) DomainDefineXML(x string, flags libvirt.DomainDefineXMLFlags) error {
	req := libvirt.RemoteDomainDefineXmlFlagsReq{Xml: x, Flags: uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := l.send(libvirt.RemoteProcDomainDefineXmlFlags, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Create start defined domain.
func (d *Domain) Create() error {
	req := libvirt.RemoteDomainCreateReq{Domain: d.RemoteDomain}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainCreate, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// DomainCreateXML start domain based on xml.
func (l *Libvirt) DomainCreateXML(x string, flags libvirt.DomainCreateFlags) error {
	req := libvirt.RemoteDomainCreateXmlReq{XML: x, Flags: uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := l.send(libvirt.RemoteProcDomainCreateXml, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// DomainLookupByName returns a domain as seen by libvirt.
func (l *Libvirt) DomainLookupByName(name string) (*Domain, error) {
	req := libvirt.RemoteDomainLookupByNameReq{Name: name}
	res := libvirt.RemoteDomainLookupByNameRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.send(libvirt.RemoteProcDomainLookupByName, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))

	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}
	fmt.Printf("OOO %+#v\n", res)

	return &Domain{RemoteDomain: res.Domain, l: l}, nil
}

// DomainLookupByUUID returns a domain as seen by libvirt.
func (l *Libvirt) DomainLookupByUUID(uuid string) (*Domain, error) {
	req := libvirt.RemoteDomainLookupByUuidReq{}
	res := libvirt.RemoteDomainLookupByUuidRes{}

	_, err := hex.Decode(req.UUID[:], []byte(strings.Replace(uuid, "-", "", -1)))
	if err != nil {
		return nil, err
	}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.send(libvirt.RemoteProcDomainLookupByUuid, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))

	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &Domain{RemoteDomain: res.Domain, l: l}, nil
}
