package gw2api

import "testing"

func TestBackStoryAnswers(t *testing.T) {
	api := New()
	if _, err := api.BackStoryAnswers(); err != nil {
		t.Errorf("BackStoryAnswers failed: '%s'", err)
	}
}

func TestBackStoryQuestions(t *testing.T) {
	api := New()
	if _, err := api.BackStoryQuestions(); err != nil {
		t.Errorf("BackStoryQuestions failed: '%s'", err)
	}
}
