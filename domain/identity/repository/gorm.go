package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	GORMDB "github.com/Sharykhin/go-payments/core/database/gorm"
	"github.com/Sharykhin/go-payments/domain/identity/repository/entity"
)

type (
	// GORMRepository struct responsible for working with
	// storage layer by using GORM ORM
	GORMRepository struct {
		conn *gorm.DB
	}
)

// Create creates a new user in a database and returns just created record
func (r GORMRepository) CreatePassword(cxt context.Context, userID int64, password string) (*entity.UserPassword, error) {
	up := entity.UserPassword{
		UserID:   userID,
		Password: password,
	}
	err := r.conn.Create(&up).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert a user password row: %v", err)
	}

	return &up, nil
}

func (r GORMRepository) FindPasswordByUserID(cxt context.Context, userID int64) ([]entity.UserPassword, error) {
	var up []entity.UserPassword
	err := r.conn.Where(entity.UserPassword{UserID: userID}).Order("created_at desc").Find(&up).Error
	if err != nil {
		return nil, fmt.Errorf("could execute find password by user id: %v", err)
	}

	return up, err
}

// NewGORMRepository is a constructor function
// that returns a new instance of GORMRepository
func NewGORMRepository() *GORMRepository {
	return &GORMRepository{
		conn: GORMDB.G,
	}
}
