package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve07 represents the MSG_SYS_reserve07
type MsgSysReserve07 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve07) Opcode() network.PacketID {
	return network.MSG_SYS_reserve07
}

// Parse parses the packet from binary
func (m *MsgSysReserve07) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve07) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
