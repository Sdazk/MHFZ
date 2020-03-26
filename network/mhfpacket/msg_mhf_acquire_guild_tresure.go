package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireGuildTresure represents the MSG_MHF_ACQUIRE_GUILD_TRESURE
type MsgMhfAcquireGuildTresure struct {
	AckHandle uint32
	Unk0      uint32
	Unk1      bool
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireGuildTresure) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_GUILD_TRESURE
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireGuildTresure) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint32()
	m.Unk1 = bf.ReadBool()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireGuildTresure) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
