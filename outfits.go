package gw2api

// Outfit is a game outfit
type Outfit struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	UnlockItems []int  `json:"unlock_items"`
}

// Outfits returns requested outfits
func (s *Session) Outfits(ids ...int) (rsp []*Outfit, err error) {
	err = s.get(concatStrings("/v2/outfits", genArgs(ids...)), &rsp)
	return
}
