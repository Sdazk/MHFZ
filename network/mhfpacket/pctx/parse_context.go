package pctx

import (
	"golang.org/x/text/encoding"
)

// PacketContext holds contextual data required for packet parsing.
type PacketContext struct {
	Encoding encoding.Encoding
}
