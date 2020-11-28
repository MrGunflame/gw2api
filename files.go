package gw2api

// File is a in-game asset ressource
type File struct {
	ID   string `json:"id"`
	Icon string `json:"icon"`
}

// Files returns the files with the given ids
func (s *Session) Files(ids ...string) (resp []*File, err error) {
	err = s.get(concatStrings("/v2/files", genArgsString(ids...)), &resp)
	return
}
