package gw2api

// Itemstat contains the stattype and attributes of an item
type Itemstat struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Attributes []struct {
		Attribute  string  `json:"attribute"`
		Multiplier float64 `json:"multiplier"`
		Value      int     `json:"value"`
	} `json:"attributes"`
}

// GetItemStats returns a list of itemstats
func (s *Session) GetItemStats(ids ...int) (itemstats []*Itemstat, err error) {
	err = s.get(concatStrings("/v2/itemstats", genArgs(ids...)), &itemstats)
	return
}
