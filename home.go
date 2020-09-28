package gw2api

// HomeCat is a home instance cat
type HomeCat struct {
	ID   int    `json:"id"`
	Hint string `json:"hint"`
}

// HomeCats returns home instance cats
func (s *Session) HomeCats(ids ...int) (rsp []*HomeCat, err error) {
	err = s.get(concatStrings("/v2/home/cats", genArgs(ids...)), &rsp)
	return
}

// HomeNode is a home instance node
type HomeNode struct {
	ID string `json:"id"`
}

// HomeNodes returns home instance nodes
func (s *Session) HomeNodes(ids ...string) (rsp []*HomeNode, err error) {
	err = s.get(concatStrings("/v2/home/nodes", genArgsString(ids...)), &rsp)
	return
}
