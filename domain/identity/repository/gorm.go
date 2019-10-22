package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	GORMDB "github.com/Sharykhin/go-payments/core/database/gorm"
	types "github.com/Sharykhin/go-payments/core/type"
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

// Update is a general update methods that can update specified number of field abstracting
// any knowledge of data source and its schema
func (r GORMRepository) Update(ctx context.Context, userID int64, fields UpdateFields) error {
	user := entity.User{
		ID: userID,
	}
	r.conn.First(&user)

	if fields.LastLogin.Valid {
		user.LastLogin = types.NullTime{Valid: true, Time: fields.LastLogin.Time}
	}
	err := r.conn.Save(&user).Error

	return err
}

// NewGORMRepository is a constructor function
// that returns a new instance of GORMRepository
func NewGORMRepository() *GORMRepository {
	return &GORMRepository{
		conn: GORMDB.G,
	}
}
