package order

import (
	"log"

	"gorm.io/gorm"

	"github.com/dodo/ecom/db"
	"github.com/dodo/ecom/types"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order *types.Order) (int, error) {
	res := db.DB.Create(&order)
	if res.Error != nil {
		return 0, res.Error
	}
	log.Println("Order created")
	return order.ID, nil
}

func (s *Store) CreateOrderItem(orderItem *types.OrderItem) error {
	res := db.DB.Create(&orderItem)
	if res.Error != nil {
		return res.Error
	}
	log.Println("Order item created")
	return nil
}
