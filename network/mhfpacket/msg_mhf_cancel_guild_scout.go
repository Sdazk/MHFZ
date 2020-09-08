package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfCancelGuildScout represents the MSG_MHF_CANCEL_GUILD_SCOUT
type MsgMhfCancelGuildScout struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfCancelGuildScout) Opcode() network.PacketID {
	return network.MSG_MHF_CANCEL_GUILD_SCOUT
}

// Parse parses the packet from binary
func (m *MsgMhfCancelGuildScout) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfCancelGuildScout) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
