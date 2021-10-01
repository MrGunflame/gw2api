package gw2api

// LegendaryArmoryItem is an item that can be added into the Legendary Armory
type LegendaryArmoryItem struct {
	ID       int `json:"id"`
	MaxCount int `json:"max_count"`
}

// LegendaryArmory returns all items that can be added to the legendary armory
func (s *Session) LegendaryArmory(ids ...int) (res []*LegendaryArmoryItem, err error) {
	err = s.get(concatStrings("/v2/legendaryarmory", genArgs(ids...)), &res)
	return
}
