package gw2api

// Item is a game item
type Item struct {
	ID           int      `json:"id"`
	ChatLink     string   `json:"chat_link"`
	Name         string   `json:"name"`
	Icon         string   `json:"icon"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	Rarity       string   `json:"rarity"`
	Level        int      `json:"level"`
	VendorValue  int      `json:"vendor_value"`
	DefaultSkin  int      `json:"default_skin"`
	Flags        []string `json:"flags"`
	GameTypes    []string `json:"game_types"`
	Restrictions []string `json:"restrictions"`
	UpgradesInto []struct {
		Upgrade string `json:"upgrade"`
		ItemID  int    `json:"item_id"`
	} `json:"upgrades_into"`
	UpgradesFrom []struct {
		Upgrade string `json:"upgrade"`
		ItemID  int    `json:"item_id"`
	} `json:"upgrades_from"`
	Details struct {
		Type          string `json:"type"`         // Armor / Consumable / Container / Gathering tools / Gizmo / Salvage kits / Trinket / Upgrade component / Weapon
		WeightClass   string `json:"weight_class"` // Armor
		Defense       int    `json:"defense"`      // Armor / Weapon
		InfusionSlots []struct {
			Flags  []string `json:"flags"`
			ItemID int      `json:"item_id"`
		} `json:"infusion_slots"` // Armor / Back item / Trinket / Weapon
		AttributeAdjustment float64 `json:"attribute_adjustment"` // Armor / Back item / Trinket / Weapon
		InfixUpgrade        struct {
			ID         int `json:"id"`
			Attributes []struct {
				Attribute string `json:"attribute"`
				Modifier  int    `json:"modifier"`
			}
			Buff struct {
				SkillID     int    `json:"skill_id"`
				Description string `json:"description"`
			} `json:"buff"`
		} `json:"infix_upgrade"` // Armor / Back item / Trinket / Upgrade component / Weapon
		SuffixItemID          int      `json:"suffix_item_id"`           // Armor / Back item / Trinket / Weapon
		SecondarySuffixItemID string   `json:"secondary_suffix_item_id"` // Armor / Back item / Trinket / Weapon
		StatChoices           []int    `json:"stat_choices"`             // Armor / Back item / Trinket / Weapon
		Size                  int      `json:"size"`                     // Bag
		NoSellOrSort          bool     `json:"no_sell_or_sort"`          // Bag
		Description           string   `json:"description"`              // Consumable
		DurationMS            int      `json:"duration_ms"`              // Consumable
		UnlockType            string   `json:"unlock_type"`              // Consumable
		ColorID               int      `json:"color_id"`                 // Consumable
		RecipeID              int      `json:"recipe_id"`                // Consumable
		ExtraRecipeIDs        []int    `json:"extra_recipe_ids"`         // Consumable
		GuildUpgradeID        int      `json:"guild_upgrade_id"`         // Consumable / Gizmo
		ApplyCount            int      `json:"apply_count"`              // Consumable
		Name                  string   `json:"name"`                     // Consumable
		Icon                  string   `json:"icon"`                     // Consumable
		Skins                 []int    `json:"skins"`                    // Consumable
		VendorIDs             []int    `json:"vendor_ids"`               // Gizmo
		MinipetID             int      `json:"minipet_id"`               // Miniature
		Charges               int      `json:"charges"`                  // Salvage kits
		Flags                 []string `json:"flags"`                    // Upgrade component
		InfusionUpgradesFlags []string `json:"infusion_upgrades_flags"`  // Upgrade component
		Suffix                string   `json:"suffix"`                   // Upgrade component
		Bonuses               []string `json:"bonuses"`                  // Upgrade component
		DamageType            string   `json:"damage_type"`              // Weapon
		MinPower              int      `json:"min_power"`                // Weapon
		MaxPower              int      `json:"max_power"`                // Weapon
	} `json:"details"`
}

// Items returns the items with the given id
func (s *Session) Items(ids ...int) (rsp []*Item, err error) {
	err = s.get(concatStrings("/v2/items", genArgs(ids...)), &rsp)
	return
}
