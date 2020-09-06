package chatgen

import "encoding/base64"

// EncodeOutfit encodes an outfit into a chatlink
func EncodeOutfit(id int) string {
	bytes := []byte{headerByteOutfit}

	for i := 0; i < 3; i++ {
		bytes = append(bytes, byte(id>>(8*i)))
	}

	bytes = append(bytes, 0)

	return addLinkEnclosing(base64.StdEncoding.EncodeToString(bytes))
}

// DecodeOutfit decodes a chatlink into the contained outfit
func DecodeOutfit(link string) (int, error) {
	bytes, err := base64.StdEncoding.DecodeString(removeLinkEnclosing(link))
	if err != nil {
		return 0, err
	}

	if len(bytes) < 5 || bytes[0] != headerByteOutfit {
		return 0, ErrInvalidHeader
	}

	var id int
	for i := 0; i < 3; i++ {
		id += int(bytes[i+1]) << (8 * i)
	}

	return id, nil
}
