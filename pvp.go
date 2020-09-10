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
	Season       string `json:"season,omitempty"`
}

// PvPStandings contains information about the account season ranking and the best sesion rankings
type PvPStandings struct {
	Current []struct {
		TotalPoints int `json:"total_points"`
		Division    int `json:"division"`
		Tier        int `json:"tier"`
		Points      int `json:"points"`
		Repeats     int `json:"repeats"`
		Rating      int `json:"rating"`
		Decay       int `json:"decay"`
	} `json:"current"`
	Best []struct {
		TotalPoints int `json:"total_points"`
		Division    int `json:"division"`
		Tier        int `json:"tier"`
		Points      int `json:"points"`
		Repeats     int `json:"repeats"`
	} `json:"best"`
	SeasonID string `json:"season_id"`
}

// PvPRank is a pvp rank
type PvPRank struct {
	ID         int    `json:"id"`
	FinisherID int    `json:"finisher_id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	MinRank    int    `json:"min_rank"`
	MaxRank    int    `json:"max_rank"`
	Levels     []struct {
		MinRank int `json:"min_rank"`
		MaxRank int `json:"max_rank"`
		Points  int `json:"points"`
	} `json:"levels"`
}

// PvPSeason is a pvp season
type PvPSeason struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Active    bool   `json:"active"`
	Divisions []struct {
		Name      string   `json:"name"`
		Flags     []string `json:"flags"`
		LargeIcon string   `json:"large_icon"`
		SmallIcon string   `json:"small_icon"`
		PipIcon   string   `json:"pip_icon"`
		Tiers     []struct {
			Points int `json:"points"`
		} `json:"tiers"`
	} `json:"divisions"`
	Leaderboards struct {
		Ladder struct {
			Settings struct {
				Name     string `json:"name"`
				Duration int    `json:"duration"`
				Scoring  string `json:"scoring"`
				Tiers    []struct {
					Range []int `json:"range"`
				}
			} `json:"settings"`
			Scorings []struct {
				ID          string `json:"id"`
				Type        string `json:"type"`
				Description string `json:"description"`
				Name        string `json:"name"`
				Ordering    string `json:"ordering"`
			} `json:"scoring"`
		} `json:"ladder"`
	} `json:"leaderboards"`
}

// PvPLeaderboardAccount is an account displayed on the pvp leaderboards
type PvPLeaderboardAccount struct {
	Name   string `json:"name"`
	Rank   int    `json:"rank"`
	ID     string `json:"id"`
	Team   string `json:"team,omitempty"`    // Only for guild leaderboards
	TeamID int    `json:"team_id,omitempty"` // Only for guild leaderboards
	Date   string `json:"date"`
	Scores []struct {
		ID    string `json:"id"`
		Value int    `json:"value"`
	} `json:"scores"`
}

// PvPStats returns the accounts pvp stats
func (s *Session) PvPStats() (stats *PvPStats, err error) {
	err = s.getWithAuth("/v2/pvp/stats", &stats)
	return
}

// PvPGames returns the last 10 played matches
func (s *Session) PvPGames(ids ...string) (games *[]PvPGame, err error) {
	err = s.getWithAuth(concatStrings("/v2/pvp/games", genArgsString(ids...)), &games)
	return
}

// PvPStandings returns the current seasons account and best ranking
func (s *Session) PvPStandings() (standings *PvPStandings, err error) {
	err = s.getWithAuth("/v2/pvp/standings", &standings)
	return
}

// PvPAmulets return pvp amulets
func (s *Session) PvPAmulets(ids ...int) (amulets *[]Item, err error) {
	err = s.get(concatStrings("/v2/pvp/amulets", genArgs(ids...)), &amulets)
	return
}

// PvPRanks returns pvp ranks
func (s *Session) PvPRanks(ids ...int) (ranks []*PvPRank, err error) {
	err = s.get(concatStrings("/v2/pvp/ranks", genArgs(ids...)), &ranks)
	return
}

// PvPSeasons returns pvp seasons
func (s *Session) PvPSeasons(ids ...string) (seasons []*PvPSeason, err error) {
	err = s.get(concatStrings("/v2/pvp/seasons", genArgsString(ids...)), &seasons)
	return
}

// PvPSeasonLeaderboards returns the leaderboards for a pvp season. Only one id can be given. Valid regions are na or eu
func (s *Session) PvPSeasonLeaderboards(id, region string) (leaderboards []*PvPLeaderboardAccount, err error) {
	err = s.get(concatStrings("/v2/pvp/seasons/", id, "/leaderboards/ladder/", region), &leaderboards)
	return
}
