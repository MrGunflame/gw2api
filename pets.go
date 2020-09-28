package gw2api

// Pet is a game ranger pet
type Pet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Skills      []struct {
		ID int `json:"id"`
	} `json:"skills"`
}

// Pets returns the requested game ranger pets
func (s *Session) Pets(ids ...int) (rsp []*Pet, err error) {
	err = s.get(concatStrings("/v2/pets", genArgs(ids...)), &rsp)
	return
}
