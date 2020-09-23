package gw2api

// Skill contains skill info
type Skill struct {
	ID              int                `json:"id"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	Icon            string             `json:"icon"`
	ChatLink        string             `json:"chat_link"`
	Type            string             `json:"type"` // Bundle, Elite, Heal, Profession, Utility, Weapon
	WeaponType      string             `json:"weapon_type"`
	Professions     []string           `json:"professions"`
	Slot            string             `json:"slot"`
	Facts           []SkillFact        `json:"facts"`
	TraitedFacts    []SkillFactTraited `json:"traited_facts"`
	Categories      []string           `json:"categories"`
	Attunement      string             `json:"attunement"`
	Cost            int                `json:"cost"`
	DualWield       string             `json:"dual_wield"`
	FlipSkill       int                `json:"flip_skill"`
	Initiative      int                `json:"initiative"`
	NextChain       int                `json:"next_chain"`
	PrevChain       int                `json:"prev_chain"`
	TransformSkills []int              `json:"transform_skills"`
	BundleSkills    []int              `json:"bundle_skills"`
	ToolbeltSkill   int                `json:"toolbelt_skill"`
}

// SkillFact contains information about a skills effect
type SkillFact struct {
	Text          string      `json:"text"`
	Icon          string      `json:"icon"`
	Type          string      `json:"type"`
	Value         interface{} `json:"value"`             // Type == AttributeAdjust || Number || Recharge || Unblockable (bool)
	Target        string      `json:"target"`            // Type == AttributeAdjust
	Status        string      `json:"status"`            // Type == Buff
	Description   string      `json:"description"`       // Type == Buff
	ApplyCount    int         `json:"apply_count"`       // Type == Buff
	Duration      int         `json:"duration"`          // Type == Buff || Duration || Time
	FieldType     string      `json:"field_type"`        // Type == ComboField
	FinisherType  string      `json:"finisher_type"`     // Type == ComboFinisher
	Percent       float64     `json:"percent"`           // Type == ComboFinisher || Percent
	HitCount      int         `json:"hit_count"`         // Type == Damage || Heal
	DmgMultiplier int         `json:"damage_multiplier"` // Type == Damage
	Distance      int         `json:"distance"`          // Type == Distance || Radius
	Prefix        struct {
		Text        string `json:"text"`
		Icon        string `json:"icon"`
		Status      string `json:"status"`
		Description string `json:"description"`
	} `json:"prefix"` // Type == PrefixedBuff
}

// SkillFactTraited contains skill effects that are overriding the default skill effects
type SkillFactTraited struct {
	*SkillFact
	RequiredTrait int `json:"required_trait"`
	Overrides     int `json:"overrides"`
}

// Skills returns the requested skills
func (s *Session) Skills(ids ...int) (rsp []*Skill, err error) {
	err = s.get(concatStrings("/v2/skills", genArgs(ids...)), &rsp)
	return
}
