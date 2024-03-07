package types

// Basket is a struct that represents a basket
type Basket struct {
	Products   map[string]*Product
	TotalPrice int
}

// String returns a string representation of a basket
func (b *Basket) String() string {
	if len(b.Products) == 0 {
		return "Basket is empty"
	}

	basketDescription := ""
	for _, product := range b.Products {
		basketDescription += product.String() + "\n"
	}

	return "Basket:\n" + basketDescription
}
