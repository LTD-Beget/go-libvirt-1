package client

import (
	"bytes"

	"github.com/davecgh/go-xdr/xdr2"
	libvirt "github.com/vtolstov/go-libvirt"
)

// Network represents a network as seen by libvirt.
type Network struct {
	*libvirt.RemoteNetwork
	l *Libvirt
}

// NetworkLookupByName returns the network associated with the provided name.
// An error is returned if the requested network is not found.
func (l *Libvirt) NetworkLookupByName(name string) (*Network, error) {
	req := libvirt.RemoteNetworkLookupByNameReq{Name: name}
	res := libvirt.RemoteNetworkLookupByNameRes{}

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

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &Network{RemoteNetwork: res.Network, l: l}, nil
}

// SetAutostart set autostart for network.
func (n *Network) SetAutostart(autostart bool) error {
	req := libvirt.RemoteNetworkSetAutostartReq{Network: n.RemoteNetwork}
	if autostart {
		req.Autostart = 1
	} else {
		req.Autostart = 0
	}

	buf, err := encode(&req)
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
