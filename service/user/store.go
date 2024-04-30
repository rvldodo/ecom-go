package user

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

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var users types.User
	if err := db.DB.Where("email = ?", email).Find(&users).Error; err != nil {
		log.Fatalf("error querying database: %v", err)
		return nil, err
	}
	return &users, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	var user types.User
	err := db.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		log.Fatalf("user not found: %v", err)
		return nil, err
	}
	return &user, nil
}

func (s *Store) CreateUser(user *types.User) error {
	res := db.DB.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	log.Println("User created")
	return nil
}
