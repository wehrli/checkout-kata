package test

import (
	"checkout-kata/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPricing(t *testing.T) {

	t.Run("OK - Product exist - return true", func(t *testing.T) {
		parser := internal.Parser{}
		pricing := internal.NewPricing(&parser, "./products_prices.yaml")
		err := pricing.InitPrices()
		if err != nil {
			t.Errorf("Error initializing prices: %s", err)
		}

		result := pricing.ProductExists("A")
		assert.True(t, result, "Unexpected result")
	})

	t.Run("OK - Product exist - return false", func(t *testing.T) {
		parser := internal.Parser{}
		pricing := internal.NewPricing(&parser, "./products_prices.yaml")
		err := pricing.InitPrices()
		if err != nil {
			t.Errorf("Error initializing prices: %s", err)
		}

		result := pricing.ProductExists("Z")
		assert.False(t, result, "Unexpected result")
	})

	t.Run("OK - Get Price", func(t *testing.T) {
		parser := internal.Parser{}
		pricing := internal.NewPricing(&parser, "./products_prices.yaml")
		err := pricing.InitPrices()
		if err != nil {
			t.Errorf("Error initializing prices: %s", err)
		}

		result, err := pricing.GetPrice("A", 3)
		assert.Nil(t, err, "Unexpected error")
		assert.Equal(t, 130, result, "Unexpected result")
	})

	t.Run("OK - Get Price - Product does not exist", func(t *testing.T) {
		parser := internal.Parser{}
		pricing := internal.NewPricing(&parser, "./products_prices.yaml")
		err := pricing.InitPrices()
		if err != nil {
			t.Errorf("Error initializing prices: %s", err)
		}

		result, err := pricing.GetPrice("Z", 3)
		assert.NotNil(t, err, "Unexpected error")
		assert.Equal(t, 0, result, "Unexpected result")
	})

}
