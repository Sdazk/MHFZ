package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireExchangeShop represents the MSG_MHF_ACQUIRE_EXCHANGE_SHOP
type MsgMhfAcquireExchangeShop struct {
	AckHandle      uint32
	RawDataPayload []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireExchangeShop) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_EXCHANGE_SHOP
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireExchangeShop) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	dataSize := bf.ReadUint16()
	m.RawDataPayload = bf.ReadBytes(uint(dataSize))
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireExchangeShop) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	bf.WriteUint32(m.AckHandle)
	bf.WriteUint16(uint16(len(m.RawDataPayload)))
	bf.WriteBytes(m.RawDataPayload)
	return nil
}
