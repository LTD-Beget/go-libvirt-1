package client

import (
	"bytes"

	"github.com/davecgh/go-xdr/xdr2"
	libvirt "github.com/vtolstov/go-libvirt"
)

// Secret represents a secret managed by the libvirt daemon.
type Secret struct {
	*libvirt.RemoteSecret
	l *Libvirt
}

// Secrets returns all secrets managed by the libvirt daemon.
func (l *Libvirt) Secrets() ([]*Secret, error) {
	req := libvirt.RemoteConnectListAllSecretsReq{
		NeedResults: 1,
		Flags:       0}
	res := libvirt.RemoteConnectListAllSecretsRes{}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.send(libvirt.RemoteProcConnectListAllSecrets, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
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

	var secrets []*Secret
	for _, secret := range res.Secrets {
		secrets = append(secrets, &Secret{l: l, RemoteSecret: secret})
	}
	return secrets, nil
}
