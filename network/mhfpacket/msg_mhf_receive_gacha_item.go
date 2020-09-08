package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfReceiveGachaItem represents the MSG_MHF_RECEIVE_GACHA_ITEM
type MsgMhfReceiveGachaItem struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfReceiveGachaItem) Opcode() network.PacketID {
	return network.MSG_MHF_RECEIVE_GACHA_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfReceiveGachaItem) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfReceiveGachaItem) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
