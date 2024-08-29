package gw2api

// JadeBot is an unlockable Jade Bot Skin.
type JadeBot struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UnlockItem  int    `json:"unlock_item"`
}

// JadeBots returns the  list of Jade Bots given by the `ids`.
func (s *Session) JadeBots(ids ...int) (bots []*JadeBot, err error) {
	err = s.get(concatStrings("/v2/jadebots", genArgs(ids...)), &bots)
	return
}
