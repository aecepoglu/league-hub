package test

import (
	is "github.com/stretchr/testify/assert"
	"testing"
)

func TestGoodLogin(t *testing.T) {
	s := startServer(t)

	jp := s.Send(t, `query {
 login(email: "admin@mail.com", password: "adminpass") {
  user {email}
  token
 }
}`)

	is.Equal(t, "admin@mail.com", jp.MustGet("$data.login.user.email"))
	is.NotNil(t, jp.MustGet("$data.login.token"))
}

func TestBadLogin(t *testing.T) {
	s := startServer(t)

	jp := s.Send(t, `query {
 login(email: "admin@mail.com", password: "badpass") {
  user {email}
  token
 }
}`)

	is.Equal(t, "bad auth", jp.MustGet("$errors[0].message"))
}

func TestWhoamiIsSecure(t *testing.T) {
	s := startServer(t)

	jp := s.Send(t, "query { whoami { email } }")

	is.Equal(t, "unauth", jp.MustGet("$errors[0].message"))
}

func TestWhoamiShowsUser(t *testing.T) {
	s := startServer(t)

	s.Login(t)
	jp := s.Send(t, "query { whoami { email } }")

	is.Equal(t, "admin@mail.com", jp.MustGet("$data.whoami.email"))
}
