package gw2api

// Item is an game item
// type Item struct {
// 	ID         int         `json:"id"`
// 	Name       string      `json:"name"`
// 	Icon       string      `json:"icon"`
// 	Attributes *Attributes `json:"attributes"`
// }

// ItemStack is the data contained in requests that return stored items, e.g. bank or shared inventory
type ItemStack struct {
	ID        string `json:"id"`
	Count     int    `json:"count"`
	Charges   int    `json:"charges"`
	Skin      int    `json:"skin"`
	Upgrades  []int  `json:"upgrades"`
	Infusions []int  `json:"infusions"`
	Binding   string `json:"binding"`
	BoundTo   string `json:"bound_to"`
}
