package gw2api

type BuildTemplate struct {
  Name string `json:"name"`
  Profession string `json:"profession"`
  Specializations []struct {
    ID int `json:"id"`
    Traits []int `json:"traits"`
  } `json:"specializations"`
  Skills struct {
    Heal int `json:"heal"`
    Utilities []int `json:"utilities"`
    Elite int `json:"elite"`
  } `json:"skills"`
  AquaticSkills struct {
    Heal int `json:"heal"`
    Utilities []int `json:"utilities"`
    Elite int `json:"elite"`
  } `json:"aquatic_skills"`
}

type BuildTemplates struct {
  Error
  Data []*BuildTemplate
}
