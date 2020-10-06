package gw2api

import (
	"testing"
)

func TestFinishers(t *testing.T) {
	api := New()
	if _, err := api.Finishers(); err != nil {
		t.Errorf("Finishers failed: '%s'", err)
	}
}
