package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfPostSeibattle represents the MSG_MHF_POST_SEIBATTLE
type MsgMhfPostSeibattle struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPostSeibattle) Opcode() network.PacketID {
	return network.MSG_MHF_POST_SEIBATTLE
}

// Parse parses the packet from binary
func (m *MsgMhfPostSeibattle) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPostSeibattle) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
