package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetCafeDurationBonusInfo represents the MSG_MHF_GET_CAFE_DURATION_BONUS_INFO
type MsgMhfGetCafeDurationBonusInfo struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetCafeDurationBonusInfo) Opcode() network.PacketID {
	return network.MSG_MHF_GET_CAFE_DURATION_BONUS_INFO
}

// Parse parses the packet from binary
func (m *MsgMhfGetCafeDurationBonusInfo) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetCafeDurationBonusInfo) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
