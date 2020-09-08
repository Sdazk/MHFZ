package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateGuildItem represents the MSG_MHF_UPDATE_GUILD_ITEM
type MsgMhfUpdateGuildItem struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateGuildItem) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_GUILD_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateGuildItem) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateGuildItem) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
