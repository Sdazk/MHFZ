package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfCheckDailyCafepoint represents the MSG_MHF_CHECK_DAILY_CAFEPOINT
type MsgMhfCheckDailyCafepoint struct {
	AckHandle uint32
	Unk       uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfCheckDailyCafepoint) Opcode() network.PacketID {
	return network.MSG_MHF_CHECK_DAILY_CAFEPOINT
}

// Parse parses the packet from binary
func (m *MsgMhfCheckDailyCafepoint) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfCheckDailyCafepoint) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
