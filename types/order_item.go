package types

type OrderItemStore interface {
	GetListOrderItem() ([]*OrderItem, error)
	GetOrderItemById(id int) (*OrderItem, error)
	CreateOrderItem(*OrderItem) error
}

type OrderItem struct {
	ID        int `json:"id"`
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
	Timestamp
}

type OrderItemPayload struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
}
