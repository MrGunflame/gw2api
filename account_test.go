package gw2api

import (
	"os"
	"testing"
)

func TestAccount(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.Account(); err != nil {
		t.Errorf("Account failed fetching the endpoint: '%s'", err)
	}
}

func TestAccountAchievements(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountAchievements(); err != nil {
		t.Errorf("AccountAchievements failed fetching the endpoint: '%s'", err)
	}
}

func TestAccountBuildStorage(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountBuildStorage(); err != nil {
		t.Errorf("AccountBuildStorage failed fetching endpoint: '%s'", err)
	}
}

func TestAccountDailyCrafting(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountDailyCrafting(); err != nil {
		t.Errorf("AccountDailyCrafting failed fetching endpoint: '%s'", err)
	}
}

func TestAccountDungeons(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountDungeons(); err != nil {
		t.Errorf("AccountDungeons failed fetching endpoint: '%s'", err)
	}
}

func TestAccountDyes(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountDyes(); err != nil {
		t.Errorf("AccountDyes failed fetching endpoint: '%s'", err)
	}
}

func TestAccountFinishers(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountFinishers(); err != nil {
		t.Errorf("AccountFinishers failed fetching endpoint: '%s'", err)
	}
}

func TestAccountGliders(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountGliders(); err != nil {
		t.Errorf("AccountGliders failed fetching endpoint: '%s'", err)
	}
}

func TestAccountHomeCats(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountHomeCats(); err != nil {
		t.Errorf("AccountHomeCats failed fetching endpoint: '%s'", err)
	}
}

func TestAccountSharedInventory(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountSharedInventory(); err != nil {
		t.Errorf("AccountSharedInventory failed fetching endpoint: '%s'", err)
	}
}

func TestAccountJadeBots(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountJadeBots(); err != nil {
		t.Errorf("AccountJadeBots failed fetching endpoint: '%s'", err)
	}
}

func TestAccountLuck(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountLuck(); err != nil {
		t.Errorf("AccountLuck failed fetching endpoint: '%s'", err)
	}
}

func TestAccountMailCarriers(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountMailCarriers(); err != nil {
		t.Errorf("AccountMailCarriers failed fetching endpoint: '%s'", err)
	}
}

func TestAccountMapChests(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountMapChests(); err != nil {
		t.Errorf("AccountMapChests failed fetching endpoint: '%s'", err)
	}
}

func TestAccountMasteries(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountMasteries(); err != nil {
		t.Errorf("AccountMasteries failed fetching endpoint: '%s'", err)
	}
}

func TestAccountMasteryPoints(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountMasteryPoints(); err != nil {
		t.Errorf("AccountMasteryPoints failed fetching endpoint: '%s'", err)
	}
}

func TestAccountMaterials(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountMaterials(); err != nil {
		t.Errorf("AccountMasteryPoints failed fetching endpoint: '%s'", err)
	}
}

func TestAccountMinis(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountMinis(); err != nil {
		t.Errorf("AccountMinis failed fetching endpoint: '%s'", err)
	}
}

func TestAccountMountSkins(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountMountSkins(); err != nil {
		t.Errorf("AccountMountSkins failed fetching endpoint: '%s'", err)
	}
}

func TestAccountMountTypes(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountMountTypes(); err != nil {
		t.Errorf("AccountMountTypes failed fetching endpoint: '%s'", err)
	}
}

func TestAccountNovelties(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountNovelties(); err != nil {
		t.Errorf("AccountNovelties failed fetching endpoint: '%s'", err)
	}
}

func TestAccountOutfits(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountOutfits(); err != nil {
		t.Errorf("AccountOutfits failed fetching endpoint: '%s'", err)
	}
}

func TestAccountPvPHeroes(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountPvPHeroes(); err != nil {
		t.Errorf("AccountPvPHeroes failed fetching endpoint: '%s'", err)
	}
}

func TestAccountRaids(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountRaids(); err != nil {
		t.Errorf("AccountRaids failed fetching endpoint: '%s'", err)
	}
}

func TestAccountRecipes(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountRecipes(); err != nil {
		t.Errorf("AccountRecipes failed fetching endpoint: '%s'", err)
	}
}

func TestAccountSkins(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountSkins(); err != nil {
		t.Errorf("AccountSkins failed fetching endpoint: '%s'", err)
	}
}

func TestAccountTitles(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountTitles(); err != nil {
		t.Errorf("AccountTitles failed fetching endpoint: '%s'", err)
	}
}

func TestAccountWallet(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountWallet(); err != nil {
		t.Errorf("AccountWallet failed fetching endpoint: '%s'", err)
	}
}

func TestAccountWorldbosses(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountWorldbosses(); err != nil {
		t.Errorf("AccountWorldbosses failed fetching endpoint: '%s'", err)
	}
}

func TestAccountLegendaryArmory(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIkey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AccountLegendaryArmory(); err != nil {
		t.Errorf("AccountLegendaryArmory failed: '%s'", err)
	}
}
