package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfReadMercenaryM represents the MSG_MHF_READ_MERCENARY_M
type MsgMhfReadMercenaryM struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfReadMercenaryM) Opcode() network.PacketID {
	return network.MSG_MHF_READ_MERCENARY_M
}

// Parse parses the packet from binary
func (m *MsgMhfReadMercenaryM) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfReadMercenaryM) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
