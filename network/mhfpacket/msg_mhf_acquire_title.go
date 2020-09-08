package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireTitle represents the MSG_MHF_ACQUIRE_TITLE
type MsgMhfAcquireTitle struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireTitle) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_TITLE
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireTitle) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireTitle) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
