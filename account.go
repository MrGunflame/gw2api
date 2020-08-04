package gw2api

import (
  "time"
)

type AccountBankItem struct {
  ID int `json:"id"`
  Count int `json:"count"`
  Skin int `json:"skin"`
  Upgrades []int `json:"upgrades"`
  UpgradeSlotIndices []int `json:"upgrade_slot_indices"`
  Dyes []int `json:"dyes"`
  Binding string `json:"binding"`
  Stats struct {
    ID int `json:"id"`
    Attributes struct {
      // TODO: ADD STATS, POWER, PRECISION, ETC.
    } `json:"attributes"`
  } `json:"stats"`
}

type AccountDailyCraftingItem struct {
  ID string `json:"id"`
}

type AccountFinisher struct {
  ID int `json:"id"`
  Permanent bool `json:"permanent"`
  Quantity int `json:"qantity"`
}

type AccountHomeCat struct {
  ID int `json:"id"`
  Hint string `json:"hint"`
}

type Account struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Age time.Duration `json:"age"`
  World int `json:"world"`
  Guilds []string `json:"guilds"`
  GuildLeader []string `json:"guild_leader"`
  Created time.Time `json:"created"`
  Access []string `json:"access"`
  Commander bool `json:"commander"`
  FractalLevel int `json:"fractal_level"`
  DailyAP int `json:"daily_ap"`
  MonthlyAP int `json:"monthly_ap"`
  WvWRank int `json:"wvw_rank"`
}

func (s *Session) GetAccount() (account Account, err error) {
  err = s.getWithAuth("/v2/account", &account)
  return
}

type AccountAchievement struct {
  ID int `json:"id"`
  Current int `json:"current"`
  Max int `json:"max"`
  Done bool `json:"done"`
}

func (s *Session) GetAccountAchievements(ids...int) (accountAchievements []*AccountAchievement, err error) {
  err = s.getWithAuth(concatStrings("/v2/account/achievements", genArgs(ids...)), &accountAchievements)
  return
}

func (s *Session) GetAccountBuildStorage(ids...int) (buildstorage BuildTemplates, err error) {
  err = s.getWithAuth(concatStrings("/v2/account/buildstorage", genArgs(ids...)), &buildstorage)
  return
}

func (s *Session) GetAccountDailyCrafting(ids...int) (items []*AccountDailyCraftingItem, err error) {
  err = s.getWithAuth(concatStrings("/v2/account/dailycrafting", genArgs(ids...)), &items)
  return
}

// NOT YET IMPLEMENTED
func (s *Session) GetAccountDungeons() {
  
}

func (s *Session) GetAccountDyes() (dyes []int, err error) {
  err = s.getWithAuth("/v2/account/dyes", &dyes)
  return
}

func (s *Session) GetAccountFinishers() (finishers []*AccountFinisher, err error) {
  err = s.getWithAuth("/v2/account/finishers", &finishers)
  return
}

func (s *Session) GetAccountGliders() (gliders []int, err error) {
  err = s.getWithAuth("/v2/account/gliders", &gliders)
  return
}

func (s *Session) GetAccountHomeCats() (cats []*AccountHomeCat, err error) {
  err = s.get("/v2/account/home/cats", &cats)
  return
}
