package gw2api

// Profession is a game profession
type Profession struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	IconBig         string `json:"icon_big"`
	Specializations []int  `json:"specializations"`
	Training        []struct {
		ID       int    `json:"id"`
		Category string `json:"category"`
		Name     string `json:"name"`
		Track    []struct {
			Cost    int    `json:"cost"`
			Type    string `json:"type"`
			SkillID int    `json:"skill_id"` // If Type == Skill
			TraitID int    `json:"trait_id"` // If Type == Trait
		} `json:"track"`
	} `json:"training"`
	Weapons map[string]struct {
		Flags          []string `json:"flags"`
		Specialization int      `json:"specialization"`
		Skills         []struct {
			ID         int    `json:"id"`
			Slot       string `json:"slot"`
			Offhand    string `json:"offhand"`
			Attunement string `json:"attunement"` // Elementalist
			Source     string `json:"source"`     // Thief stolen skills
		} `json:"skills"`
	} `json:"weapons"`
	Flags []string `json:"flags"`
}

// Professions returns the requested game professions
func (s *Session) Professions(ids ...string) (rsp []*Profession, err error) {
	err = s.get(concatStrings("/v2/professions", genArgsString(ids...)), &rsp)
	return
}
