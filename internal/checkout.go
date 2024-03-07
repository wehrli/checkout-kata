package internal

import (
	"fmt"

	"checkout-kata/internal/types"
)

// IPricing is an interface for pricing
type IPricing interface {
	ProductExists(item string) bool
	GetPrice(item string, quantity int) (int, error)
}

// Checkout is a struct that represents a checkout
type Checkout struct {
	pricing IPricing
	basket  *types.Basket
}

// NewCheckout returns a new Checkout instance
func NewCheckout(pricing IPricing) *Checkout {
	return &Checkout{
		pricing: pricing,
		basket:  &types.Basket{Products: make(map[string]*types.Product), TotalPrice: 0},
	}
}

// Scan adds a product to the basket
func (c *Checkout) Scan(item string) {
	if !c.pricing.ProductExists(item) {
		fmt.Printf("Product %s not found\n", item)
		return
	}

	product, ok := c.basket.Products[item]
	if !ok {
		c.basket.Products[item] = &types.Product{SKU: item, Quantity: 1}
		product = c.basket.Products[item]
	} else {
		product.Quantity++
	}

	cost, err := c.pricing.GetPrice(item, product.Quantity)
	if err != nil {
		fmt.Printf("Error getting price for product %s: %s\n", item, err)
		return
	}
	product.Cost = cost
}

func (c *Checkout) String() string {
	return c.basket.String()
}

// GetTotalPrice returns the total price for the basket
func (c *Checkout) GetTotalPrice() int {
	total := 0

	for _, product := range c.basket.Products {
		total += product.Cost
	}

	return total
}
