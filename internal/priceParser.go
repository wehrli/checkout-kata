package internal

import (
	"os"

	"gopkg.in/yaml.v2"

	"checkout-kata/internal/types"
)

// Special is a struct that represents a special price
type Special struct {
	Quantity int `yaml:"quantity"`
	Price    int `yaml:"price"`
}

// Product is a struct that represents a product
type Product struct {
	SKU          string   `yaml:"sku"`
	UnitPrice    int      `yaml:"unit_price"`
	SpecialPrice *Special `yaml:"special"`
}

type pricesConfig struct {
	Products []Product `yaml:"prices"`
}

// Parser is a struct that represents a parser
type Parser struct{}

// ParsePrices parses the prices from a file
func (p *Parser) ParsePrices(filePath string) (map[string]types.ProductDefinition, error) {
	var config pricesConfig
	prices := make(map[string]types.ProductDefinition)

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	for _, p := range config.Products {
		var s types.Special
		if p.SpecialPrice != nil {
			s = types.Special{
				Quantity: p.SpecialPrice.Quantity,
				Price:    p.SpecialPrice.Price,
			}
		}

		prices[p.SKU] = types.ProductDefinition{
			SKU:          p.SKU,
			UnitPrice:    p.UnitPrice,
			SpecialPrice: &s,
		}
	}

	return prices, nil
}
