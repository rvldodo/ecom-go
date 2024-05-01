package types

type OrderStore interface {
	CreateOrder(*Order) (int, error)
	CreateOrderItem(*OrderItem) error
}

type Order struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	Total   float64 `json:"total"`
	Status  string  `json:"status"`
	Address string  `json:"address"`
	Timestamp
}

type CartCheckoutPayload struct {
	Items []CartCheckoutItem `json:"items" validate:"required"`
}
