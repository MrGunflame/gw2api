package gw2api

// Build contains the id of a game build
type Build struct {
	ID int `json:"id"`
}

// Build returns the latest game build
func (s *Session) Build() (r *Build, err error) {
	err = s.get("/v2/build", &r)
	return
}
