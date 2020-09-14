package gw2api

import "testing"

func TestWorldBosses(t *testing.T) {
	api := New()
	if _, err := api.WorldBosses(); err != nil {
		t.Errorf("WorldBosses failed: '%s'", err)
	}
}
