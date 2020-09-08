package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve19B represents the MSG_SYS_reserve19B
type MsgSysReserve19B struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve19B) Opcode() network.PacketID {
	return network.MSG_SYS_reserve19B
}

// Parse parses the packet from binary
func (m *MsgSysReserve19B) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve19B) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
