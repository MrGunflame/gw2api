package gw2api

import "strconv"

// WvWAbility is a trainable ability in wvw
type WvWAbility struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Ranks       []struct {
		Cost   int    `json:"cost"`
		Effect string `json:"effect"`
	}
}

// WvWMatch is a match in wvw. All map values will have the key strings red, green and blue
type WvWMatch struct {
	ID            string           `json:"id"`
	StartTime     string           `json:"start_time"`
	EndTime       string           `json:"end_time"`
	Scores        map[string]int   `json:"scores"`
	Worlds        map[string]int   `json:"worlds"`
	AllWorlds     map[string][]int `json:"all_worlds"`
	Deaths        map[string]int   `json:"deaths"`
	Kills         map[string]int   `json:"kills"`
	VictoryPoints map[string]int   `json:"victory_points"`
	Skirmishes    []struct {
		ID        int            `json:"id"`
		Scores    map[string]int `json:"scores"`
		MapScores []struct {
			Type   string         `json:"type"`
			Scores map[string]int `json:"scores"`
		} `json:"map_scores"`
	} `json:"skirmishes"`
	Maps []struct {
		ID      int            `json:"id"`
		Type    string         `json:"type"`
		Scores  map[string]int `json:"scores"`
		Bonuses []struct {
			Type  string `json:"type"`
			Owner string `json:"owner"`
		} `json:"bonuses"`
		Objectives []struct {
			ID            string `json:"id"`
			Type          string `json:"type"`
			Owner         string `json:"owner"`
			LastFlipped   string `json:"last_flipped"`
			ClaimedBy     string `json:"claimed_by"`
			ClaimedAt     string `json:"claimed_at"`
			PointsTick    int    `json:"points_tick"`
			PointsCapture int    `json:"points_capture"`
			YaksDelivered int    `json:"yaks_delivered"`
			GuildUpgrades []int  `json:"guild_upgrades"`
		} `json:"objectives"`
		Deaths map[string]int `json:"deaths"`
		Kills  map[string]int `json:"kills"`
	} `json:"maps"`
}

// WvWObjective is a capturable structure in wvw
type WvWObjective struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	SectorID   int       `json:"sector_id"`
	Type       string    `json:"type"`
	MapType    string    `json:"map_type"`
	MapID      int       `json:"map_id"`
	UpgradeID  int       `json:"upgrade_id"`
	Coord      []float64 `json:"coord"`
	LabelCoord []float64 `json:"label_coord"`
	Marker     string    `json:"marker"`
	ChatLink   string    `json:"chat_link"`
}

// WvWRank is a rank as seen by enemy players
type WvWRank struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	MinRank int    `json:"min_rank"`
}

// WvWUpgrade is an automatic objective upgrade
type WvWUpgrade struct {
	ID    int `json:"id"`
	Tiers []struct {
		Name         string `json:"name"`
		YaksRequired int    `json:"yaks_required"`
		Upgrades     []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"upgrades"`
	} `json:"tiers"`
}

// WvWAbilities returns trainable wvw abilities
func (s *Session) WvWAbilities(ids ...int) (abilities []*WvWAbility, err error) {
	err = s.get(concatStrings("/v2/wvw/abilities", genArgs(ids...)), &abilities)
	return
}

// WvWMatches returns currently active matches
func (s *Session) WvWMatches(ids ...string) (matches []*WvWMatch, err error) {
	err = s.get(concatStrings("/v2/wvw/matches", genArgsString(ids...)), &matches)
	return
}

// WvWMatchByWorldID returns the match based on the participating world
func (s *Session) WvWMatchByWorldID(id int) (match *WvWMatch, err error) {
	err = s.get(concatStrings("/v2/wvw/matches?world=", strconv.Itoa(id)), &match)
	return
}

// WvWObjectives returns wvw objectives
func (s *Session) WvWObjectives(ids ...string) (objectives []*WvWObjective, err error) {
	err = s.get(concatStrings("/v2/wvw/objectives", genArgsString(ids...)), &objectives)
	return
}

// WvWRanks returns wvw ranks
func (s *Session) WvWRanks(ids ...int) (ranks []*WvWRank, err error) {
	err = s.get(concatStrings("/v2/wvw/ranks", genArgs(ids...)), &ranks)
	return
}

// WvWUpgrades returns wvw objectives upgrades
func (s *Session) WvWUpgrades(ids ...int) (upgrades []*WvWUpgrade, err error) {
	err = s.get(concatStrings("/v2/wvw/upgrades", genArgs(ids...)), &upgrades)
	return
}
