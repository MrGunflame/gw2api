package gw2api

import "testing"

func TestBuild(t *testing.T) {
	api := New()
	if _, err := api.Build(); err != nil {
		t.Errorf("Build failed: '%s'", err)
	}
}
