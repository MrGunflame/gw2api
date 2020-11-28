package gw2api

// Map is a game map
type Map struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	MinLevel      int     `json:"min_level"`
	MaxLevel      int     `json:"max_level"`
	DefaultFloot  int     `json:"default_floor"`
	Floors        []int   `json:"floors"`
	RegionID      int     `json:"region_id"`
	RegionName    string  `json:"region_name"`
	ContinentID   int     `json:"continent_id"`
	ContinentName string  `json:"continent_name"`
	MapRect       [][]int `json:"map_rect"`
	ContinentRect [][]int `json:"continent_rect"`
}

// Maps returns the maps with the given ids
func (s *Session) Maps(ids ...int) (resp []*Map, err error) {
	err = s.getWithLang(concatStrings("/v2/maps", genArgs(ids...)), &resp)
	return
}
