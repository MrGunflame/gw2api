package chatgen

import "encoding/base64"

// EncodeSkin encodes a skin into a chatlink
func EncodeSkin(id int) string {
	bytes := []byte{headerByteWardrope}

	for i := 0; i < 3; i++ {
		bytes = append(bytes, byte(id>>(8*i)))
	}

	bytes = append(bytes, 0)

	return addLinkEnclosing(base64.StdEncoding.EncodeToString(bytes))
}

// DecodeSkin decodes a chatlink into the contained skin
func DecodeSkin(link string) (int, error) {
	bytes, err := base64.StdEncoding.DecodeString(removeLinkEnclosing(link))
	if err != nil {
		return 0, err
	}

	var id int
	for i := 0; i < 3; i++ {
		id += int(bytes[i+1]) << (8 * i)
	}

	return id, nil
}
