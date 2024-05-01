package product

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
	return &Store{
		db: db,
	}
}

func (s *Store) GetListProducts() (*[]types.Product, error) {
	var products []types.Product
	if err := db.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (s *Store) GetProductById(id []int) ([]types.Product, error) {
	var product []types.Product
	if err := db.DB.Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Store) CreateProduct(product *types.Product) error {
	res := db.DB.Create(product)
	if res.Error != nil {
		return res.Error
	}
	log.Print("Product created")
	return nil
}

func (s *Store) UpdateProduct(product *types.Product) error {
	res := db.DB.Save(&product)
	if res.Error != nil {
		return res.Error
	}
	log.Println("Product updated")
	return nil
}
