package gw2api

// Story contains information about story journal stories
type Story struct {
	ID          int    `json:"id"`
	Season      string `json:"season"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Timeline    string `json:"timeline"`
	Level       int    `json:"level"`
	Order       int    `json:"order"`
	Chapters    []struct {
		Name string `json:"name"`
	} `json:"chapters"`
	Races []string `json:"races"`
	Flags []string `json:"flags"`
}

// Stories returns the stories with the given ids
func (s *Session) Stories(ids ...int) (resp []*Story, err error) {
	err = s.getWithLang(concatStrings("/v2/stories", genArgs(ids...)), &resp)
	return
}
