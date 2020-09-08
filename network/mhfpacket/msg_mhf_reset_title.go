package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfResetTitle represents the MSG_MHF_RESET_TITLE
type MsgMhfResetTitle struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfResetTitle) Opcode() network.PacketID {
	return network.MSG_MHF_RESET_TITLE
}

// Parse parses the packet from binary
func (m *MsgMhfResetTitle) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfResetTitle) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
