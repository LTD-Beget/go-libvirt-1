// Code generated by "stringer -type=MessageType"; DO NOT EDIT.

package libvirt

import "fmt"

const _MessageType_name = "MessageTypeCallMessageTypeReplyMessageTypeMessageMessageTypeStreamMessageTypeCallWithFDsMessageTypeReplyWithFDsMessageTypeStreamSkip"

var _MessageType_index = [...]uint8{0, 15, 31, 49, 66, 88, 111, 132}

func (i MessageType) String() string {
	if i >= MessageType(len(_MessageType_index)-1) {
		return fmt.Sprintf("MessageType(%d)", i)
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}
