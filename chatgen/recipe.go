package chatgen

import "encoding/base64"

// EncodeRecipe encodes a recipe into the chatcode
func EncodeRecipe(id int) string {
	bytes := []byte{headerByteRecipe}

	for i := 0; i < 3; i++ {
		bytes = append(bytes, byte(id>>(8*i)))
	}

	bytes = append(bytes, 0)

	return addLinkEnclosing(base64.StdEncoding.EncodeToString(bytes))
}

// DecodeRecipe decodes a chatlink into the contained recipe
func DecodeRecipe(link string) (int, error) {
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
