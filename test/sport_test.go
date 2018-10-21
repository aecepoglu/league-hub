package test

import (
	is "github.com/stretchr/testify/assert"
	"testing"
)

func TestListSports(t *testing.T) {
	s := startServer(t)

	jp := s.Send(t, `query {
 listSports() {
  name
  teamSize
  createdBy {
   email
  }
 }
}`)

	is := is.New(t)

	is.Equal("Squash", jp.MustGet("$data.listSports[0].name"))
	is.Equal(float64(1), jp.MustGet("$data.listSports[0].teamSize"))
	is.Equal("admin@mail.com", jp.MustGet("$data.listSports[0].createdBy.email"))

	is.Equal("Tennis", jp.MustGet("$data.listSports[1].name"))
	is.Equal(float64(1), jp.MustGet("$data.listSports[1].teamSize"))
	is.Equal("admin@mail.com", jp.MustGet("$data.listSports[1].createdBy.email"))
}

func TestCreateSport(t *testing.T) {
	s := startServer(t)

	s.Login(t)
	jp := s.Send(t, `mutation {
 createSport(in: {name: "new sport", teamSize: 5}) {
  name
  teamSize
  createdBy {
   email
  }
 }
}`)

	is.Equal(t, "new sport", jp.MustGet("$data.createSport.name"))
	is.Equal(t, float64(5), jp.MustGet("$data.createSport.teamSize"))
	is.Equal(t, "admin@mail.com", jp.MustGet("$data.createSport.createdBy.email"))
}
