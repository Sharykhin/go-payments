package repository

import (
	"fmt"

	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/user/entity"
	"github.com/jinzhu/gorm"
)

type (
	GORMRepository struct {
		conn *gorm.DB
	}
)

func (r GORMRepository) Create(user entity.User) (*entity.User, error) {
	err := r.conn.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert a new user row: %v", err)
	}

	return &user, nil
}

func NewGORMRepository() *GORMRepository {
	return &GORMRepository{
		conn: database.G,
	}
}
