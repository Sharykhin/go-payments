package service

import (
	"context"
	"time"
)

type (
	// TODO: maybe it's better to have some sort of PasswordService?
	// UserIdentifier interface provides methods around user identity
	// like password and also tracks last login
	UserIdentifier interface {
		CreatePassword(ctx context.Context, userID int64, password string) (string, error)
		FindUserPassword(ctx context.Context, userID int64) (string, error)
		ValidatePassword(ctx context.Context, password string, compare string) (bool, error)
		UpdateLastLogin(ctx context.Context, userID int64, lastLogin time.Time) error
	}

	//TODO:
	// I have concerns regarding UpdateLastLogin method

	// Tokener is a general interface that provides method for working
	// with identification through the token string. Providers can be different, like
	// JWT, OAuth etc.
	Tokener interface {
		Generate(ctx context.Context, claims map[string]interface{}, expiration time.Duration) (string, error)
		Validate(ctx context.Context, token string) (map[string]interface{}, error)
	}
)
