package gw2api

// MountSkin contains information about a mount skin
type MountSkin struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Icon     string     `json:"icon"`
	Mount    string     `json:"mount"`
	DyeSlots []*DyeSlot `json:"dye_slots"`
}

// MountSkins returns the requested mount skins
func (s *Session) MountSkins(ids ...string) (rsp *[]MountSkin, err error) {
	err = s.get(concatStrings("/v2/mounts/skins", genArgsString(ids...)), &rsp)
	return
}

// MountType contains information about a mount type
type MountType struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	DefaultSkin int      `json:"default_skin"`
	Skins       []int    `json:"skins"`
	Skills      []*Skill `json:"skills"` // Only id, slot
}

// MountTypes returns the requested mount types
func (s *Session) MountTypes(ids ...string) (rsp []*MountType, err error) {
	err = s.get(concatStrings("/v2/mounts/types", genArgsString(ids...)), &rsp)
	return
}
