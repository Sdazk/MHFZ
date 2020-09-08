package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireUdItem represents the MSG_MHF_ACQUIRE_UD_ITEM
type MsgMhfAcquireUdItem struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireUdItem) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_UD_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireUdItem) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireUdItem) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
