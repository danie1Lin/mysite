package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"mysite/pkg/cache"
	"mysite/pkg/user/models"
)

type UserService interface {
	Login(ctx context.Context, loginEmail, password string) (id, email string, createdAt, updatedAt time.Time, err error)
	SignUp(ctx context.Context, email, password string) error
}

type userService struct {
	logger     log.Logger
	store      models.UserStore
	cacheStore cache.CacheStore
}

func NewUserService(logger log.Logger, db *sql.DB, cacheStore cache.CacheStore) *userService {
	return &userService{
		logger:     logger,
		store:      models.NewPostgresStore(db),
		cacheStore: cacheStore,
	}
}

func (s *userService) Login(ctx context.Context, loginEmail, password string) (id, email string, createdAt, updatedAt time.Time, err error) {
	level.Info(s.logger).Log("action", "login", "email", loginEmail, "pswd", password, "jwt", ctx.Value(jwt.JWTClaimsContextKey))
	user, err := s.store.FindByEmailAndPswd(ctx, loginEmail, password)
	return user.ID, user.Email, user.CreatedAt, user.UpdatedAt, err
}

func (s *userService) SignUp(ctx context.Context, email, password string) error {
	level.Info(s.logger).Log("action", "signup", "email", email)
	return s.store.CreateUser(ctx, email, password)
}
