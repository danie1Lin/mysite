package service

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"mysite/pkg/cache"
	"mysite/pkg/errors"
	"mysite/pkg/models"
)

type UserService interface {
	Login(ctx context.Context, loginEmail, password string) (*models.User, error)
	SignUp(ctx context.Context, email, password string) error
	GetProfile(ctx context.Context) (*models.User, error)
}

type userService struct {
	logger       log.Logger
	store        models.UserStore
	cacheStore   cache.CacheStore
	sessionStore cache.SessionStore
}

func NewUserService(logger log.Logger, db *sql.DB, cacheStore cache.CacheStore, sessionStore cache.SessionStore) *userService {
	return &userService{
		logger:       logger,
		store:        models.NewPostgresStore(db, logger),
		cacheStore:   cacheStore,
		sessionStore: sessionStore,
	}
}

func (s *userService) Login(ctx context.Context, loginEmail, password string) (*models.User, error) {
	level.Info(s.logger).Log("action", "login", "email", loginEmail, "pswd", password, "jwt")
	user, err := s.store.FindByEmail(ctx, loginEmail)
	if err != nil {
		return nil, err
	}
	if !user.VerifyPassword(password) {
		return nil, errors.New("LOGIN_FAILED", "帳號或密碼錯誤")
	}
	if err := s.saveSessionData(ctx, user); err != nil {
		return nil, err
	}
	return user, err
}

func (s *userService) getSessionData(ctx context.Context) (map[string]string, error) {
	sessionId := ctx.Value(s.sessionStore.ContextKey()).(string)
	data, err := s.sessionStore.GetData(sessionId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *userService) saveSessionData(ctx context.Context, user *models.User) error {
	sessionId := ctx.Value(s.sessionStore.ContextKey()).(string)
	err := s.sessionStore.WriteKey(sessionId, "user_id", user.ID)
	return err
}

func (s *userService) SignUp(ctx context.Context, email, password string) error {
	level.Info(s.logger).Log("action", "signup", "email", email)
	user, err := s.store.CreateUser(ctx, email, password)
	if err != nil {
		return err
	}
	err = s.saveSessionData(ctx, user)
	return err
}

func (s *userService) GetProfile(ctx context.Context) (*models.User, error) {
	data, err := s.getSessionData(ctx)
	if err != nil {
		return nil, err
	}
	return s.store.FindByID(ctx, data["user_id"])
}
