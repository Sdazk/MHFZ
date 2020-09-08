package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSetLoginwindow represents the MSG_MHF_SET_LOGINWINDOW
type MsgMhfSetLoginwindow struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSetLoginwindow) Opcode() network.PacketID {
	return network.MSG_MHF_SET_LOGINWINDOW
}

// Parse parses the packet from binary
func (m *MsgMhfSetLoginwindow) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSetLoginwindow) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
