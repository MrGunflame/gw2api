package gw2api

import (
	"image"
	"math/rand"

	_ "image/jpeg" // JPEG is the format the tile service returns images as
	"net/http"
	"strconv"
	"strings"
)

var tileService = []string{
	"tiles.guildwars2.com",
	"tiles1.guildwars2.com",
	"tiles2.guildwars2.com",
	"tiles3.guildwars2.com",
	"tiles4.guildwars2.com",
}

// Tile fetches a tile from the tile service
// The request is randomly send to one of the entries in the tileService slice
func (s *Session) Tile(continent, floor, zoom, x, y int) (image.Image, error) {

	// Convert all params to strings
	params := []int{continent, floor, zoom, x, y}
	paramsStr := make([]string, len(params))
	for i, p := range params {
		paramsStr[i] = strconv.Itoa(p)
	}

	resp, err := http.Get(concatStrings("https://", tileService[rand.Intn(4)], "/", strings.Join(paramsStr, "/"), ".jpg"))
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
