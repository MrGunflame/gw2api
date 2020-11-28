package gw2api

// Novelty is a game novelty
type Novelty struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Slot        string `json:"slot"`
	UnlockItem  []int  `json:"unlock_item"`
}

// Novelties returns the novelties with the given ids
func (s *Session) Novelties(ids ...int) (resp []*Novelty, err error) {
	err = s.getWithLang(concatStrings("/v2/novelties", genArgs(ids...)), &resp)
	return
}
