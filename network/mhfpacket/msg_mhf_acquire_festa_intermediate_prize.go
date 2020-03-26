package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireFestaIntermediatePrize represents the MSG_MHF_ACQUIRE_FESTA_INTERMEDIATE_PRIZE
type MsgMhfAcquireFestaIntermediatePrize struct {
	AckHandle uint32
	Unk0      uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireFestaIntermediatePrize) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_FESTA_INTERMEDIATE_PRIZE
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireFestaIntermediatePrize) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireFestaIntermediatePrize) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
