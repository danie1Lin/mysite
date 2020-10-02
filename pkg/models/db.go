package models

import (
	"context"
	"database/sql"
	"mysite/pkg/errors"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type postgresStore struct {
	db     *sql.DB
	logger log.Logger
}

func (p *postgresStore) FindByEmail(ctx context.Context, email string) (*User, error) {
	u := &User{}
	stmt, err := p.db.PrepareContext(
		ctx,
		"SELECT id, email, pswd, created_at, updated_at FROM users WHERE email = $1",
	)
	if err != nil {
		return u, err
	}

	err = stmt.QueryRowContext(ctx, email).Scan(&u.ID, &u.Email, &u.pswd, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		err = RecordNotFoundError
	}

	return u, err
}

func (p *postgresStore) FindByID(ctx context.Context, id string) (*User, error) {
	u := &User{}
	stmt, err := p.db.PrepareContext(
		ctx,
		"SELECT id, email, pswd, created_at, updated_at FROM users WHERE id = $1",
	)
	if err != nil {
		return u, err
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&u.ID, &u.Email, &u.pswd, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		err = RecordNotFoundError
	}
	return u, err
}

func (p *postgresStore) CreateUser(ctx context.Context, email, pswd string) (*User, error) {
	pswdByte := []byte(pswd)
	hash, err := bcrypt.GenerateFromPassword(pswdByte, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	sql := "INSERT INTO users (email, pswd, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING (id, email, created_at, updated_at)"

	stmt, err := p.db.Prepare(sql)

	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	user := User{}
	err = stmt.QueryRowContext(ctx, email, hash, now, now).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		switch v := err.(type) {
		case *pq.Error:
			if v.Code == "23505" && v.Constraint == "users_email_key" {
				return nil, errors.New("EMAIL_USED", "email 已經被使用")
			}
			p.logger.Log("db error", v)
		}
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewPostgresStore(db *sql.DB, logger log.Logger) *postgresStore {
	return &postgresStore{
		db:     db,
		logger: logger,
	}
}
