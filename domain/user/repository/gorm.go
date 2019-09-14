package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/domain/user/repository/entity"
)

type (
	// GORMRepository struct responsible for working with
	// storage layer by using GORM ORM
	GORMRepository struct {
		conn *gorm.DB
	}
)

// Create creates a new user in a database and returns just created record
func (r GORMRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	err := r.conn.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert a new user row: %v", err)
	}

	return &user, nil
}

// NewGORMRepository is a constructor function
// that returns a new instance of GORMRepository
func NewGORMRepository() *GORMRepository {
	return &GORMRepository{
		conn: database.G,
	}
}
