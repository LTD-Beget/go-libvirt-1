package client

import libvirt "github.com/vtolstov/go-libvirt"

// StorageVolume represents a volume as seen by libvirt.
type StorageVolume struct {
	libvirt.RemoteStorageVolume
	l *Libvirt
}

// Resize resize a volume.
func (v *StorageVolume) Resize(size uint64, flags StorageVolumeResizeFlags) error {
	payload := struct {
		Vol   StorageVolume
		Size  uint64
		Flags StorageVolumeResizeFlags
	}{
		Vol:   *v,
		Size:  size,
		Flags: flags,
	}

	buf, err := encode(&payload)
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
func (v *StorageVolume) Delete(flags StorageVolumeDeleteFlags) error {
	payload := struct {
		Vol   StorageVolume
		Flags StorageVolumeDeleteFlags
	}{
		Vol:   *v,
		Flags: flags,
	}

	buf, err := encode(&payload)
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
func (v *StorageVolume) Download(s *Stream, offset uint64, length uint64, flags StorageVolumeDownloadFlags) error {
	payload := struct {
		Vol    StorageVolume
		Offset uint64
		Length uint64
		Flags  StorageVolumeDownloadFlags
	}{
		Vol:    *v,
		Offset: offset,
		Length: length,
		Flags:  flags,
	}

	buf, err := encode(&payload)
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
func (v *StorageVolume) Upload(s *Stream, offset uint64, length uint64, flags StorageVolumeUploadFlags) error {
	payload := struct {
		Vol    StorageVolume
		Offset uint64
		Length uint64
		Flags  StorageVolumeUploadFlags
	}{
		Vol:    *v,
		Offset: offset,
		Length: length,
		Flags:  flags,
	}

	buf, err := encode(&payload)
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
