package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	pswd      string
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.pswd), []byte(password))
	if err != nil {
		return false
	}
	return true
}
