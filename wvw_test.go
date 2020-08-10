package gw2api

import "testing"

func TestWvW(t *testing.T) {
	api := New()

	if _, err := api.GetWvWAbilities(); err != nil {
		t.Error("GetWvWAbilities failed: ", err)
	}

	if _, err := api.GetWvWMatches(); err != nil {
		t.Error("GetWvWMatches failed: ", err)
	}

	if _, err := api.GetWvWMatchByWorldID(2003); err != nil {
		t.Error("GetWvWMatchByWorldID failed: ", err)
	}

	if _, err := api.GetWvWObjectives(); err != nil {
		t.Error("GetWvWObjectives failed: ", err)
	}

	if _, err := api.GetWvWRanks(); err != nil {
		t.Error("GetWvWRanks failed: ", err)
	}

	if _, err := api.GetWvWUpgrades(); err != nil {
		t.Error("GetWvWUpgrades failed: ", err)
	}
}
