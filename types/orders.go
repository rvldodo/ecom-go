package types

type OrderStore interface {
	GetListOrder() ([]*Order, error)
	GetOrderById(id int) (*Order, error)
	CreateOrder(*Order) error
}

type Order struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Total   int    `json:"total"`
	Status  string `json:"status"`
	Address string `json:"address"`
	Timestamp
}

type OrderPayload struct {
	UserID  int    `json:"user_id" validate:"required"`
	Total   int    `json:"total"   validate:"required"`
	Status  string `json:"status"  validate:"required"`
	Address string `json:"address" validate:"required"`
}
