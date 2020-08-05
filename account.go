package gw2api

import (
	"time"
)

// AccountBankItem is an item from the account bank
type AccountBankItem struct {
	ID                 int    `json:"id"`
	Count              int    `json:"count"`
	Skin               int    `json:"skin"`
	Upgrades           []int  `json:"upgrades"`
	UpgradeSlotIndices []int  `json:"upgrade_slot_indices"`
	Dyes               []int  `json:"dyes"`
	Binding            string `json:"binding"`
	Stats              struct {
		ID         int `json:"id"`
		Attributes struct {
			// TODO: ADD STATS, POWER, PRECISION, ETC.
		} `json:"attributes"`
	} `json:"stats"`
}

// AccountDailyCraftingItem is an object crafted daily
type AccountDailyCraftingItem struct {
	ID string `json:"id"`
}

// AccountFinisher is an finisher unluck for an account
type AccountFinisher struct {
	ID        int  `json:"id"`
	Permanent bool `json:"permanent"`
	Quantity  int  `json:"qantity"`
}

// AccountHomeCat is a home cat node
type AccountHomeCat struct {
	ID   int    `json:"id"`
	Hint string `json:"hint"`
}

// Account includes general information
type Account struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Age          time.Duration `json:"age"`
	World        int           `json:"world"`
	Guilds       []string      `json:"guilds"`
	GuildLeader  []string      `json:"guild_leader"`
	Created      time.Time     `json:"created"`
	Access       []string      `json:"access"`
	Commander    bool          `json:"commander"`
	FractalLevel int           `json:"fractal_level"`
	DailyAP      int           `json:"daily_ap"`
	MonthlyAP    int           `json:"monthly_ap"`
	WvWRank      int           `json:"wvw_rank"`
}

// AccountAchievement contains the unlock status for an achievement
type AccountAchievement struct {
	ID      int  `json:"id"`
	Current int  `json:"current"`
	Max     int  `json:"max"`
	Done    bool `json:"done"`
}

// GetAccount returns the owner account of the apikey provided
func (s *Session) GetAccount() (account Account, err error) {
	err = s.getWithAuth("/v2/account", &account)
	return
}

// GetAccountAchievements returns the accounts achievement status for the given ids
func (s *Session) GetAccountAchievements(ids ...int) (accountAchievements []*AccountAchievement, err error) {
	err = s.getWithAuth(concatStrings("/v2/account/achievements", genArgs(ids...)), &accountAchievements)
	return
}

// GetAccountBuildStorage returns all builds stored in the accounts build storage
func (s *Session) GetAccountBuildStorage(ids ...int) (buildstorage []*BuildTemplate, err error) {
	err = s.getWithAuth(concatStrings("/v2/account/buildstorage", genArgs(ids...)), &buildstorage)
	return
}

// GetAccountDailyCrafting returns the accounts daily crafted items
func (s *Session) GetAccountDailyCrafting(ids ...int) (items []*AccountDailyCraftingItem, err error) {
	err = s.getWithAuth(concatStrings("/v2/account/dailycrafting", genArgs(ids...)), &items)
	return
}

// NOT YET IMPLEMENTED
func (s *Session) GetAccountDungeons() {

}

// GetAccountDyes returns all dye unlocks
func (s *Session) GetAccountDyes() (dyes []int, err error) {
	err = s.getWithAuth("/v2/account/dyes", &dyes)
	return
}

// GetAccountFinishers returns all finisher unlocks
func (s *Session) GetAccountFinishers() (finishers []*AccountFinisher, err error) {
	err = s.getWithAuth("/v2/account/finishers", &finishers)
	return
}

// GetAccountGliders returns all glider unlocks
func (s *Session) GetAccountGliders() (gliders []int, err error) {
	err = s.getWithAuth("/v2/account/gliders", &gliders)
	return
}

// GetAccountHomeCats returns all home cat unlocks
func (s *Session) GetAccountHomeCats() (cats []*AccountHomeCat, err error) {
	err = s.getWithAuth("/v2/account/home/cats", &cats)
	return
}

// GetAccountHomeNodes returns all home node unlocks
func (s *Session) GetAccountHomeNodes() (nodes []string, err error) {
	err = s.getWithAuth("/v2/account/home/nodes", &nodes)
	return
}

// GetAccountSharedInventory returns the items stored in the shared inventory slots
func (s *Session) GetAccountSharedInventory() (items []*ItemStack, err error) {
	err = s.getWithAuth("/v2/account/inventory", &items)
	return
}

// AccountLuck is the accounts consumed luck
type AccountLuck struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

// GetAccountLuck returns the accounts luck
func (s *Session) GetAccountLuck() (luck AccountLuck, err error) {
	err = s.getWithAuth("/v2/account/luck", &luck)
	return
}

// GetMailCarries returns all mail carrier unlocks
func (s *Session) GetMailCarriers() (carries []int, err error) {
	err = s.getWithAuth("/v2/account/mailcarries", &carries)
	return
}

// GetAccountMapChests returns all Hero's choice chests unlocked since daily reset
func (s *Session) GetAccountMapChests() (chests []string, err error) {
	err = s.getWithAuth("/v2/account/mapchests", &chests)
	return
}

// AccountMastery contains the unlocked mastery level
type AccountMastery struct {
	ID    string `json:"id"`
	Level int    `json:"level"`
}

// GetAccountMasteries returns all masteries and their levels unlocked
func (s *Session) GetAccountMasteries() (masteries []*AccountMastery, err error) {
	err = s.getWithAuth("/v2/account/masteries", &masteries)
	return
}

// AccountMasteryPoints contains the accounts mastery points for each region
type AccountMasteryPoints struct {
	Totals []struct {
		Region string `json:"region"` // The mastery region
		Spent  int    `json:"spent"`  // Amount of masteries of this region spent in mastery tracks
		Earned int    `json:"earned"` // Amount of masteries of this region earned for the account
	} `json:"totals"`
	Unlocked []int `json:"unlocked"` // Array of mastery ids
}

// GetAccountMasteryPoints returns the mastery points for each region
func (s *Session) GetAccountMasteryPoints() (masteryPoints AccountMasteryPoints, err error) {
	err = s.getWithAuth("/v2/account/mastery/points", &masteryPoints)
	return
}

// A Material is an item stored in the players vault/material storage
type Material struct {
	ID       int    `json:"id"`
	Category int    `json:"category"`
	Binding  string `json:"binding"`
	Count    int    `json:"count"`
}

// GetAccountMaterials returns the accounts materials stored
func (s *Session) GetAccountMaterials() (materials []*Material, err error) {
	err = s.getWithAuth("/v2/account/materials", &materials)
	return
}
