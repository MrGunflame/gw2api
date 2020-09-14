package gw2api

// DailyCraftingItem contains data for a daily craftable item
type DailyCraftingItem struct {
	ID string `json:"id"`
}

// DailyCrafting returns the daily craftable items
func (s *Session) DailyCrafting(ids ...string) (rsp []*DailyCraftingItem, err error) {
	err = s.get(concatStrings("/v2/dailycrafting", genArgsString(ids...)), &rsp)
	return nil, err
}
