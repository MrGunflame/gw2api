package gw2api

import "strconv"

// CommerceDelivery contains the coins and items available for pickup
type CommerceDelivery struct {
	Coins int `json:"coins"`
	Items []struct {
		ID    int `json:"id"`
		Count int `json:"count"`
	} `json:"items"`
}

// CommerceDelivery returns coins and items available for pickup for the account
func (s *Session) CommerceDelivery() (delivery CommerceDelivery, err error) {
	err = s.getWithAuth("/v2/commerce/delivery", &delivery)
	return
}

// CommerceCoinGemExchange contains information about exchanging gems to coins and vice versa
type CommerceCoinGemExchange struct {
	CoinsPerGem int `json:"coins_per_gem"`
	Quantity    int `json:"quantity"`
}

// CommerceExchangeCoins returns how many gems you get for the provided amount of coins
func (s *Session) CommerceExchangeCoins(quantity int) (exchange CommerceCoinGemExchange, err error) {
	err = s.get(concatStrings("/v2/commerce/exchange/coins?quantity=", strconv.Itoa(quantity)), &exchange)
	return
}

// CommerceExchangeGems returns how many coins you get for the provided amount of gems
func (s *Session) CommerceExchangeGems(quantity int) (exchange CommerceCoinGemExchange, err error) {
	err = s.get(concatStrings("/v2/commerce/exchange/gems?quantity=", strconv.Itoa(quantity)), &exchange)
	return
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

// CommerceListings returns all buy and sell orders for the items
func (s *Session) CommerceListings(ids ...int) (items []*CommerceListing, err error) {
	err = s.get(concatStrings("/v2/commerce/listings", genArgs(ids...)), &items)
	return
}

// CommerceListingsIds returns ids of all items available on trading post
func (s *Session) CommerceListingsIds() (ids []int, err error) {
	err = s.get("/v2/commerce/listings", &ids)
	return
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

// CommercePrices returns the current tp prices for the requested ids
func (s *Session) CommercePrices(ids ...int) (st []*CommercePrice, err error) {
	err = s.get(concatStrings("/v2/commerce/prices", genArgs(ids...)), &st)
	return
}

// CommerceTransaction is an item waiting for exchanging or exchanged in the trading post
type CommerceTransaction struct {
	ID        int    `json:"id"`
	ItemID    int    `json:"item_id"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Created   string `json:"created"`
	Purchased string `json:"purchased"` // Only available from history
}

// CommerceTransactionsCurrent returns the accounts current transactions
func (s *Session) CommerceTransactionsCurrent() (transactions []*CommerceTransaction, err error) {
	err = s.getWithAuth("/v2/commerce/transactions/current", &transactions)
	return
}

// CommerceTransactionsHistory returns the accounts transaction history
func (s *Session) CommerceTransactionsHistory() (transactions []*CommerceTransaction, err error) {
	err = s.getWithAuth("/v2/commerce/transactions/history", &transactions)
	return
}
