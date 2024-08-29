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

// Account contains general information
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

// Account returns general account information
func (s *Session) Account() (account Account, err error) {
	err = s.getWithAuth("/v2/account", &account)
	return
}

// AccountAchievement contains the unlock status for an achievement
type AccountAchievement struct {
	ID      int  `json:"id"`
	Current int  `json:"current"`
	Max     int  `json:"max"`
	Done    bool `json:"done"`
}

// AccountAchievements returns the accounts achievement status for the given ids
func (s *Session) AccountAchievements(ids ...int) (accountAchievements []*AccountAchievement, err error) {
	err = s.getWithAuth(concatStrings("/v2/account/achievements", genArgs(ids...)), &accountAchievements)
	return
}

// AccountBuildStorage returns all builds stored in the accounts build storage
func (s *Session) AccountBuildStorage(ids ...int) (buildstorage []*BuildTemplate, err error) {
	err = s.getWithAuth(concatStrings("/v2/account/buildstorage", genArgs(ids...)), &buildstorage)
	return
}

// AccountDailyCraftingItem is an object crafted daily
type AccountDailyCraftingItem struct {
	ID string `json:"id"`
}

// AccountDailyCrafting returns the accounts daily crafted items
func (s *Session) AccountDailyCrafting(ids ...int) (items []*AccountDailyCraftingItem, err error) {
	err = s.getWithAuth(concatStrings("/v2/account/dailycrafting", genArgs(ids...)), &items)
	return
}

// AccountDungeons returns the dungeons paths completed since daily reset
func (s *Session) AccountDungeons() (paths []string, err error) {
	err = s.getWithAuth("/v2/account/dungeons", &paths)
	return
}

// AccountDyes returns all dye unlocks
func (s *Session) AccountDyes() (dyes []int, err error) {
	err = s.getWithAuth("/v2/account/dyes", &dyes)
	return
}

// AccountFinisher is an finisher unluck for an account
type AccountFinisher struct {
	ID        int  `json:"id"`
	Permanent bool `json:"permanent"`
	Quantity  int  `json:"qantity"`
}

// AccountFinishers returns all finisher unlocks
func (s *Session) AccountFinishers() (finishers []*AccountFinisher, err error) {
	err = s.getWithAuth("/v2/account/finishers", &finishers)
	return
}

// AccountGliders returns all glider unlocks
func (s *Session) AccountGliders() (gliders []int, err error) {
	err = s.getWithAuth("/v2/account/gliders", &gliders)
	return
}

// AccountHomeCat is a home cat node
type AccountHomeCat struct {
	ID   int    `json:"id"`
	Hint string `json:"hint"`
}

// AccountHomeCats returns all home cat unlocks
func (s *Session) AccountHomeCats() (cats []*AccountHomeCat, err error) {
	err = s.getWithAuth("/v2/account/home/cats", &cats)
	return
}

// AccountHomeNodes returns all home node unlocks
func (s *Session) AccountHomeNodes() (nodes []string, err error) {
	err = s.getWithAuth("/v2/account/home/nodes", &nodes)
	return
}

// AccountSharedInventory returns the items stored in the shared inventory slots
func (s *Session) AccountSharedInventory() (items []*ItemStack, err error) {
	err = s.getWithAuth("/v2/account/inventory", &items)
	return
}

// AccountLuck is the accounts consumed luck
type AccountLuck struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

// AccountJadebots returns the list of Jade Bots unlocked by the account.
func (s *Session) AccountJadeBots() (bots []int, err error) {
	err = s.getWithAuth("/v2/account/jadebots", &bots)
	return
}

// AccountLuck returns the accounts luck
func (s *Session) AccountLuck() (luck AccountLuck, err error) {
	err = s.getWithAuth("/v2/account/luck", &luck)
	return
}

// AccountMailCarriers the accounts unlocked mailcarriers
func (s *Session) AccountMailCarriers() (carriers []int, err error) {
	err = s.getWithAuth("/v2/account/mailcarries", &carriers)
	return
}

