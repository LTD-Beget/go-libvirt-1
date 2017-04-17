package client

import libvirt "github.com/vtolstov/go-libvirt"

// StorageVolume represents a volume as seen by libvirt.
type StorageVolume struct {
	libvirt.RemoteStorageVolume
	l *Libvirt
}

// Resize resize a volume.
func (v *StorageVolume) Resize(size uint64, flags libvirt.StorageVolumeResizeFlags) error {
	req := libvirt.RemoteStorageVolResizeReq{
		Vol: &libvirt.RemoteStorageVolume{
			Pool: v.Pool,
			Name: v.Name,
			Key:  v.Key,
		},
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
		Vol: &libvirt.RemoteStorageVolume{
			Pool: v.Pool,
			Name: v.Name,
			Key:  v.Key,
		},
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

// Download downloads a volume.
func (v *StorageVolume) Download(s *Stream, offset uint64, length uint64, flags libvirt.StorageVolumeDownloadFlags) error {
	req := libvirt.RemoteStorageVolDownloadReq{
		Vol: &libvirt.RemoteStorageVolume{
			Pool: v.Pool,
			Name: v.Name,
			Key:  v.Key,
		},
		Offset: offset,
		Length: length,
		Flags:  uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolDownload, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	s.serial = r.Header.Serial
	s.procedure = libvirt.RemoteProcStorageVolDownload
	v.l.addStream(r.Header.Serial, s)

	return nil
}

// Upload uploads a volume.
func (v *StorageVolume) Upload(s *Stream, offset uint64, length uint64, flags libvirt.StorageVolumeUploadFlags) error {
	req := libvirt.RemoteStorageVolUploadReq{
		Vol: &libvirt.RemoteStorageVolume{
			Pool: v.Pool,
			Name: v.Name,
			Key:  v.Key,
		},
		Offset: offset,
		Length: length,
		Flags:  uint32(flags)}

	buf, err := encode(&req)
	if err != nil {
		return err
	}

	resp, err := v.l.send(libvirt.RemoteProcStorageVolUpload, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	s.serial = r.Header.Serial
	s.procedure = libvirt.RemoteProcStorageVolUpload
	v.l.addStream(r.Header.Serial, s)

	return nil
}
