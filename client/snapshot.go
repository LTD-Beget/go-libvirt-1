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

// SnapshotCreateXML creates a new snapshot of domain based on xml.
func (d *Domain) SnapshotCreateXML(x []byte, flags DomainSnapshotCreateFlags) (*Snapshot, error) {
	res := libvirt.RemoteDomainSnapshotCreateXmlRes{}
	req := libvirt.RemoteDomainSnapshotCreateXmlReq{
		Domain: &libvirt.RemoteDomain{Name: d.Name, UUID: d.UUID, ID: d.ID},
		XML:    string(x),
		Flags:  uint32(flags),
	}

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

func (d *Domain) ListAllSnapshots(flags DomainSnapshotListFlags) ([]*Snapshot, error) {
	res := libvirt.RemoteDomainListAllSnapshotsRes{}
	req := libvirt.RemoteDomainListAllSnapshotsReq{
		Domain: &libvirt.RemoteDomain{Name: d.Name, UUID: d.UUID, ID: d.ID},
		Flags:  uint32(flags),
	}

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
	res := libvirt.RemoteDomainSnapshotLookupByNameRes{}
	req := libvirt.RemoteDomainSnapshotLookupByNameReq{
		Domain: &libvirt.RemoteDomain{Name: d.Name, UUID: d.UUID, ID: d.ID},
		Name:   name,
		Flags:  uint32(0),
	}

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

func (s *Snapshot) Delete(flags DomainSnapshotDeleteFlags) error {
	req := libvirt.RemoteDomainSnapshotDeleteReq{
		Snapshot: &libvirt.RemoteDomainSnapshot{Name: s.Name, Domain: s.Domain},
		Flags:    uint32(flags),
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

func (s *Snapshot) Revert(flags DomainSnapshotRevertFlags) error {
	req := libvirt.RemoteDomainRevertToSnapshotReq{
		Snapshot: &libvirt.RemoteDomainSnapshot{Name: s.Name, Domain: s.Domain},
		Flags:    uint32(flags),
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
