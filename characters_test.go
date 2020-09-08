package gw2api

import (
	"os"
	"testing"
)

func TestCharacters(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.Characters(); err != nil {
		t.Errorf("Characters failed: '%s'", err)
	}
}

func TestCharacterBackstory(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	char := os.Getenv("CHARACTER")
	if char == "" {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterBackstory(char); err != nil {
		t.Errorf("CharacterBackstory failed: '%s'", err)
	}
}

func TestCharacterCore(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	char := os.Getenv("CHARACTER")
	if char == "" {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterCore(char); err != nil {
		t.Errorf("CharacterCore failed: '%s'", err)
	}
}

func TestCharacterCrafting(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	char := os.Getenv("CHARACTER")
	if char == "" {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterCrafting(char); err != nil {
		t.Errorf("CharacterCrafting failed: '%s'", err)
	}
}

func TestCharacterEquipment(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	char := os.Getenv("CHARACTER")
	if len(char) < 3 {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterEquipment(char); err != nil {
		t.Errorf("CharacterEquipment failed: '%s'", err)
	}
}

func TestCharacterHeropoints(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	char := os.Getenv("CHARACTER")
	if len(char) < 3 {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterHeropoints(char); err != nil {
		t.Errorf("CharacterHeropoints failed: '%s'", err)
	}
}

func TestCharacterInventory(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	char := os.Getenv("CHARACTER")
	if len(char) < 3 {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterInventory(char); err != nil {
		t.Errorf("CharacterInventory failed: '%s'", err)
	}
}

func TestCharacterRecipes(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	char := os.Getenv("CHARACTER")
	if len(char) < 3 {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterRecipes(char); err != nil {
		t.Errorf("CharacterRecipes failed: '%s'", err)
	}
}

func TestCharacterSab(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	char := os.Getenv("CHARACTER")
	if len(char) < 3 {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterSab(char); err != nil {
		t.Errorf("CharacterSab failed: '%s'", err)
	}
}

func TestCharacterSkills(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	char := os.Getenv("CHARACTER")
	if len(char) < 3 {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterSkills(char); err != nil {
		t.Errorf("CharacterSkills failed: '%s'", err)
	}
}

func TestCharacterSpecializations(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	char := os.Getenv("CHARACTER")
	if len(char) < 3 {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterSpecializations(char); err != nil {
		t.Errorf("CharacterSpecializations failed: '%s'", err)
	}
}

func TestCharacterTraining(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	char := os.Getenv("CHARACTER")
	if len(char) < 3 {
		t.Skip("Unable to test without CHARACTER")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CharacterTraining(char); err != nil {
		t.Errorf("CharacterTraining failed: '%s'", err)
	}
}
