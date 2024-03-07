package types

// Special is a struct that represents a special price
type Special struct {
	Quantity int `yaml:"quantity"`
	Price    int `yaml:"price"`
}
