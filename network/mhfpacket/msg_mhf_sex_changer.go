package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSexChanger represents the MSG_MHF_SEX_CHANGER
type MsgMhfSexChanger struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSexChanger) Opcode() network.PacketID {
	return network.MSG_MHF_SEX_CHANGER
}

// Parse parses the packet from binary
func (m *MsgMhfSexChanger) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSexChanger) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
