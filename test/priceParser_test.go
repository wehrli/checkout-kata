package test

import (
	"checkout-kata/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriceParser(t *testing.T) {

	t.Run("OK - Parse price", func(t *testing.T) {
		parser := internal.Parser{}
		prices, err := parser.ParsePrices("./products_prices.yaml")
		if err != nil {
			t.Errorf("Error parsing prices: %s", err)
		}

		assert.Equal(t, 4, len(prices), "Unexpected result")
		assert.Equal(t, 50, prices["A"].UnitPrice, "Unexpected result")
		assert.Equal(t, 30, prices["B"].UnitPrice, "Unexpected result")
		assert.Equal(t, 20, prices["C"].UnitPrice, "Unexpected result")
		assert.Equal(t, 15, prices["D"].UnitPrice, "Unexpected result")
	})

	t.Run("OK - Parse price - file does not exist", func(t *testing.T) {
		parser := internal.Parser{}
		prices, err := parser.ParsePrices("./not_found.yaml")
		assert.NotNil(t, err, "Unexpected error")
		assert.Nil(t, prices, "Unexpected result")
	})
}
