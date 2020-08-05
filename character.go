package gw2api

type CharacterCore struct {
	Name       string `json:"name"`
	Race       string `json:"race"`
	Gender     string `json:"gender"`
	Profession string `json:"profession"`
	Level      int    `json:"level"`
	Guild      string `json:"guild"`
	Age        int    `jsonb:"age"`
	Created    string `json:"created"`
	Deaths     int    `json:"deaths"`
	Title      int    `json:"title"`
}

type CharacterCrafting struct {
	Discipline string `json:"discipline"`
	Rating     int    `json:"rating"`
	Active     bool   `json:"active"`
}

// type CharacterEquipment struct {
// 	ID int `json:"id"`
// 	Slot string `json:"slot"`
// 	Infusions []int `json:"infusions"`
// }

// GetCharacters returns all character names
func (s *Session) GetCharacters() (chars []string, err error) {
	err = s.getWithAuth("/v2/characters", &chars)
	return
}

// GetCharacterBackStory returns a characters initial story decisions
func (s *Session) GetCharacterBackStory(character string) (backstory []string, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/backstory"), &backstory)
	return
}

// GetCharacterCore returns a characters core information
func (s *Session) GetCharacterCore(character string) (core CharacterCore, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/core"), &core)
	return
}
