package gw2api

import "strconv"

// Recipe is a game recipe
type Recipe struct {
	ID              int      `json:"id"`
	Type            string   `json:"type"`
	OutputItemID    int      `json:"output_item_id"`
	OutputItemCount int      `json:"output_item_count"`
	TimeToCraftMS   int      `json:"time_to_craft_ms"`
	Disciplines     []string `json:"disciplines"`
	MinRating       int      `json:"min_rating"`
	Flags           []string `json:"flags"`
	Ingredients     []struct {
		ItemID int `json:"item_id"`
		Count  int `json:"count"`
	} `json:"ingrdients"`
	GuildIngredients []struct {
		UpgradeID int `json:"upgrade_id"`
		Count     int `json:"count"`
	} `json:"guild_ingredients"`
	OutputUpgradeID int    `json:"output_upgrade_id"`
	ChatLink        string `json:"chat_link"`
}

// Recipes returns the recipes with the requested ids
func (s *Session) Recipes(ids ...int) (resp []*Recipe, err error) {
	err = s.get(concatStrings("/v2/recipes", genArgs(ids...)), &resp)
	return
}

// RecipesSearchInput searches for recipes that use the item as an input ingredient
func (s *Session) RecipesSearchInput(id int) (resp []int, err error) {
	err = s.get(concatStrings("/v2/recipes/search?input=", strconv.Itoa(id)), &resp)
	return
}

// RecipesSearchOutput searches for recipes that craft the item
func (s *Session) RecipesSearchOutput(id int) (resp []int, err error) {
	err = s.get(concatStrings("/v2/recipes/search?output=", strconv.Itoa(id)), &resp)
	return
}
