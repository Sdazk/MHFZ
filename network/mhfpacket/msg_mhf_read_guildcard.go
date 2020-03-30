package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfReadGuildcard represents the MSG_MHF_READ_GUILDCARD
type MsgMhfReadGuildcard struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfReadGuildcard) Opcode() network.PacketID {
	return network.MSG_MHF_READ_GUILDCARD
}

// Parse parses the packet from binary
func (m *MsgMhfReadGuildcard) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfReadGuildcard) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
