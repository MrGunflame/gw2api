package gw2api

// Achievement is a game achievement
type Achievement struct {
	ID          int      `json:"id"`
	Icon        string   `json:"icon"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	LockedText  string   `json:"locked_text"`
	Type        string   `json:"type"`
	Flags       []string `json:"flags"`
	Tiers       []struct {
		Count  int `json:"count"`
		Points int `json:"points"`
	} `json:"tiers"`
	Prerequisites []int `json:"prerequisites"`
	Rewards       []struct {
		Type string `json:"type"`
	} `json:"rewards"`
	Bits []struct {
		Type string `json:"type"`
		ID   int    `json:"id"`
		Text string `json:"text"`
	} `json:"bits"`
	PointCap int `json:"point_cap"`
}

// Achievements returns an game achievement
func (s *Session) Achievements(ids ...int) (achievements []*Achievement, err error) {
	err = s.get(concatStrings("/v2/achievements", genArgs(ids...)), &achievements)
	return
}

// DailyAchievement is a daiy achievement
type DailyAchievement struct {
	ID    int `json:"id"`
	Level struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"level"`
	RequiredAccess []string `json:"required_access"`
}

// DailyAchievements returns a map of daily achievements. Each map key stands for a different section, e.g. pve, wvw, fractals, ...
func (s *Session) DailyAchievements() (daily map[string][]*DailyAchievement, err error) {
	err = s.get("/v2/achievements/daily", &daily)
	return
}

// TomorrowDailyAchievements works the same as DailyAchievements. It returns the daily achievements for tomorrow
func (s *Session) TomorrowDailyAchievements() (daily map[string][]*DailyAchievement, err error) {
	err = s.get("/v2/achievements/daily/tomorrow", &daily)
	return
}

// AchievementGroup is an achievement group
type AchievementGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	Categories  []int  `json:"categories"`
}

// AchievementGroups returns all achievement groups
func (s *Session) AchievementGroups(ids ...string) (groups []*AchievementGroup, err error) {
	err = s.get(concatStrings("/v2/achievements/groups", genArgsString(ids...)), &groups)
	return
}

// AchievementCategory is an achievement category
type AchievementCategory struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Order        int    `json:"order"`
	Icon         string `json:"icon"`
	Achievements []int  `json:"achievement"`
}

// AchievementCategories returns all achievement categories
func (s *Session) AchievementCategories(ids ...int) (categories []*AchievementCategory, err error) {
	err = s.get(concatStrings("/v2/achievements/categories", genArgs(ids...)), &categories)
	return
}
