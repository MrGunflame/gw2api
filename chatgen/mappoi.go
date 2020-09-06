package chatgen

import (
	"encoding/base64"
)

// EncodeMapPoI encodes a map poi into a chatcode
func EncodeMapPoI(id int) string {

	bytes := []byte{headerByteMapPoI}

	for i := 0; i < 3; i++ {
		bytes = append(bytes, byte(id>>(8*i)))
	}

	bytes = append(bytes, 0)

	return addLinkEnclosing(base64.StdEncoding.EncodeToString(bytes))
}

// DecodeMapPoI decodes a map poi chatcode and returns the contained id
func DecodeMapPoI(link string) (int, error) {

	bytes, err := base64.StdEncoding.DecodeString(removeLinkEnclosing(link))
	if err != nil {
		return 0, nil
	}

	if len(bytes) < 5 || bytes[0] != headerByteMapPoI {
		return 0, ErrInvalidHeader
	}

	var id int
	for i := 0; i < 3; i++ {
		id += int(bytes[i+1]) << (8 * i)
	}

	return id, nil
}
