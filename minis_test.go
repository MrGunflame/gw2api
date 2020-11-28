package gw2api

import "testing"

func TestMinis(t *testing.T) {
	api := New()
	if _, err := api.Minis(); err != nil {
		t.Errorf("Minis failed: %s", err)
	}
}
