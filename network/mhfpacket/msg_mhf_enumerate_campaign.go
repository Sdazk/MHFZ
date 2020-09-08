package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfEnumerateCampaign represents the MSG_MHF_ENUMERATE_CAMPAIGN
type MsgMhfEnumerateCampaign struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateCampaign) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_CAMPAIGN
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateCampaign) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateCampaign) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
