package gw2api

// Legend is a revenant legend
type Legend struct {
	ID        string `json:"id"`
	Swap      int    `json:"swap"`
	Heal      int    `json:"heal"`
	Elite     int    `json:"elite"`
	Utilities []int  `json:"utilities"`
}

// Legends returns information about revenant legends
func (s *Session) Legends(ids ...string) (rsp []*Legend, err error) {
	err = s.get(concatStrings("/v2/legends", genArgsString(ids...)), &rsp)
	return
}
