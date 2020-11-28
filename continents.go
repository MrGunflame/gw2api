package gw2api

// Continent is a game continent
type Continent struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ContinentDims []int  `json:"continent_dims"`
	MinZoom       int    `json:"min_zoom"`
	MaxZoom       int    `json:"max_zoom"`
	Floors        []int  `json:"floors"`
}

// Continents returns the continents with the given id
func (s *Session) Continents(ids ...int) (resp []*Continent, err error) {
	err = s.get(concatStrings("/v2/continents", genArgs(ids...)), &resp)
	return
}
