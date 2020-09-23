package gw2api

import "testing"

func TestSkills(t *testing.T) {
	api := New()
	if _, err := api.Skills(); err != nil {
		t.Errorf("Skills failed: '%s'", err)
	}
}
