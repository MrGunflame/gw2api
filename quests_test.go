package gw2api

import "testing"

func TestQuests(t *testing.T) {
	api := New()
	if _, err := api.Quests(); err != nil {
		t.Errorf("Quests failed: %s", err)
	}
}
