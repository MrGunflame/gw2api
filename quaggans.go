package gw2api

// Quaggan is a quaggaon image ressource
type Quaggan struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// Quaggans returns the quaggans with the given ids
func (s *Session) Quaggans(ids ...string) (resp []*Quaggan, err error) {
	err = s.get(concatStrings("/v2/quaggans", genArgsString(ids...)), &resp)
	return
}
