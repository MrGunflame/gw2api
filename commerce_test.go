package gw2api

import (
	"os"
	"testing"
)

func TestCommerceDelivery(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CommerceDelivery(); err != nil {
		t.Errorf("Commerce Delivery failed: '%s'", err)
	}
}

func TestCommerceExchangeCoins(t *testing.T) {
	api := New()
	if _, err := api.CommerceExchangeCoins(10000); err != nil {
		t.Errorf("CommerceExchangeCoins failed: '%s'", err)
	}
}

func TestCommerceExchangeGems(t *testing.T) {
	api := New()
	if _, err := api.CommerceExchangeGems(100); err != nil {
		t.Errorf("CommerceExchangeGems failed: '%s'", err)
	}
}

func TestCommerceListings(t *testing.T) {
	api := New()
	if _, err := api.CommerceListings(104); err != nil {
		t.Errorf("CommerceListings failed: '%s'", err)
	}
}

func TestCommerceListingsIds(t *testing.T) {
	api := New()
	if _, err := api.CommerceListingsIds(); err != nil {
		t.Errorf("CommerceListingsIds failed: '%s'", err)
	}
}

func TestCommercePrices(t *testing.T) {
	api := New()
	if _, err := api.CommercePrices(104); err != nil {
		t.Errorf("CommercePrices failed: '%s'", err)
	}
}

func TestCommerceTransactionsCurrent(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CommerceTransactionsCurrent(); err != nil {
		t.Errorf("CommerceTransactionsCurrent failed: '%s'", err)
	}
}

func TestCommerceTransactionHistory(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKey")
	}

	api := New().WithAccessToken(key)
	if _, err := api.CommerceTransactionsHistory(); err != nil {
		t.Errorf("CommerceTransactionsHistory failed: '%s'", err)
	}
}
