package graphql

import (
	"context"
	"errors"
	"github.com/satori/go.uuid"
	"time"
)

type Auth struct {
	user  User
	token string
}

var TOKEN_EXPIRE = time.Duration(time.Duration.Minutes(30))

func (a *Auth) User() *User {
	return &a.user
}

func (a *Auth) Token() string {
	return a.token
}

func (_ *resolvers) Login(_ context.Context, s *struct{ Email, Password string }) (*Auth, error) {
	var u User
	if db.Where("email = ? AND password = ?", s.Email, encryptPass(s.Password)).First(&u).RecordNotFound() {
		return nil, errors.New("bad auth")
	}

	t, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	ts := t.String()

	a := Auth{
		user:  u,
		token: ts,
	}

	//TODO instead of just the email, I might want to put the id or the whole user
	redisClient.Set(ts, u.Email, TOKEN_EXPIRE)

	return &a, nil
}

func (_ *resolvers) Whoami(ctx context.Context) (*User, error) {
	u, err := getCtxUser(ctx)

	return u, err
}
