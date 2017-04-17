package client

import libvirt "github.com/vtolstov/go-libvirt"

// StorageVolume represents a volume as seen by libvirt.
type StorageVolume struct {
	libvirt.RemoteStorageVolume
	l *Libvirt
}

// StorageVolumeDownloadFlags specifies options when performing a volume download.
type StorageVolumeDownloadFlags uint32

const (
	// StorageVolumeDownloadFlagNone default
	StorageVolumeDownloadFlagNone StorageVolumeDownloadFlags = 0

	// StorageVolumeDownloadFlagSparseStream use sparse stream.
	StorageVolumeDownloadFlagSparseStream StorageVolumeDownloadFlags = 1 << iota
)

// StorageVolumeUploadFlags specifies options when performing a volume upload.
type StorageVolumeUploadFlags uint32

const (
	// StorageVolumeUploadFlagNone default
	StorageVolumeUploadFlagNone StorageVolumeUploadFlags = 0

	// StorageVolumeUploadFlagSparseStream use sparse stream.
	StorageVolumeUploadFlagSparseStream StorageVolumeUploadFlags = 1 << iota
)

// StorageVolumeCreateFlags specifies options when performing a volume creation.
type StorageVolumeCreateFlags uint32

const (
	_ StorageVolumeCreateFlags = iota
	// StorageVolumeCreateFlagPreallocMetadata preallocates metadata
	StorageVolumeCreateFlagPreallocMetadata

	// StorageVolumeCreateFlagReflink use btrfs light copy
	StorageVolumeCreateFlagReflink
)

// StorageVolumeDeleteFlags specifies options when performing a volume deletion.
type StorageVolumeDeleteFlags uint32

const (
	// StorageVolumeDeleteFlagNormal delete metadata only (fast)
	StorageVolumeDeleteFlagNormal StorageVolumeDeleteFlags = iota

	// StorageVolumeDeleteFlagZeroes clear all data to zeros (slow)
	StorageVolumeDeleteFlagZeroes

	// StorageVolumeDeleteFlagWithSnapshots force removal of volume, even if in use
	StorageVolumeDeleteFlagWithSnapshots
)

// StorageVolumeResizeFlags specifies options when performing a volume deletion.
type StorageVolumeResizeFlags uint32

const (
	_ StorageVolumeResizeFlags = 0
	// StorageVolumeResizeFlagAllocate force allocation of new size.
	StorageVolumeResizeFlagAllocate StorageVolumeResizeFlags = 1 << (iota - 1)

	// StorageVolumeResizeFlagDelta size is relative to current.
	StorageVolumeResizeFlagDelta

	// StorageVolumeResizeFlagShrink allow decrease in capacity.
	StorageVolumeResizeFlagShrink
)

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
