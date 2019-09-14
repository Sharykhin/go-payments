package repository

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/domain/user/entity"
	"github.com/jinzhu/gorm"
)

type (
	GORMRepository struct {
		conn *gorm.DB
	}
)

func (r GORMRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
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
