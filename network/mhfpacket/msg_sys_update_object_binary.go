package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysUpdateObjectBinary represents the MSG_SYS_UPDATE_OBJECT_BINARY
type MsgSysUpdateObjectBinary struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysUpdateObjectBinary) Opcode() network.PacketID {
	return network.MSG_SYS_UPDATE_OBJECT_BINARY
}

// Parse parses the packet from binary
func (m *MsgSysUpdateObjectBinary) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysUpdateObjectBinary) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	panic("Not implemented")
}
