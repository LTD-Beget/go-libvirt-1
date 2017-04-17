package client

import (
	"bytes"

	"github.com/davecgh/go-xdr/xdr2"
	libvirt "github.com/vtolstov/go-libvirt"
)

// Network represents a network as seen by libvirt.
type Network struct {
	libvirt.RemoteNetwork
	l *Libvirt
}

// NetworkLookupByName returns the network associated with the provided name.
// An error is returned if the requested network is not found.
func (l *Libvirt) NetworkLookupByName(name string) (*Network, error) {
	req := struct {
		Name string
	}{
		Name: name,
	}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.send(libvirt.RemoteProcNetworkLookupByName, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return nil, decodeError(r.Payload)
	}

	result := struct {
		Net Network
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}

	result.Net.l = l
	return &result.Net, nil
}

// SetAutostart set autostart for network.
func (n *Network) SetAutostart(autostart bool) error {
	payload := struct {
		Network   Network
		Autostart int32
	}{}

	payload.Network = *n
	if autostart {
		payload.Autostart = 1
	} else {
		payload.Autostart = 0
	}

	buf, err := encode(&payload)
	if err != nil {
		return err
	}

	resp, err := n.l.send(libvirt.RemoteProcNetworkSetAutostart, 0, libvirt.MessageTypeCall, libvirt.RemoteProgram, libvirt.MessageStatusOK, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Header.Status != libvirt.MessageStatusOK {
		return decodeError(r.Payload)
	}

	return nil
}
