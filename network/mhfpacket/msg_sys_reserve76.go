package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve76 represents the MSG_SYS_reserve76
type MsgSysReserve76 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve76) Opcode() network.PacketID {
	return network.MSG_SYS_reserve76
}

// Parse parses the packet from binary
func (m *MsgSysReserve76) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve76) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
