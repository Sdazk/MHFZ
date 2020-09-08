package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysAuthData represents the MSG_SYS_AUTH_DATA
type MsgSysAuthData struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysAuthData) Opcode() network.PacketID {
	return network.MSG_SYS_AUTH_DATA
}

// Parse parses the packet from binary
func (m *MsgSysAuthData) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysAuthData) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
