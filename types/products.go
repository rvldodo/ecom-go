package types

type ProductStore interface {
	GetListProducts() ([]*Product, error)
	GetProductById(id int) (*Product, error)
	CreateProduct(*Product) error
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Quantity    int    `json:"quantity"`
	Timestamp
}

type ProductPayload struct {
	Name        string `json:"name"        validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"       validate:"required"`
	Quantity    int    `json:"quantity"    validate:"required"`
}
