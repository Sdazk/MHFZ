package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve05 represents the MSG_SYS_reserve05
type MsgSysReserve05 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve05) Opcode() network.PacketID {
	return network.MSG_SYS_reserve05
}

// Parse parses the packet from binary
func (m *MsgSysReserve05) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve05) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
