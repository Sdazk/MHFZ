package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetUdTacticsRanking represents the MSG_MHF_GET_UD_TACTICS_RANKING
type MsgMhfGetUdTacticsRanking struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetUdTacticsRanking) Opcode() network.PacketID {
	return network.MSG_MHF_GET_UD_TACTICS_RANKING
}

// Parse parses the packet from binary
func (m *MsgMhfGetUdTacticsRanking) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetUdTacticsRanking) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
