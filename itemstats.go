package gw2api

// Itemstat contains the stattype and attributes of an item
type Itemstat struct {
	ID         int            `json:"id"`
	Attributes map[string]int `json:"attributes"`
}
