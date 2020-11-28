package gw2api

// Raid is a game raid instance
type Raid struct {
	ID    string `json:"id"`
	Wings []struct {
		ID     string `json:"id"`
		Events []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"events"`
	} `json:"wings"`
}

// Raids returns the raids with the given ids
func (s *Session) Raids(ids ...string) (resp []*Raid, err error) {
	err = s.getWithLang(concatStrings("/v2/raids", genArgsString(ids...)), &resp)
	return
}
