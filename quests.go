package gw2api

// Quest contains story journal mission information
type Quest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
	Story int    `json:"story"`
	Goals []struct {
		Active   string `json:"active"`
		Complete string `json:"complete"`
	} `json:"goals"`
}

// Quests returns the quests with the given ids
func (s *Session) Quests(ids ...int) (resp []*Quest, err error) {
	err = s.getWithLang(concatStrings("/v2/quests", genArgs(ids...)), &resp)
	return
}
