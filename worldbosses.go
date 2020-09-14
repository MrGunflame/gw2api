package gw2api

// WorldBoss is a worldboss chest that can be opened once per day
type WorldBoss struct {
	ID string `json:"id"`
}

// WorldBosses returns worldboss chest that can be opened once per day
func (s *Session) WorldBosses(ids ...string) (rsp *[]WorldBoss, err error) {
	err = s.get(concatStrings("/v2/worldbosses", genArgsString(ids...)), &rsp)
	return
}
