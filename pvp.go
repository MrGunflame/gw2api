package gw2api

// PvPStats contains the accounts pvp stats
type PvPStats struct {
	PvPRank          int                        `json:"pvp_rank"`
	PvPRankPoints    int                        `json:"pvp_rank_points"`
	PvPRankRollovers int                        `json:"pvp_rank_rollovers"`
	Aggregate        PvPWinLossStats            `json:"aggregate"`
	Professions      map[string]PvPWinLossStats `json:"professions"`
	Ladders          map[string]PvPWinLossStats `json:"ladders"`
}

// PvPWinLossStats contains the wins and losses
type PvPWinLossStats struct {
	Wins       int `json:"wins"`
	Losses     int `json:"losses"`
	Desertions int `json:"desertions"`
	Byes       int `json:"byes"`
	Forfeits   int `json:"forfeits"`
}

// PvPGame contains information about played pvp games
type PvPGame struct {
	ID      string `json:"id"`
	MapID   int    `json:"map_id"`
	Started string `json:"started"`
	Ended   string `json:"ended"`
	Result  string `json:"result"`
	Team    string `json:"team"`
	Scores  struct {
		Red  int `json:"red"`
		Blue int `json:"blue"`
	} `json:"scores"`
	RatingType   string `json:"rating_type"`
	RatingChange int    `json:"rating_change"`
	Season       string `json:"season"`
}
