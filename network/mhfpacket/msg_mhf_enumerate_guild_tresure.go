package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfEnumerateGuildTresure represents the MSG_MHF_ENUMERATE_GUILD_TRESURE
type MsgMhfEnumerateGuildTresure struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateGuildTresure) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_GUILD_TRESURE
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateGuildTresure) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateGuildTresure) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
