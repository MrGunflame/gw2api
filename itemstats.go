package gw2api

type Itemstat struct {
	ID         int            `json:"id"`
	Attributes map[string]int `json:"attributes"`
}
