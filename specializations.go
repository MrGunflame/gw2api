package gw2api

// Specialization is a game specialization
type Specialization struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Profession  string `json:"professions"`
	Elite       bool   `json:"elite"`
	Icon        string `json:"icon"`
	Background  string `json:"background"`
	MinorTraits []int  `json:"minor_traits"`
	MajorTraits []int  `json:"major_traits"`
}

// Specializations returns the requested specializations
func (s *Session) Specializations(ids ...int) (rsp []*Specialization, err error) {
	err = s.get(concatStrings("/v2/specializations", genArgs(ids...)), &rsp)
	return
}
