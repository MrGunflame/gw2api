package chatgen

import (
	"encoding/base64"
)

// EncodeSkill encodes a skill into the chatlink
func EncodeSkill(id int) string {

	bytes := []byte{headerByteSkill}

	for i := 0; i < 3; i++ {
		bytes = append(bytes, byte(id>>(8*i)))
	}

	bytes = append(bytes, 0)

	return addLinkEnclosing(base64.StdEncoding.EncodeToString(bytes))
}

// DecodeSkill decodes a skill chatlink into the contained id
func DecodeSkill(link string) (int, error) {
	bytes, err := base64.StdEncoding.DecodeString(removeLinkEnclosing(link))
	if err != nil {
		return 0, err
	}

	if len(bytes) < 5 || bytes[0] != headerByteSkill {
		return 0, ErrInvalidHeader
	}

	var id int
	for i := 0; i < 3; i++ {
		id += int(bytes[i+1]) << (8 * i)
	}

	return id, nil
}
