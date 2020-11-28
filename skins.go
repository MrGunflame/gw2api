package gw2api

// Skin is a game item skin
type Skin struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Flags        []string `json:"flags"`
	Restrictions []string `json:"restrictions"`
	Icon         string   `json:"icon"`
	Rarity       string   `json:"rarity"`
	Description  string   `json:"description"`
	Details      struct {
		Type        string `json:"type"`
		WeightClass string `json:"weight_class"`
		DyeSlots    struct {
			Default []struct {
				ColorID  int    `json:"color_id"`
				Material string `json:"material"`
			} `json:"default"`
		} `json:"dye_slots"`
		Overrides  map[string]string `json:"overrides"`
		DamageType string            `json:"damage_type"`
	} `json:"details"`
}

// Skins returns the skins with the given ids
func (s *Session) Skins(ids ...int) (resp []*Skin, err error) {
	err = s.get(concatStrings("/v2/skins", genArgs(ids...)), &resp)
	return
}
