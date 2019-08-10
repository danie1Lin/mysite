package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kit/kit/log"
)

type postgresStore struct {
	db     *sql.DB
	logger log.Logger
}

func (p *postgresStore) FindByEmailAndPswd(ctx context.Context, email string, pswd string) (*User, error) {
	u := &User{}
	stmt, err := p.db.PrepareContext(
		ctx,
		"SELECT id, email, created_at, updated_at FROM users WHERE email = $1 AND pswd = $2",
	)
	if err != nil {
		return u, err
	}
	err = stmt.QueryRowContext(ctx, email, pswd).Scan(&u.ID, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		err = RecordNotFoundError
	}
	return u, err
}

func (p *postgresStore) FindByEmail(ctx context.Context, email string) (*User, error) {
	u := &User{}
	err := p.db.QueryRowContext(
		ctx,
		"select id, email, created_at, updated_at From users Where users.email = $1",
		email,
	).Scan(u.ID, u.Email, u.CreatedAt, u.UpdatedAt)
	if err == sql.ErrNoRows {
		err = RecordNotFoundError
	}
	return u, err
}

func (p *postgresStore) CreateUser(ctx context.Context, email, pswd string) error {
	sql := "INSERT INTO users (email, pswd, created_at, updated_at) VALUES ($1, $2, $3, $4)"

	stmt, err := p.db.Prepare(sql)

	if err != nil {
		return err
	}
	now := time.Now().UTC()

	_, err = stmt.ExecContext(ctx, email, pswd, now, now)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewPostgresStore(db *sql.DB) *postgresStore {
	return &postgresStore{
		db: db,
	}
}
