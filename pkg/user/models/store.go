package models

import (
	"context"
)

type UserStore interface {
	FindByEmailAndPswd(ctx context.Context, email string, pswd string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, email, pswd string) error
}
