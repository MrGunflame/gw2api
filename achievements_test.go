package gw2api

import (
	"os"
	"testing"
)

func TestAchievements(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.Achievements(); err != nil {
		t.Errorf("Achievements failed fetching endpoint: '%s'", err)
	}
}

func TestDailyAchievements(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.DailyAchievements(); err != nil {
		t.Errorf("DailyAchievements failed fetching endpoint: '%s'", err)
	}
}

func TestTomorrowDailyAchievements(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.TomorrowDailyAchievements(); err != nil {
		t.Errorf("TomorrowDailyAchievements failed fetching endpoint: '%s'", err)
	}
}

func TestAchievementGroups(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AchievementGroups(); err != nil {
		t.Errorf("AchievementGroups failed fetching endpoint: '%s'", err)
	}
}

func TestAchievementCategories(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.AchievementCategories(); err != nil {
		t.Errorf("AchievementCategories failed fetching endpoint: '%s'", err)
	}
}
