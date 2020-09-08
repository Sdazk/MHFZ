package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfEnumerateGuildItem represents the MSG_MHF_ENUMERATE_GUILD_ITEM
type MsgMhfEnumerateGuildItem struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateGuildItem) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_GUILD_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateGuildItem) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateGuildItem) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
