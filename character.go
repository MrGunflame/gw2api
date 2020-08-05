package gw2api

// Character contains all the listed items from subrequests
type Character struct {
	CharacterCore
	Backstory []string `json:"backstory"`
	CharacterExtras
	Specializations []*Character
	CharacterSkills
	CharacterEquipment
	Crafting []*DisciplineProgress `json:"crafting"`
	Recipes  []int                 `json:"recipes"`
	Bags     []*InventoryBag       `json:"bags"`
	Training []*TrainingTree       `json:"training"`
}

// CharacterCore contains general character information
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

// DisciplineProgrss defines a crafting discipline and the characters progress
type DisciplineProgress struct {
	Discipline string `json:"discipline"`
	Rating     int    `json:"rating"`
	Active     bool   `json:"active"`
}

// CharacterEquipment defines an equipment item carried by an character
type CharacterEquipment struct {
	ID        int      `json:"id"`
	Slot      string   `json:"slot"`
	Infusions []int    `json:"infusions"`
	Upgrades  []int    `json:"upgrades"`
	Skin      int      `json:"skin"`
	Stats     Itemstat `json:"stats"`
	Binding   string   `json:"binding"`
	Charges   int      `json:"charges"`
	BoundTo   string   `json:"bound_to"`
	Dyes      []int    `json:"dyes"`
}

// InventoryBag is a bag placed in the inventory bag slots. It also contains the items inside
type InventoryBag struct {
	ID        int         `json:"id"`
	Size      int         `json:"size"`
	Inventory []ItemStack `json:"inventory"`
}

// SuperAdventureBox contains the character's super adventure box progression
type SuperAdventureBox struct {
	Zones []struct {
		ID    int    `json:"id"`
		Mode  string `json:"mode"`
		World int    `json:"world"`
		Zone  int    `json:"zone"`
	} `json:"zones"`
	Unlocks []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"unlocks"`
	Songs []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"songs"`
}

// CharacterSkills contains the skills for all gamemodes
type CharacterSkills struct {
	PvE Skills `json:"pve"`
	PvP Skills `json:"pvp"`
	WvW Skills `json:"wvw"`
}

// Skills defines a set of character skills
type Skills struct {
	Heal      int   `json:"heal"`
	Utilities []int `json:"utilities"`
	Elite     int   `json:"elite"`
	Legends   int   `json:"legends"`
}

// CharacterSpecializations contains the specializations for all gamemodes
type CharacterSpecializations struct {
	PvE []*Specialization `json:"pve"`
	PvP []*Specialization `json:"pvp"`
	WvW []*Specialization `json:"wvw"`
}

// Specialization defines a specialization and its trait choices
type Specialization struct {
	ID     int   `json:"id"`
	Traits []int `json:"traits"`
}

// CharacterExtras contains special data contained in the general request
type CharacterExtras struct {
	WvWAbilities []struct {
		ID   int `json:"id"`
		Rank int `json:"rank"`
	} `json:"wvw_abilities"`
	EquipmentPvP struct {
		Amulet int   `json:"amulet"`
		Rune   int   `json:"rune"`
		Sigils []int `json:"sigils"`
	} `json:"equipment_pvp`
	Flags []string `json:"flag"`
}

// TrainingTree defines a training tree for both skills and traits as displayed in the game
type TrainingTree struct {
	ID    int  `json:"id"`
	Spent int  `json:"spent"`
	Done  bool `json:"done"`
}

// GetCharacters returns all character names
func (s *Session) GetCharacters() (chars []string, err error) {
	err = s.getWithAuth("/v2/characters", &chars)
	return
}

// GetCharacterBackStory returns the characters initial story decisions
func (s *Session) GetCharacterBackstory(character string) (backstory []string, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/backstory"), &backstory)
	return
}

// GetCharacterCore returns the characters core information
func (s *Session) GetCharacterCore(character string) (core CharacterCore, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/core"), &core)
	return
}

// GetCharacterCrafting returns the characters crafting disciplines
func (s *Session) GetCharacterCrafting(character string) (crafting []*DisciplineProgress, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/crafting"), &crafting)
	return
}

// GetCharacterEquipment returns the character equipment items
func (s *Session) GetCharacterEquipment(character string) (equipment []*CharacterEquipment, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/equipment"), &equipment)
	return
}

// GetCharacterHeropoints returns the characters unlocked hero challenges
func (s *Session) GetCharacterHeropoints(character string) (heropoints []string, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/heropoints"), &heropoints)
	return
}

// GetCharacterInventory returns the characters inventory bags and its contents
func (s *Session) GetCharacterInventory(character string) (inventory []*InventoryBag, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/inventory"), &inventory)
	return
}

// GetCharacterRecipes returns the characters unlocked recipes
func (s *Session) GetCharacterRecipes(character string) (recipes []int, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/recipes"), &recipes)
	return
}

// GetCharacterSab returns the characters super adventure box progress
func (s *Session) GetCharacterSab(character string) (sab SuperAdventureBox, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters", character, "/sab"), &sab)
	return
}

// GetCharacterSkills returns the characters equipped skills
func (s *Session) GetCharacterSkills(character string) (skills CharacterSkills, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/skills"), &skills)
	return
}

// GetCharacterSpecializations returns the characters equipped specializations
func (s *Session) GetCharacterSpecializations(character string) (specializations CharacterSpecializations, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/specializations"), &specializations)
	return
}

// GetCharacterTraining returns the characters training progress
func (s *Session) GetCharacterTraining(character string) (training []*TrainingTree, err error) {
	err = s.getWithAuth(concatStrings("/v2/characters/", character, "/training"), &training)
	return
}
