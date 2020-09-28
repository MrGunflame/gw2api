package gw2api

import "testing"

func TestProfessions(t *testing.T) {
	api := New()
	if _, err := api.Professions(); err != nil {
		t.Errorf("Professions failed: '%s'", err)
	}
}
