package test

import (
	"checkout-kata/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {

	t.Run("OK - Scan went well", func(t *testing.T) {
		priceParser := internal.Parser{}
		pricing := internal.NewPricing(&priceParser, "./products_prices.yaml")
		err := pricing.InitPrices()
		if err != nil {
			t.Errorf("Error initializing prices: %s", err)
		}
		checkout := internal.NewCheckout(pricing)

		checkout.Scan("A")
		tt := checkout.GetTotalPrice()
		assert.Equal(t, 50, tt, "Unexpected result")
	})

	t.Run("OK - Scan when a product has a special price", func(t *testing.T) {
		priceParser := internal.Parser{}
		pricing := internal.NewPricing(&priceParser, "./products_prices.yaml")
		err := pricing.InitPrices()
		if err != nil {
			t.Errorf("Error initializing prices: %s", err)
		}
		checkout := internal.NewCheckout(pricing)

		checkout.Scan("A")
		checkout.Scan("A")
		checkout.Scan("A")

		tt := checkout.GetTotalPrice()
		assert.Equal(t, 130, tt, "Unexpected result")
	})

	t.Run("OK - When product does not exist", func(t *testing.T) {
		priceParser := internal.Parser{}
		pricing := internal.NewPricing(&priceParser, "./products_prices.yaml")
		err := pricing.InitPrices()
		if err != nil {
			t.Errorf("Error initializing prices: %s", err)
		}
		checkout := internal.NewCheckout(pricing)

		checkout.Scan("Z")
		tt := checkout.GetTotalPrice()
		assert.Equal(t, 0, tt, "Unexpected result")
	})

}
