package types

import "fmt"

// Product is a struct that represents a product
type Product struct {
	Quantity int
	SKU      string
	Cost     int
}

// String returns a string representation of a product
func (p *Product) String() string {
	return fmt.Sprintf("SKU: %s, Quantity: %d, Cost: %d", p.SKU, p.Quantity, p.Cost)
}

// ProductDefinition is a struct that represents a product definition
type ProductDefinition struct {
	SKU          string
	UnitPrice    int
	SpecialPrice *Special
}
