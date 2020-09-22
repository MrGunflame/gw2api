package gw2api

// Mastery contains information about a unlockable mastery
type Mastery struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Requirement string `json:"requirement"`
	Order       int    `json:"order"`
	Background  string `json:"background"`
	Region      string `json:"region"`
	Levels      []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Instruction string `json:"instruction"`
		Icon        string `json:"icon"`
		PointCost   int    `json:"point_cost"`
		ExpCost     int    `json:"exp_cost"`
	} `json:"levels"`
}

// Masteries returns all trainable masteries
func (s *Session) Masteries(ids ...int) (rsp []*Mastery, err error) {
	err = s.get(concatStrings("/v2/masteries", genArgs(ids...)), &rsp)
	return
}
