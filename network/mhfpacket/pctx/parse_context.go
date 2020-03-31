package pctx

import (
	"github.com/Andoryuuta/Erupe/common/stringsupport"
)

// PacketContext holds contextual data required for packet parsing.
type PacketContext struct {
	StrConv *stringsupport.StringConverter
}
