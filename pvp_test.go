package gw2api

import (
	"os"
	"testing"
)

func TestPvPStats(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	api := New().WithAccessToken(key)
	if _, err := api.PvPStats(); err != nil {
		t.Errorf("PvPStats failed: '%s'", err)
	}
}

func TestPvPGames(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	api := New()
	if _, err := api.PvPGames(); err != nil {
		t.Errorf("PvPGames failed: '%s'", err)
	}
}

func TestPvPStandings(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	api := New()
	if _, err := api.PvPStandings(); err != nil {
		t.Errorf("PvPStandings failed: '%s'", err)
	}
}

func TestPvPRanks(t *testing.T) {
	api := New()
	if _, err := api.PvPRanks(); err != nil {
		t.Errorf("PvPRanks failed: '%s'", err)
	}
}

func TestPvPSeasons(t *testing.T) {
	api := New()
	if _, err := api.PvPSeasons(); err != nil {
		t.Errorf("PvPSeasons failed: '%s'", err)
	}
}

func TestPvPSeasonLeaderboards(t *testing.T) {
	api := New()
	if _, err := api.PvPSeasonLeaderboards("020A557D-387A-4A6F-904B-D0A0DBEE68FB", "eu"); err != nil {
		t.Errorf("PvPSeasonLeaderboards failed: '%s'", err)
	}
}
