package internal

import (
	"fmt"

	"checkout-kata/internal/types"
)

// IPriceParser is an interface for parsing prices
type IPriceParser interface {
	ParsePrices(filePath string) (map[string]types.ProductDefinition, error)
}

// Pricing is a struct that represents the pricing
type Pricing struct {
	parser   IPriceParser
	prices   map[string]types.ProductDefinition
	filePath string
}

// NewPricing returns a new Pricing instance
func NewPricing(p IPriceParser, filePath string) *Pricing {
	return &Pricing{
		parser:   p,
		prices:   make(map[string]types.ProductDefinition),
		filePath: filePath,
	}
}

// InitPrices initializes the prices from the file
func (p *Pricing) InitPrices() error {
	prices, err := p.parser.ParsePrices(p.filePath)
	if err != nil {
		return fmt.Errorf("error parsing prices: %w", err)
	}

	p.prices = prices

	return nil
}

// ProductExists checks if a product exists
func (p *Pricing) ProductExists(item string) bool {
	_, ok := p.prices[item]
	return ok
}

// GetPrice returns the total price for a given item and quantity
func (p *Pricing) GetPrice(item string, quantity int) (int, error) {
	product, ok := p.prices[item]
	if !ok {
		return 0, types.ErrProductNotFound
	}

	if product.SpecialPrice == nil {
		return product.UnitPrice * quantity, nil
	}

	specials := quantity / product.SpecialPrice.Quantity
	remaining := quantity % product.SpecialPrice.Quantity
	totalForItem := specials*product.SpecialPrice.Price + remaining*product.UnitPrice

	return totalForItem, nil
}
