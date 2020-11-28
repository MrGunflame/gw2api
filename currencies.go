package gw2api

// Currency is a game wallet currency
type Currency struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Order       int    `json:"order"`
}

// Currencies returns the currencies with the given ids
func (s *Session) Currencies(ids ...int) (resp []*Currency, err error) {
	err = s.getWithLang(concatStrings("/v2/currencies", genArgs(ids...)), &resp)
	return
}
