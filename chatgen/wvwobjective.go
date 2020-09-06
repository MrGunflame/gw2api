package chatgen

import (
	"encoding/base64"
)

// EncodeWvWObjective encodes a wvw objective into a chatlink
func EncodeWvWObjective(objectiveID, mapID int) string {
	bytes := []byte{headerByteWvWObjective}

	for i := 0; i < 4; i++ {
		bytes = append(bytes, byte(objectiveID>>(8*i)))
	}

	for i := 0; i < 4; i++ {
		bytes = append(bytes, byte(mapID>>(8*i)))
	}

	return addLinkEnclosing(base64.StdEncoding.EncodeToString(bytes))
}

// DecodeWvWObjective decodes a chatlink into the contained wvw objective
func DecodeWvWObjective(link string) (int, int, error) {
	bytes, err := base64.StdEncoding.DecodeString(removeLinkEnclosing(link))
	if err != nil {
		return 0, 0, err
	}

	if len(bytes) < 9 || bytes[0] != headerByteWvWObjective {
		return 0, 0, ErrInvalidHeader
	}

	var objectiveID int
	for i := 0; i < 4; i++ {
		objectiveID += int(bytes[i+1]) << (8 * i)
	}

	var mapID int
	for i := 0; i < 4; i++ {
		mapID += int(bytes[i+5]) << (8 * i)
	}

	return objectiveID, mapID, nil
}
