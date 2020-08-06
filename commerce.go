package gw2api

import "strconv"

// CommerceDelivery contains the coins and items avaliable for pickup
type CommerceDelivery struct {
	Coins int `json:"coins"`
	Items []struct {
		ID    int `json:"id"`
		Count int `json:"count"`
	} `json:"items"`
}

// CommerceCoinGemExchange contains information about exchanging gems to coins and vice versa
type CommerceCoinGemExchange struct {
	CoinsPerGem int `json:"coins_per_gem"`
	Quantity    int `json:"quantity"`
}

// CommerceListing contains an items buy and sell orders
type CommerceListing struct {
	ID   int `json:"id"`
	Buys []struct {
		Listings  int `json:"listings"`
		UnitPrice int `json:"unit_price"`
		Quantity  int `json:"quantity"`
	} `json:"buys"`
	Sells []struct {
		Listings  int `json:"listings"`
		UnitPrice int `json:"unit_price"`
		Quantity  int `json:"quantity"`
	} `json:"sells"`
}

// CommercePrice contains the currently highest buy price and lowest sell price
type CommercePrice struct {
	ID          int  `json:"id"`          // The item id
	Whitelisted bool `json:"whitelisted"` // Whether free to play accounts may purchase/sell this item
	Buys        struct {
		Quantity  int `json:"quantity"`
		UnitPrice int `json:"unit_price"`
	} `json:"buys"`
	Sells struct {
		Quantity  int `json:"quantity"`
		UnitPrice int `json:"unit_price"`
	} `json:"sells"`
}

// CommerceTransaction is an item waiting for exchanging or exchanged in the trading post
type CommerceTransaction struct {
	ID        int    `json:"id"`
	ItemID    int    `json:"item_id"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Created   string `json:"created"`
	Purchased string `json:"purchased"` // Only avaliable from history
}

// GetCommerceDelivery returns coins and items avaiable for pickup for the account
func (s *Session) GetCommerceDelivery() (delivery CommerceDelivery, err error) {
	err = s.getWithAuth("/v2/commerce/delivery", &delivery)
	return
}

// GetCommerceExchangeCoins returns how many gems you get for the provided amount of coins
func (s *Session) GetCommerceExchangeCoins(quantity int) (exchange CommerceCoinGemExchange, err error) {
	err = s.get(concatStrings("/v2/commerce/exchange/coins?quantity=", strconv.Itoa(quantity)), &exchange)
	return
}

// GetCommerceExchangeGems returns how many coins you get for the provided amount of gems
func (s *Session) GetCommerceExchangeGems(quantity int) (exchange CommerceCoinGemExchange, err error) {
	err = s.get(concatStrings("/v2/commerce/exchange/gems?quantity=", strconv.Itoa(quantity)), &exchange)
	return
}

// GetCommerceListings returns all buy and sell orders for the items
func (s *Session) GetCommerceListings(ids ...int) (items []*CommerceListing, err error) {
	err = s.get(concatStrings("/v2/commerce/listings", genArgs(ids...)), &items)
	return
}

// GetCommerceTransactionsCurrent returns the accounts current transactions
func (s *Session) GetCommerceTransactionsCurrent() (transactions []*CommerceTransaction, err error) {
	err = s.getWithAuth("/v2/commerce/transactions/current", &transactions)
	return
}

// GetCommerceTransactionsHistory returns the accounts transaction history
func (s *Session) GetCommerceTransactionsHistory() (transactions []*CommerceTransaction, err error) {
	err = s.getWithAuth("/v2/commerce/transactions/history", &transactions)
	return
}
