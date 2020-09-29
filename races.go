package gw2api

// Race is a game race
type Race struct {
	ID     string `json:"id"`
	Skills []int  `json:"skills"`
}

// Races returns the requested races
func (s *Session) Races(ids ...string) (rsp []*Race, err error) {
	s.get(concatStrings("/v2/races", genArgsString(ids...)), &rsp)
	return
}
