package gw2api

// BuildTemplate is a template in the accounts storage
type BuildTemplate struct {
	Name            string `json:"name"`
	Profession      string `json:"profession"`
	Specializations []struct {
		ID     int   `json:"id"`
		Traits []int `json:"traits"`
	} `json:"specializations"`
	Skills struct {
		Heal      int   `json:"heal"`
		Utilities []int `json:"utilities"`
		Elite     int   `json:"elite"`
	} `json:"skills"`
	AquaticSkills struct {
		Heal      int   `json:"heal"`
		Utilities []int `json:"utilities"`
		Elite     int   `json:"elite"`
	} `json:"aquatic_skills"`
}
