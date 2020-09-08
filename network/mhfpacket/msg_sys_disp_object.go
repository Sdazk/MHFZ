package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysDispObject represents the MSG_SYS_DISP_OBJECT
type MsgSysDispObject struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysDispObject) Opcode() network.PacketID {
	return network.MSG_SYS_DISP_OBJECT
}

// Parse parses the packet from binary
func (m *MsgSysDispObject) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysDispObject) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
