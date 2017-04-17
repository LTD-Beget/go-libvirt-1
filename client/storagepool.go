package client

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/davecgh/go-xdr/xdr2"
	libvirt "github.com/vtolstov/go-libvirt"
)

// StoragePool represents a storage pool as seen by libvirt.
type StoragePool struct {
	libvirt.RemoteStoragePool
	l *Libvirt
}

// StoragePoolsFlags specifies storage pools to list.
type StoragePoolsFlags uint32

// These flags come in groups; if all bits from a group are 0,
// then that group is not used to filter results.
const (
	StoragePoolsFlagInactive = 1 << iota
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

// StoragePoolLookupByName returns the storage pool associated with the provided name.
// An error is returned if the requested storage pool is not found.
func (l *Libvirt) StoragePoolLookupByName(name string) (*StoragePool, error) {
	req := struct {
		Name string
	}{
		Name: name,
	}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.send(libvirt.RemoteProcStoragePoolLookupByName, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	result := struct {
		Pool StoragePool
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}

	result.Pool.l = l
	return &result.Pool, nil
}

// StoragePoolLookupByUUID returns the storage pool associated with the provided uuid.
// An error is returned if the requested storage pool is not found.
func (l *Libvirt) StoragePoolLookupByUUID(uuid string) (*StoragePool, error) {
	req := struct {
		UUID libvirt.UUID
	}{}

	_, err := hex.Decode(req.UUID[:], []byte(strings.Replace(uuid, "-", "", -1)))
	if err != nil {
		return nil, err
	}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.send(libvirt.RemoteProcStoragePoolLookupByUuid, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	result := struct {
		Pool StoragePool
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}

	result.Pool.l = l
	return &result.Pool, nil
}

// Refresh refreshes the storage pool.
func (p *StoragePool) Refresh(flags uint32) error {
	req := struct {
		Pool  StoragePool
		Flags uint32
	}{
		Pool:  *p,
		Flags: flags, // unused per libvirt source, callers should pass 0
	}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := p.l.send(libvirt.RemoteProcStoragePoolRefresh, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// StoragePools returns a list of defined storage pools. Pools are filtered by
// the provided flags. See StoragePools*.
func (l *Libvirt) StoragePools(flags StoragePoolsFlags) ([]StoragePool, error) {
	req := struct {
		NeedResults uint32
		Flags       StoragePoolsFlags
	}{
		NeedResults: 1,
		Flags:       flags,
	}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.send(libvirt.RemoteProcConnectListAllStoragePools, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	result := struct {
		Pools []StoragePool
		Count uint32
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}

	for _, p := range result.Pools {
		p.l = l
	}
	return result.Pools, nil
}

// SetAutostart set autostart for domain.
func (p *StoragePool) SetAutostart(autostart bool) error {
	payload := struct {
		Pool      StoragePool
		Autostart int32
	}{}

	payload.Pool = *p
	if autostart {
		payload.Autostart = 1
	} else {
		payload.Autostart = 0
	}

	buf, err := encode(&payload)
	if err != nil {
		return err
	}

	resp, err := p.l.send(libvirt.RemoteProcStoragePoolSetAutostart, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// StorageVolumeCreateXML creates a volume.
func (p *StoragePool) StorageVolumeCreateXML(x []byte, flags StorageVolumeCreateFlags) (*StorageVolume, error) {
	payload := struct {
		Pool  StoragePool
		XML   []byte
		Flags StorageVolumeCreateFlags
	}{
		Pool:  *p,
		XML:   x,
		Flags: flags,
	}

	buf, err := encode(&payload)
	if err != nil {
		return nil, err
	}

	resp, err := p.l.send(libvirt.RemoteProcStorageVolCreateXml, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	result := struct {
		Volume StorageVolume
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	result.Volume.l = p.l

	return &result.Volume, nil
}

// StorageVolumeLookupByName returns a volume as seen by libvirt.
func (p *StoragePool) StorageVolumeLookupByName(name string) (*StorageVolume, error) {
	payload := struct {
		Pool StoragePool
		Name string
	}{
		Pool: *p,
		Name: name,
	}

	buf, err := encode(&payload)
	if err != nil {
		return nil, err
	}

	resp, err := p.l.send(libvirt.RemoteProcStorageVolLookupByName, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	result := struct {
		Volume StorageVolume
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}

	result.Volume.l = p.l
	return &result.Volume, nil
}
