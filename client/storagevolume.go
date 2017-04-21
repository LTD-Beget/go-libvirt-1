package client

import (
	"bytes"

	"github.com/davecgh/go-xdr/xdr2"
	libvirt "github.com/vtolstov/go-libvirt"
)

type StorageVolumeInfo struct {
	libvirt.RemoteStorageVolGetInfoFlagsRes
}

// StorageVolume represents a volume as seen by libvirt.
type StorageVolume struct {
	*libvirt.RemoteStorageVolume
	l *Libvirt
}

// Resize resize a volume.
func (v *StorageVolume) Resize(size uint64, flags libvirt.StorageVolumeResizeFlags) error {
	req := libvirt.RemoteStorageVolResizeReq{
		Vol:      v.RemoteStorageVolume,
		Capacity: size,
		Flags:    uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolResize, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Delete deletes a volume.
func (v *StorageVolume) Delete(flags libvirt.StorageVolumeDeleteFlags) error {
	req := libvirt.RemoteStorageVolDeleteReq{
		Vol:   v.RemoteStorageVolume,
		Flags: uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolDelete, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// WipePattern wipes volume using specified alg.
func (v *StorageVolume) WipePattern(alg libvirt.StorageVolumeWipeAlgFlags, flags uint32) error {
	req := libvirt.RemoteStorageVolWipePatternReq{
		Vol:       v.RemoteStorageVolume,
		Algorithm: uint32(alg),
		Flags:     uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolWipePattern, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Wipe wipes volume.
func (v *StorageVolume) Wipe(flags uint32) error {
	req := libvirt.RemoteStorageVolWipeReq{
		Vol:   v.RemoteStorageVolume,
		Flags: uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolWipe, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// XML dump xml for volume.
func (v *StorageVolume) XML(flags uint32) (string, error) {
	req := libvirt.RemoteStorageVolGetXmlDescReq{
		Vol:   v.RemoteStorageVolume,
		Flags: uint32(flags)}
	res := libvirt.RemoteStorageVolGetXmlDescRes{}

	buf, err := encode(&req)
	if err != nil {
		return "", err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolGetXmlDesc, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

// Download downloads a volume.
func (v *StorageVolume) Download(offset uint64, length uint64, flags libvirt.StorageVolumeDownloadFlags) (*Stream, error) {
	req := libvirt.RemoteStorageVolDownloadReq{
		Vol:    v.RemoteStorageVolume,
		Offset: offset,
		Length: length,
		Flags:  uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolDownload, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	s, err := v.l.StreamNew()
	if err != nil {
		return nil, err
	}

	s.serial = r.Header.Serial
	s.procedure = libvirt.RemoteProcStorageVolDownload
	v.l.addStream(r.Header.Serial, s)

	return s, nil
}

// Upload uploads a volume.
func (v *StorageVolume) Upload(offset uint64, length uint64, flags libvirt.StorageVolumeUploadFlags) (*Stream, error) {
	req := libvirt.RemoteStorageVolUploadReq{
		Vol:    v.RemoteStorageVolume,
		Offset: offset,
		Length: length,
		Flags:  uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolUpload, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	s, err := v.l.StreamNew()
	if err != nil {
		return nil, err
	}

	s.serial = r.Header.Serial
	s.procedure = libvirt.RemoteProcStorageVolUpload
	v.l.addStream(r.Header.Serial, s)

	return s, nil
}

// Info bout volume.
func (v *StorageVolume) Info(flags libvirt.StorageVolumeInfoFlags) (*StorageVolumeInfo, error) {
	req := libvirt.RemoteStorageVolGetInfoFlagsReq{Vol: v.RemoteStorageVolume, Flags: uint32(flags)}
	res := libvirt.RemoteStorageVolGetInfoFlagsRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolGetInfoFlags, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

	return &StorageVolumeInfo{RemoteStorageVolGetInfoFlagsRes: res}, nil
}
