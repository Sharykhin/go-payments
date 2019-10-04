package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	GORMDB "github.com/Sharykhin/go-payments/core/database/gorm"
	"github.com/Sharykhin/go-payments/core/errors"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/domain/user/repository/entity"
)

type (
	// GORMRepository struct responsible for working with
	// storage layer by using GORM ORM
	GORMRepository struct {
		conn *gorm.DB
	}
)

// NewGORMRepository is a constructor function
// that returns a new instance of GORMRepository
func NewGORMRepository() *GORMRepository {
	return &GORMRepository{
		//conn: database.G,
		conn: GORMDB.G,
	}
}

// Create creates a new user in a database and returns just created record
func (r GORMRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	err := r.conn.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert a new user row: %v", err)
	}

	return &user, nil
}

// FindByEmail looks for user by its email and takes the first one.
// Actually email should be unique
func (r GORMRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := entity.User{Email: email}
	err := r.conn.Where(&user).First(&user).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.ResourceNotFound
		}
		logger.Error("gorm could not execute select statement to find user by email: %v", err)
		return nil, fmt.Errorf("repository failed to find user by email %s: %v", email, err)
	}

	return &user, nil
}
