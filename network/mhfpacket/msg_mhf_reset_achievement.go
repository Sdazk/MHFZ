package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfResetAchievement represents the MSG_MHF_RESET_ACHIEVEMENT
type MsgMhfResetAchievement struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfResetAchievement) Opcode() network.PacketID {
	return network.MSG_MHF_RESET_ACHIEVEMENT
}

// Parse parses the packet from binary
func (m *MsgMhfResetAchievement) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfResetAchievement) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
