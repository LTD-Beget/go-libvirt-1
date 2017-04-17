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

// StoragePoolLookupByName returns the storage pool associated with the provided name.
// An error is returned if the requested storage pool is not found.
func (l *Libvirt) StoragePoolLookupByName(name string) (*StoragePool, error) {
	req := libvirt.RemoteStoragePoolLookupByNameReq{Name: name}
	res := libvirt.RemoteStoragePoolLookupByNameRes{}

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

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	pool := &StoragePool{RemoteStoragePool: *res.Pool, l: l}
	return pool, nil
}

// StoragePoolLookupByUUID returns the storage pool associated with the provided uuid.
// An error is returned if the requested storage pool is not found.
func (l *Libvirt) StoragePoolLookupByUUID(uuid string) (*StoragePool, error) {
	req := libvirt.RemoteStoragePoolLookupByUuidReq{}
	res := libvirt.RemoteStoragePoolLookupByUuidRes{}

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

	pool := &StoragePool{RemoteStoragePool: *res.Pool, l: l}
	return pool, nil
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
func (l *Libvirt) StoragePools(flags libvirt.StoragePoolsFlags) ([]*StoragePool, error) {
	req := libvirt.RemoteConnectListAllStoragePoolsReq{NeedResults: 1, Flags: uint32(flags)}
	res := libvirt.RemoteConnectListAllStoragePoolsRes{}

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

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	var pools []*StoragePool
	for _, pool := range res.Pools {
		pools = append(pools, &StoragePool{RemoteStoragePool: *pool, l: l})
	}
	return pools, nil
}

// SetAutostart set autostart for domain.
func (p *StoragePool) SetAutostart(autostart bool) error {
	req := libvirt.RemoteStoragePoolSetAutostartReq{
		Pool: &libvirt.RemoteStoragePool{
			Name: p.Name,
			UUID: p.UUID,
		}}

	if autostart {
		req.Autostart = 1
	} else {
		req.Autostart = 0
	}

	buf, err := encode(&req)
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
func (p *StoragePool) StorageVolumeCreateXML(x string, flags libvirt.StorageVolumeCreateFlags) (*StorageVolume, error) {
	req := libvirt.RemoteStorageVolCreateXmlReq{
		Pool: &libvirt.RemoteStoragePool{
			Name: p.Name,
			UUID: p.UUID,
		},
		Xml:   x,
		Flags: uint32(flags)}
	res := libvirt.RemoteStorageVolCreateXmlRes{}

	buf, err := encode(&req)
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

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	vol := &StorageVolume{RemoteStorageVolume: *res.Vol, l: p.l}
	return vol, nil
}

// StorageVolumeLookupByName returns a volume as seen by libvirt.
func (p *StoragePool) StorageVolumeLookupByName(name string) (*StorageVolume, error) {
	req := libvirt.RemoteStorageVolLookupByNameReq{
		Pool: &libvirt.RemoteStoragePool{
			Name: p.Name,
			UUID: p.UUID,
		},
		Name: name}
	res := libvirt.RemoteStorageVolLookupByNameRes{}

	buf, err := encode(&req)
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

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	vol := &StorageVolume{RemoteStorageVolume: *res.Vol, l: p.l}
	return vol, nil
}
