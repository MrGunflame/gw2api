package chatgen

import (
	"fmt"
	"strings"
)

// header bytes for all items
const (
	headerByteCoin          = 0x01
	headerByteItem          = 0x02
	headerByteNPC           = 0x03
	headerByteMapObject     = 0x04
	headerBytePvPGame       = 0x05
	headerByteSkill         = 0x06
	headerByteTrait         = 0x07
	headerByteUser          = 0x08
	headerByteRecipe        = 0x09
	headerByteWardrope      = 0x0A
	headerByteOutfit        = 0x0B
	headerByteWvWObjective  = 0x0C
	headerByteBuildTemplate = 0x0D
)

const terminatorByte = 0

func removeLinkEnclosing(link string) string {
	return strings.Replace(strings.Replace(link, "]", "", 1), "[&", "", 1)
}

func addLinkEnclosing(base string) string {
	return fmt.Sprint("[&", base, "]")
}
