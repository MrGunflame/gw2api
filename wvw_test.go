package gw2api

import "testing"

func TestWvWAbilities(t *testing.T) {
	api := New()
	if _, err := api.WvWAbilities(); err != nil {
		t.Errorf("WvWAbilities failed: '%s'", err)
	}
}

func TestWvWMatches(t *testing.T) {
	api := New()
	if _, err := api.WvWMatches(); err != nil {
		t.Errorf("WvWmatches failed: '%s'", err)
	}
}

func TestWvWMatchByWorldID(t *testing.T) {
	api := New()
	if _, err := api.WvWMatchByWorldID(2008); err != nil {
		t.Errorf("WvWmatchByWorldID failed: '%s'", err)
	}
}

func TestWvWObjectives(t *testing.T) {
	api := New()
	if _, err := api.WvWObjectives(); err != nil {
		t.Errorf("WvWObjectives failed: '%s'", err)
	}
}

func TestWvWRanks(t *testing.T) {
	api := New()
	if _, err := api.WvWRanks(); err != nil {
		t.Errorf("WvWRanks failed: '%s'", err)
	}
}

func TestWvWUpgrades(t *testing.T) {
	api := New()
	if _, err := api.WvWUpgrades(); err != nil {
		t.Errorf("WvWUpgrades failed: '%s'", err)
	}
}
