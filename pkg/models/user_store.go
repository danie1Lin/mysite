package models

import (
	"context"
)

type UserStore interface {
	// User service
	FindByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, email, pswd string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
}
