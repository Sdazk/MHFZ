package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysRotateObject represents the MSG_SYS_ROTATE_OBJECT
type MsgSysRotateObject struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysRotateObject) Opcode() network.PacketID {
	return network.MSG_SYS_ROTATE_OBJECT
}

// Parse parses the packet from binary
func (m *MsgSysRotateObject) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysRotateObject) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	panic("Not implemented")
}
