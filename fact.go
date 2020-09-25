package gw2api

// Fact contains information about a skills/trait effect
type Fact struct {
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

// FactTraited contains skill effects that are overriding the default skill/trait effects
type FactTraited struct {
	*Fact
	RequiredTrait int `json:"required_trait"`
	Overrides     int `json:"overrides"`
}
