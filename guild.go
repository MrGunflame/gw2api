package gw2api

// Guild contains the core information about a guild
type Guild struct {
	Level          int    `json:"level"`
	MOTD           string `json:"motd"`
	Influence      int    `json:"influence"`
	Aetherium      int    `json:"aetherium"`
	Favor          int    `json:"favor"`
	MemberCount    int    `json:"member_count"`
	MemberCapacity int    `json:"member_capacity"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Tag            string `json:"tag"`
	Emblem         struct {
		Background struct {
			ID     int   `json:"id"`
			Colors []int `json:"colors"`
		} `json:"background"`
		Foreground struct {
			ID     int   `json:"id"`
			Colors []int `json:"colors"`
		} `json:"foreground"`
	} `json:"emblem"`
	Flags []string `json:"flags"`
}

// GuildPermission is a assignable guild permission
type GuildPermission struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GuildUpgrade contains information about upgrades for guild halls
type GuildUpgrade struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Type          string `json:"type"`
	BagMaxItems   int    `json:"bag_max_items"` // When Type is "BankBag"
	BagMaxCoins   int    `json:"bag_max_coins"` // When Type us "BankBag"
	Icon          string `json:"icon"`
	BuildTime     int    `json:"build_time"`
	RequiredLevel int    `json:"required_level"`
	Experience    int    `json:"experience"`
	Prerequisites []int  `json:"prerequisites"`
	Costs         []struct {
		Type   string `json:"type"`
		Name   string `json:"name"`
		Count  int    `json:"count"`
		ItemID int    `json:"item_id"`
	} `json:"costs"`
}

// GuildLogEntry is an entry in the guild log
type GuildLogEntry struct {
	ID        int    `json:"id"`
	Time      string `json:"time"`
	User      string `json:"user"`
	Type      string `json:"type"`
	InvitedBy string `json:"invited_by"` // When Type is "invited"
	KickedBy  string `json:"kicked_by"`  // When Type is "kick"
	ChangedBy string `json:"changed_by"` // When Type is "rank_change"
	OldRank   string `json:"old_rank"`   // When Type is "rank_change"
	NewRank   string `json:"new_rank"`   // When Type is "rank_change"
	ItemID    int    `json:"item_id"`    // When Type is "treasury" or "stash"
	Count     int    `json:"count"`      // When Type is "treasury" or "stash"
	Operation string `json:"operation"`  // When Type is "stash"
	Coins     int    `json:"coins"`      // When Type is "stash"
	MOTD      string `json:"motd"`       // When Type is "motd"
	Action    string `json:"action"`     // When Type is "upgrade"
	UpgradeID int    `json:"upgrade_id"` // When Type is "upgrade"
	RecipeID  int    `json:"recipe_id"`  // When Type is "upgrade", optional
}

// GuildMember is a guild member
type GuildMember struct {
	Name   string `json:"name"`
	Rank   string `json:"rank"`
	Joined string `json:"joined"`
}

// GuildRank is a guild rank
type GuildRank struct {
	ID          string   `json:"id"`
	Order       int      `json:"order"`
	Permissions []string `json:"permissions"`
	Icon        string   `json:"icon"`
}

// GuildStashContainer is a section of the guild stash holding items
type GuildStashContainer struct {
	UpgradeID int    `json:"upgrade_id"`
	Size      int    `json:"size"`
	Coins     int    `json:"coins"`
	Note      string `json:"note"`
	Iventory  []struct {
		ID    int `json:"id"`
		Count int `json:"count"`
	} `json:"inventory"`
}

// GuildTreasurySlot is a slot for items needed for guild ugrades
type GuildTreasurySlot struct {
	ItemID   int `json:"item_id"`
	Count    int `json:"count"`
	NeededBy []struct {
		UpgradeID int `json:"upgrade_id"`
		Count     int `json:"count"`
	} `json:"needed_by"`
}

// GuildTeam contains information about a guild pvp team
type GuildTeam struct {
	ID      int `json:"id"`
	Members []struct {
		Name string `json:"name"`
		Role string `json:"role"`
	} `json:"members"`
	Name      string          `json:"name"`
	Aggregate PvPWinLossStats `json:"aggregate"`
	Ladders   PvPWinLossStats `json:"ladders"`
	Games     []*PvPGame      `json:"games"` // Profession field is omitted
	Seasons   []struct {
		ID     string `json:"id"`
		Wins   int    `json:"wins"`
		Losses int    `json:"losses"`
		Rating int    `json:"rating"`
	} `json:"seasons"`
}

// Guild returns the guilds general information
func (s *Session) Guild(guild string, auth bool) (resp Guild, err error) {
	err = s.getWithAuth(concatStrings("/v2/guild/", guild), &resp)
	return
}

// GuildPermissions returns the assignable guild permissions
func (s *Session) GuildPermissions(ids ...string) (resp []*GuildPermission, err error) {
	err = s.get(concatStrings("/v2/guild/permissions", genArgsString(ids...)), &resp)
	return
}

// GuildSearch searches for guilds based on the given name. It returns an array of guild ids
func (s *Session) GuildSearch(guild string) (resp []string, err error) {
	err = s.get(concatStrings("/v2/guild/search?name=", guild), &resp)
	return
}

// GuildHallUpgrades returns unlockable upgrades for the guild hall. Note: This method is call GetGuildHallUpgrades as GetGuildUpgrades returns the guilds unlocked upgrades
func (s *Session) GuildHallUpgrades(ids ...int) (resp []*GuildUpgrade, err error) {
	err = s.get(concatStrings("/v2/guild/upgrades", genArgs(ids...)), &resp)
	return
}

// GuildLog returns the guild log events
func (s *Session) GuildLog(guild string) (resp []*GuildLogEntry, err error) {
	err = s.getWithAuth(concatStrings("/v2/guild/", guild, "/log"), &resp)
	return
}

// GuildMembers returns all guild members
func (s *Session) GuildMembers(guild string) (resp []*GuildMember, err error) {
	err = s.getWithAuth(concatStrings("/v2/guild/", guild, "/members"), &resp)
	return
}

// GuildRanks returns all guild ranks
func (s *Session) GuildRanks(guild string) (resp []*GuildRank, err error) {
	err = s.getWithAuth(concatStrings("/V2/guild/", guild, "/ranks"), &resp)
	return
}

// GuildStash returns all items in the guild stash
func (s *Session) GuildStash(guild string) (resp []*GuildStashContainer, err error) {
	err = s.getWithAuth(concatStrings("/v2/guild/", guild, "/stash"), &resp)
	return
}

// GuildTreasury returns all items in the guild treasury
func (s *Session) GuildTreasury(guild string) (resp []*GuildTreasurySlot, err error) {
	err = s.getWithAuth(concatStrings("/v2/guild/", guild, "/treasury"), &resp)
	return
}

// GuildTeams returns all guild teams
func (s *Session) GuildTeams(guild string) (resp []*GuildTeam, err error) {
	err = s.getWithAuth(concatStrings("/v2/guild/", guild, "/teams"), &resp)
	return
}

// GuildUpgrades returns all unlocked guild upgrades
func (s *Session) GuildUpgrades(guild string) (resp []int, err error) {
	err = s.getWithAuth(concatStrings("/v2/guild/", guild, "/upgrades"), &resp)
	return
}