// AccountMapChests returns all Hero's choice chests unlocked since daily reset
func (s *Session) AccountMapChests() (chests []string, err error) {
	err = s.getWithAuth("/v2/account/mapchests", &chests)
	return
}

// AccountMastery contains the unlocked mastery level
type AccountMastery struct {
	ID    string `json:"id"`
	Level int    `json:"level"`
}

// AccountMasteries returns all masteries and their levels unlocked
func (s *Session) AccountMasteries() (masteries []*AccountMastery, err error) {
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

// AccountMasteryPoints returns the mastery points for each region
func (s *Session) AccountMasteryPoints() (masteryPoints AccountMasteryPoints, err error) {
	err = s.getWithAuth("/v2/account/mastery/points", &masteryPoints)
	return
}

// A AccountMaterial is an item stored in the players vault/material storage
type AccountMaterial struct {
	ID       int    `json:"id"`
	Category int    `json:"category"`
	Binding  string `json:"binding"`
	Count    int    `json:"count"`
}

// AccountMaterials returns the accounts materials stored
func (s *Session) AccountMaterials() (materials []*Material, err error) {
	err = s.getWithAuth("/v2/account/materials", &materials)
	return
}

// AccountMinis returns the accounts mini unlocks
func (s *Session) AccountMinis() (minis []int, err error) {
	err = s.getWithAuth("/v2/account/minis", &minis)
	return
}

// AccountMountSkins returns the accounts mountskin unlocks
func (s *Session) AccountMountSkins() (skins []int, err error) {
	err = s.getWithAuth("/v2/account/mounts/skins", &skins)
	return
}

// AccountMountTypes returns the accounts mount unlocks
func (s *Session) AccountMountTypes() (types []int, err error) {
	err = s.getWithAuth("/v2/account/mounts/types", &types)
	return
}

// AccountNovelties returns the accounts novelties unlocks
func (s *Session) AccountNovelties() (novelties []int, err error) {
	err = s.getWithAuth("/v2/account/novelties", &novelties)
	return
}

// AccountOutfits returns the accounts outfit unlocks
func (s *Session) AccountOutfits() (outfits []int, err error) {
	err = s.getWithAuth("/v2/account/outfits", &outfits)
	return
}

// AccountPvPHeroes returns the accounts pvp hero unlocks
func (s *Session) AccountPvPHeroes() (heroes []int, err error) {
	err = s.getWithAuth("/v2/account/pvp/heroes", &heroes)
	return
}

// AccountRaids returns the completed raid encounters since the weekly raid reset
func (s *Session) AccountRaids() (bosses []string, err error) {
	err = s.getWithAuth("/v2/account/raids", &bosses)
	return
}

// AccountRecipes returns information about recipes that are unlocked for an account
func (s *Session) AccountRecipes() (recipes []int, err error) {
	err = s.getWithAuth("/v2/account/recipes", &recipes)
	return
}

// AccountSkins returns the unlocked skins of the account
func (s *Session) AccountSkins() (skins []int, err error) {
	err = s.getWithAuth("/v2/account/skins", &skins)
	return
}

// AccountTitles returns information about titles that are unlocked for an account
func (s *Session) AccountTitles() (titles []int, err error) {
	err = s.getWithAuth("/v2/account/titles", &titles)
	return
}

// AccountCurrency represents objects in the account wallet
type AccountCurrency struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

// AccountWallet returns the currencies of the account
func (s *Session) AccountWallet() (wallet []*AccountCurrency, err error) {
	err = s.getWithAuth("/v2/account/wallet", &wallet)
	return
}

// AccountWorldbosses returns information about which world bosses have been killed by the account since daily reset
func (s *Session) AccountWorldbosses() (bosses []string, err error) {
	err = s.getWithAuth("/v2/account/worldbosses", &bosses)
	return
}

// AccountLegendaryArmoryItem is an item added in the account's legendary armory
type AccountLegendaryArmoryItem struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}

// AccountLegendaryArmory returns the items in the account's legendary armory
func (s *Session) AccountLegendaryArmory() (res []*AccountLegendaryArmoryItem, err error) {
	err = s.getWithAuth("/v2/account/legendaryarmory", &res)
	return
}
