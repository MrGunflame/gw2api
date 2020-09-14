package gw2api

// MapChest is a chest lootable once per day
type MapChest struct {
	ID string `json:"id"`
}

// MapChests returns all dayil lootable mapchests
func (s *Session) MapChests(ids ...string) (rsp []*MapChest, err error) {
	err = s.get(concatStrings("/v2/mapchests", genArgsString(ids...)), &rsp)
	return
}
