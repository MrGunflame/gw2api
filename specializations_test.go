package gw2api

import "testing"

func TestSpecializations(t *testing.T) {
	api := New()
	if _, err := api.Specializations(); err != nil {
		t.Errorf("Specializations failed: '%s'", err)
	}
}
