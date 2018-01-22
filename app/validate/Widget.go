package validate

type Widget struct {
	Name string     `validate:"nonzero"`
	Color string     `validate:"nonzero"`
	Inventory int     `validate:"nonzero"`
	Price string     `validate:"nonzero"`
	Melts bool     `validate:"nonzero"`
}