package gw2api

import (
	"testing"
)

func TestAchievements(t *testing.T) {
	api := New()

	if _, err := api.GetAchievements(1); err != nil {
		t.Error("GetAchievements failed: ", err)
	}

	if _, err := api.GetDailyAchievements(); err != nil {
		t.Error("GetDailyAchievements failed: ", err)
	}

	if _, err := api.GetTomorrowDailyAchievements(); err != nil {
		t.Error("GetTomorrowDailyAchievements failed: ", err)
	}

	if _, err := api.GetAchievementGroups(); err != nil {
		t.Error("GetAchievementGroups failed: ", err)
	}

	if _, err := api.GetAchievementCategories(); err != nil {
		t.Error("GetAchievementCategories failed: ", err)
	}

}
