package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireItem represents the MSG_MHF_ACQUIRE_ITEM
type MsgMhfAcquireItem struct {
	AckHandle uint32
	Unk0      uint32
	Unk1      uint16 // Hardcoded 0 in binary.
	// count uint32
	Unk2 []uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireItem) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireItem) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint32()
	m.Unk1 = bf.ReadUint16()
	count := bf.ReadUint16()
	m.Unk2 = make([]uint32, uint(count))
	for i := 0; i < int(count); i++ {
		m.Unk2[i] = bf.ReadUint32()
	}
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireItem) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
