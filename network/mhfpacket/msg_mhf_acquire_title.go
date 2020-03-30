package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireTitle represents the MSG_MHF_ACQUIRE_TITLE
type MsgMhfAcquireTitle struct {
	AckHandle uint32
	//Count uint16
	Unk1 uint16 // Hardcoded 0 in binary
	Unk2 []uint16
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireTitle) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_TITLE
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireTitle) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	count := bf.ReadUint16()
	m.Unk1 = bf.ReadUint16()
	m.Unk2 = make([]uint16, int(count))
	for i := 0; i < int(count); i++ {
		m.Unk2[i] = bf.ReadUint16()
	}
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireTitle) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
