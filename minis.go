package gw2api

// Mini is a game miniature
type Mini struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Unlock string `json:"unlock"`
	Icon   string `json:"icon"`
	Order  int    `json:"order"`
	ItemID int    `json:"item_id"`
}

// Minis returns the minis with the given ids
func (s *Session) Minis(ids ...int) (resp []*Mini, err error) {
	err = s.getWithLang(concatStrings("/v2/minis", genArgs(ids...)), &resp)
	return
}
