package gw2api

// Trait contains trait info
type Trait struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Icon           string         `json:"icon"`
	Description    string         `json:"description"`
	Specialization int            `json:"specialization"`
	Tier           int            `json:"tier"`
	Slot           string         `json:"slot"`
	Facts          []*Fact        `json:"facts"`
	TraitedFacts   []*FactTraited `json:"traited_facts"`
	Skills         []*Skill       `json:"skills"` // NOTE: Only id, name, description, icon, facts and traited_facts are provided
}

// Traits returns the requested traits
func (s *Session) Traits(ids ...int) (rsp []*Trait, err error) {
	err = s.get(concatStrings("/v2/traits", genArgs(ids...)), &rsp)
	return
}
