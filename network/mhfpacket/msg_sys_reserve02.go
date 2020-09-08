package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve02 represents the MSG_SYS_reserve02
type MsgSysReserve02 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve02) Opcode() network.PacketID {
	return network.MSG_SYS_reserve02
}

// Parse parses the packet from binary
func (m *MsgSysReserve02) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve02) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
