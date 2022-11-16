package model

import (
	"fmt"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type UserName string

func NewUserName(name string) (UserName, error) {
	if name == "" || utf8.RuneCountInString(name) > 64 {
		return "", fmt.Errorf("user name must be 1-64 characters")
	}

	return UserName(name), nil
}

type UserPassword string

func NewUserPassword(rawPw string) (UserPassword, error) {
	if rawPw == "" || utf8.RuneCountInString(rawPw) > 128 {
		return "", fmt.Errorf("user password must be 1-128 characters")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(rawPw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return UserPassword(pw), nil
}

type UserRole string

func NewUserRole(role string) (UserRole, error) {
	if role == "" || utf8.RuneCountInString(role) > 64 {
		return "", fmt.Errorf("user role must be 1-64 characters")
	}

	return UserRole(role), nil
}

type User struct {
	ID        UUID
	Name      UserName
	Password  UserPassword
	Role      UserRole
	CreatedAt time.Time
}

func NewUser(name UserName, pw UserPassword, role UserRole, created_at time.Time) *User {
	return &User{
		ID:        NewUUID(),
		Name:      name,
		Password:  pw,
		Role:      role,
		CreatedAt: created_at,
	}
}
