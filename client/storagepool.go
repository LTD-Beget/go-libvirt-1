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
	*libvirt.RemoteStoragePool
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

	pool := &StoragePool{RemoteStoragePool: res.Pool, l: l}
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

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	pool := &StoragePool{RemoteStoragePool: res.Pool, l: l}
	return pool, nil
}

// Refresh refreshes the storage pool.
func (p *StoragePool) Refresh(flags uint32) error {
	req := libvirt.RemoteStoragePoolRefreshReq{Pool: p.RemoteStoragePool, Flags: uint32(flags)}

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

// ListAllStoragePools returns a list of defined storage pools. Pools are filtered by
// the provided flags. See StoragePools*.
func (l *Libvirt) ListAllStoragePools(flags libvirt.StoragePoolsFlags) ([]*StoragePool, error) {
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
		pools = append(pools, &StoragePool{RemoteStoragePool: pool, l: l})
	}
	return pools, nil
}

// ListAllStorageVolumes returns a list of defined storage volumes.
// the provided flags. See StoragePools*.
func (p *StoragePool) ListAllStorageVolumes(flags uint32) ([]*StorageVolume, error) {
	req := libvirt.RemoteStoragePoolListAllVolumesReq{Pool: p.RemoteStoragePool, NeedResults: 1, Flags: uint32(flags)}
	res := libvirt.RemoteStoragePoolListAllVolumesRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := p.l.send(libvirt.RemoteProcStoragePoolListAllVolumes, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

	var volumes []*StorageVolume
	for _, vol := range res.Vols {
		volumes = append(volumes, &StorageVolume{RemoteStorageVolume: vol, l: p.l})
	}
	return volumes, nil
}

// SetAutostart set autostart for domain.
func (p *StoragePool) SetAutostart(autostart bool) error {
	req := libvirt.RemoteStoragePoolSetAutostartReq{Pool: p.RemoteStoragePool}

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
		Pool:  p.RemoteStoragePool,
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

	vol := &StorageVolume{RemoteStorageVolume: res.Vol, l: p.l}
	return vol, nil
}

// XML dump xml for pool.
func (p *StoragePool) XML(flags libvirt.StorageXmlFlags) (string, error) {
	req := libvirt.RemoteStoragePoolGetXmlDescReq{
		Pool:  p.RemoteStoragePool,
		Flags: uint32(flags)}
	res := libvirt.RemoteStoragePoolGetXmlDescRes{}

	buf, err := encode(&req)
	if err != nil {
		return "", err
	}

	resp, err := p.l.send(libvirt.RemoteProcStoragePoolGetXmlDesc, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

// StorageVolumeCreateXMLFrom creates a volume from another volume.
func (p *StoragePool) StorageVolumeCreateXMLFrom(x string, v *StorageVolume, flags libvirt.StorageVolumeCreateFlags) (*StorageVolume, error) {
	req := libvirt.RemoteStorageVolCreateXmlFromReq{
		Pool:     p.RemoteStoragePool,
		Xml:      x,
		Clonevol: v.RemoteStorageVolume,
		Flags:    uint32(flags)}
	res := libvirt.RemoteStorageVolCreateXmlFromRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := p.l.send(libvirt.RemoteProcStorageVolCreateXmlFrom, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

	vol := &StorageVolume{RemoteStorageVolume: res.Vol, l: p.l}
	return vol, nil
}

// StorageVolumeLookupByName returns a volume as seen by libvirt.
func (p *StoragePool) StorageVolumeLookupByName(name string) (*StorageVolume, error) {
	req := libvirt.RemoteStorageVolLookupByNameReq{
		Pool: p.RemoteStoragePool,
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

	vol := &StorageVolume{RemoteStorageVolume: res.Vol, l: p.l}
	return vol, nil
}
