package gw2api

// Title is a game title
type Title struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Achievement  int    `json:"achievement"`
	Achievements []int  `json:"achievements"`
	APRequired   int    `json:"ap_required"`
}

// Titles returns the titles with the given ids
func (s *Session) Titles(ids ...int) (resp []*Title, err error) {
	err = s.getWithLang(concatStrings("/v2/titles", genArgs(ids...)), &resp)
	return
}
