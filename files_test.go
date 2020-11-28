package gw2api

import "testing"

func TestFiles(t *testing.T) {
	api := New()
	if _, err := api.Files(); err != nil {
		t.Errorf("Files failed: %s", err)
	}
}
