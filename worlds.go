package gw2api

// World is a game world
type World struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population string `json:"population"`
}

// GetWorld returns worlds
func (s *Session) GetWorld(ids ...int) (worlds []*World, err error) {
	err = s.get(concatStrings("/v2/worlds", genArgs(ids...)), &worlds)
	return
}
