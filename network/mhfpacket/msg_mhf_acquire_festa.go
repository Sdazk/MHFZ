package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireFesta represents the MSG_MHF_ACQUIRE_FESTA
type MsgMhfAcquireFesta struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireFesta) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_FESTA
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireFesta) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireFesta) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
