package gw2api

// Emblem is a guild emblem
type Emblem struct {
	ID     int      `json:"id"`
	Layers []string `json:"layers"`
}

// EmblemBackgrounds returns the requested background emblems
func (s *Session) EmblemBackgrounds(ids ...int) (rsp []*Emblem, err error) {
	err = s.get(concatStrings("/v2/emblem/backgrounds", genArgs(ids...)), &rsp)
	return
}

// EmblemForegrounds returns the requested foreground emblems
func (s *Session) EmblemForegrounds(ids ...int) (rsp []*Emblem, err error) {
	err = s.get(concatStrings("/v2/emblem/foregrounds", genArgs(ids...)), &rsp)
	return
}
