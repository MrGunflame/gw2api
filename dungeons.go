package gw2api

// Dungeon is a game dungeon
type Dungeon struct {
	ID    string `json:"id"`
	Paths []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"paths"`
}

// Dungeons returns the dungeons with the given ids
func (s *Session) Dungeons(ids ...string) (resp []*Dungeon, err error) {
	err = s.getWithLang(concatStrings("/v2/dungeons", genArgsString(ids...)), &resp)
	return
}
