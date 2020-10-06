package gw2api

// Finisher is a game finisher
type Finisher struct {
	ID            int    `json:"id"`
	UnlockDetails string `json:"unlock_details"`
	UnlockItems   []int  `json:"unlock_items"`
	Order         int    `json:"order"`
	Icon          string `json:"icon"`
	Name          string `json:"name"`
}

// Finishers returns the requested finishers
func (s *Session) Finishers(ids ...int) (rsp []*Finisher, err error) {
	err = s.get(concatStrings("/v2/finishers", genArgs(ids...)), &rsp)
	return
}
