package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateBeatLevel represents the MSG_MHF_UPDATE_BEAT_LEVEL
type MsgMhfUpdateBeatLevel struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateBeatLevel) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_BEAT_LEVEL
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateBeatLevel) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateBeatLevel) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
