package client

import (
	"bytes"

	"github.com/davecgh/go-xdr/xdr2"
	libvirt "github.com/vtolstov/go-libvirt"
)

// Snapshot represents a Snapshot as seen by libvirt.
type Snapshot struct {
	libvirt.RemoteDomainSnapshot
	l *Libvirt
}

// SnapshotCreateXML creates a new snapshot of domain based on xml.
func (d *Domain) SnapshotCreateXML(x string, flags libvirt.DomainSnapshotCreateFlags) (*Snapshot, error) {
	req := libvirt.RemoteDomainSnapshotCreateXmlReq{
		Domain: &libvirt.RemoteDomain{
			Name: d.Name,
			UUID: d.UUID,
			ID:   d.ID,
		},
		XML:   x,
		Flags: uint32(flags)}
	res := libvirt.RemoteDomainSnapshotCreateXmlRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainSnapshotCreateXml, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

	return &Snapshot{RemoteDomainSnapshot: *res.Snapshot, l: d.l}, nil
}

func (d *Domain) ListAllSnapshots(flags libvirt.DomainSnapshotListFlags) ([]*Snapshot, error) {
	req := libvirt.RemoteDomainListAllSnapshotsReq{
		Domain: &libvirt.RemoteDomain{
			Name: d.Name,
			UUID: d.UUID,
			ID:   d.ID,
		},
		Flags: uint32(flags),
	}
	res := libvirt.RemoteDomainListAllSnapshotsRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainListAllSnapshots, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

	snaps := make([]*Snapshot, res.Ret)
	for idx, snap := range res.Snapshots {
		snaps[idx] = &Snapshot{RemoteDomainSnapshot: *snap, l: d.l}
	}

	return snaps, nil
}

func (d *Domain) SnapshotLookupByName(name string, flags uint32) (*Snapshot, error) {
	req := libvirt.RemoteDomainSnapshotLookupByNameReq{
		Domain: &libvirt.RemoteDomain{
			Name: d.Name,
			UUID: d.UUID,
			ID:   d.ID,
		},
		Name:  name,
		Flags: uint32(0),
	}
	res := libvirt.RemoteDomainSnapshotLookupByNameRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := d.l.send(libvirt.RemoteProcDomainSnapshotLookupByName, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

	return &Snapshot{RemoteDomainSnapshot: *res.Snapshot, l: d.l}, nil
}

func (s *Snapshot) Delete(flags libvirt.DomainSnapshotDeleteFlags) error {
	req := libvirt.RemoteDomainSnapshotDeleteReq{
		Snapshot: &libvirt.RemoteDomainSnapshot{
			Name:   s.Name,
			Domain: s.Domain,
		},
		Flags: uint32(flags),
	}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := s.l.send(libvirt.RemoteProcDomainSnapshotDelete, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

func (s *Snapshot) Revert(flags libvirt.DomainSnapshotRevertFlags) error {
	req := libvirt.RemoteDomainRevertToSnapshotReq{
		Snapshot: &libvirt.RemoteDomainSnapshot{
			Name:   s.Name,
			Domain: s.Domain,
		},
		Flags: uint32(flags),
	}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := s.l.send(libvirt.RemoteProcDomainRevertToSnapshot, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}
