package gw2api

// Skill contains skill info
type Skill struct {
	ID              int            `json:"id"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Icon            string         `json:"icon"`
	ChatLink        string         `json:"chat_link"`
	Type            string         `json:"type"` // Bundle, Elite, Heal, Profession, Utility, Weapon
	WeaponType      string         `json:"weapon_type"`
	Professions     []string       `json:"professions"`
	Slot            string         `json:"slot"`
	Facts           []*Fact        `json:"facts"`
	TraitedFacts    []*FactTraited `json:"traited_facts"`
	Categories      []string       `json:"categories"`
	Attunement      string         `json:"attunement"`
	Cost            int            `json:"cost"`
	DualWield       string         `json:"dual_wield"`
	FlipSkill       int            `json:"flip_skill"`
	Initiative      int            `json:"initiative"`
	NextChain       int            `json:"next_chain"`
	PrevChain       int            `json:"prev_chain"`
	TransformSkills []int          `json:"transform_skills"`
	BundleSkills    []int          `json:"bundle_skills"`
	ToolbeltSkill   int            `json:"toolbelt_skill"`
}

// Skills returns the requested skills
func (s *Session) Skills(ids ...int) (rsp []*Skill, err error) {
	err = s.get(concatStrings("/v2/skills", genArgs(ids...)), &rsp)
	return
}
