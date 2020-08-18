package gw2api

import "testing"

func TestPvP(t *testing.T) {
	api := New()
	apiAuth := New().WithAccessToken(testAccessToken)

	if _, err := apiAuth.GetPvPStats(); err != nil {
		t.Error("GetPvPStats failed: ", err)
	}

	if _, err := apiAuth.GetPvPGames(); err != nil {
		t.Error("GetPvPGames failed: ", err)
	}

	if _, err := apiAuth.GetPvPStandings(); err != nil {
		t.Error("GetPvPStandings failed: ", err)
	}

	if _, err := api.GetPvPAmulets(); err != nil {
		t.Error("GetPvPAmulets failed: ", err)
	}

	if _, err := api.GetPvPRanks(); err != nil {
		t.Error("GetPvPRanks failed: ", err)
	}

	if _, err := api.GetPvPSeasons(); err != nil {
		t.Error("GetPvPSeasons failed: ", err)
	}

	if _, err := api.GetPvPSeasonLeaderboards("088D9941-21E7-4335-A0C2-4365A8D46B1F"); err != nil {
		t.Error("GetPvPSeasonLeaderboards failed: ", err)
	}
}
