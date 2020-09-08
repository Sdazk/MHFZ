package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfEnumerateUnionItem represents the MSG_MHF_ENUMERATE_UNION_ITEM
type MsgMhfEnumerateUnionItem struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateUnionItem) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_UNION_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateUnionItem) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateUnionItem) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
