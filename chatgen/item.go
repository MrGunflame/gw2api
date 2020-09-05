package chatgen

import (
	"encoding/base64"
)

// Item contains information for encoding or decoding item links
type Item struct {
	Quantity uint8
	ID       int
	Upgrades [2]int
	Skin     int
}

// EncodeItem returns the encoded chatlink for that item
func EncodeItem(item Item) string {

	// Header and item quantity
	bytes := []byte{headerByteItem, byte(item.Quantity)}

	// Item id
	for i := 0; i < 3; i++ {
		bytes = append(bytes, byte(item.ID>>(8*i)))
	}

	// Flag byte
	bytes = append(bytes, 0)
	flagByteIndex := len(bytes) - 1

	// Item Skin
	if item.Skin != 0 {
		bytes[flagByteIndex] += byte(0x80)

		for i := 0; i < 3; i++ {
			bytes = append(bytes, byte(item.Skin>>(8*i)))
		}

		bytes = append(bytes, 0)
	}

	// Upgrades
	for i, u := range item.Upgrades {
		if u == 0 {
			break
		}

		bytes[flagByteIndex] += byte(64 >> i)

		for j := 0; j < 3; j++ {
			bytes = append(bytes, byte(u>>(8*j)))
		}

		bytes = append(bytes, 0)
	}

	return addLinkEnclosing(base64.StdEncoding.EncodeToString(bytes))
}

// DecodeItem reads the item data from a link
func DecodeItem(link string) (*Item, error) {

	bytes, err := base64.StdEncoding.DecodeString(removeLinkEnclosing(link))
	if err != nil {
		return nil, err
	}

	if len(bytes) < 5 || bytes[0] != headerByteItem {
		return nil, ErrInvalidHeader
	}

	item := &Item{}
	var flagByte byte

	for i, b := range bytes {
		switch i {
		case 1:
			item.Quantity = uint8(b)
		case 2, 3, 4:
			item.ID += int(b) << (8 * (i - 2))
		case 5:
			flagByte = b
		case 6, 7, 8:
			if (flagByte & 0x80) != 0 {
				item.Skin += int(b) << (8 * (i - 6))
			} else {
				item.Upgrades[0] += int(b) << (8 * (i - 6))
			}
		case 10, 11, 12:
			if (flagByte & 0x80) != 0 {
				item.Upgrades[0] += int(b) << (8 * (i - 10))
			} else {
				item.Upgrades[1] += int(b) << (8 * (i - 10))
			}
		case 14, 15, 16:
			item.Upgrades[1] += int(b) << (8 * (i - 14))
		}
	}

	return item, nil
}
