package main

import (
	"context"
	"errors"
)

type resolvers struct{}

type User struct {
	username string
	email string
	phone string
}

type Auth struct {
	user User
	token string
}

func (_ *resolvers) Hello() string {
	return "Hello, World!"
}

func (a *Auth) User() *User {
	return &a.user
}

func (a *Auth) Token() string {
	return a.token
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Phone() string {
	return u.phone
}

func (_ *resolvers) Login(_ context.Context, s *struct{Username, Password string}) (*Auth, error) {
	if s.Username != "user" || s.Password != "pass" {
		return nil, errors.New("bad auth")
	}

	u := User {
		username: s.Username,
		email: "existing email",
		phone: "existing phone",
	}

	a := Auth {
		user: u,
		token: "generated auth token",
	}

	return &a, nil
}
