package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysHideObject represents the MSG_SYS_HIDE_OBJECT
type MsgSysHideObject struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysHideObject) Opcode() network.PacketID {
	return network.MSG_SYS_HIDE_OBJECT
}

// Parse parses the packet from binary
func (m *MsgSysHideObject) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysHideObject) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
