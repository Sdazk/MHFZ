package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysSetStageBinary represents the MSG_SYS_SET_STAGE_BINARY
type MsgSysSetStageBinary struct {
	BinaryType0    uint8
	BinaryType1    uint8  // Index
	StageIDLength  uint8  // <= 0x20
	DataSize       uint16 // <= 0x400
	StageID        string
	RawDataPayload []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysSetStageBinary) Opcode() network.PacketID {
	return network.MSG_SYS_SET_STAGE_BINARY
}

// Parse parses the packet from binary
func (m *MsgSysSetStageBinary) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.BinaryType0 = bf.ReadUint8()
	m.BinaryType1 = bf.ReadUint8()
	m.StageIDLength = bf.ReadUint8()
	m.DataSize = bf.ReadUint16()
	m.StageID = string(bf.ReadBytes(uint(m.StageIDLength)))
	m.RawDataPayload = bf.ReadBytes(uint(m.DataSize))
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgSysSetStageBinary) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	panic("Not implemented")
}
