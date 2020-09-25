package gw2api

// DyeSlot contains information about a dye used in a dye slot
type DyeSlot struct {
	ColorID  int    `json:"color_id"`
	Material string `json:"material"`
}
