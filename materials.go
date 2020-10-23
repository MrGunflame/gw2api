package gw2api

// Material is a material in the material storage
type Material struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Items []int  `json:"items"`
	Order int    `json:"order"`
}

// Materials returns the materials with the given id
func (s *Session) Materials(ids ...int) (rsp []*Material, err error) {
	err = s.get(concatStrings("/v2/materials", genArgs(ids...)), &rsp)
	return
}
