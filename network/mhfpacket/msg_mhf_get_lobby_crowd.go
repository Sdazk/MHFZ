package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetLobbyCrowd represents the MSG_MHF_GET_LOBBY_CROWD
type MsgMhfGetLobbyCrowd struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetLobbyCrowd) Opcode() network.PacketID {
	return network.MSG_MHF_GET_LOBBY_CROWD
}

// Parse parses the packet from binary
func (m *MsgMhfGetLobbyCrowd) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetLobbyCrowd) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
