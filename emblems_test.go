package gw2api

import "testing"

func TestEmblemBackground(t *testing.T) {
	api := New()
	if _, err := api.EmblemBackgrounds(); err != nil {
		t.Errorf("EmblemBackgrounds failed: '%s'", err)
	}
}

func TestEmblemForegrounds(t *testing.T) {
	api := New()
	if _, err := api.EmblemForegrounds(); err != nil {
		t.Errorf("EmblemForegrounds failed: '%s'", err)
	}
}
