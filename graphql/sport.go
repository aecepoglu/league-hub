package graphql

import (
	"context"
	"log"
)

type Sport struct {
	Name          string `gorm:"primary_key"`
	TeamSize      int
	CreatedBy     User `gorm:"foreignkey:CreatedByUser;association_autoupdate:false;association_autocreate:false"`
	CreatedByUser string
}

type sportInput struct {
	Name     string
	TeamSize int32
}

func (s *Sport) NAME() string {
	return s.Name
}

func (s *Sport) TEAMSIZE() int32 {
	return int32(s.TeamSize)
}

func (s *Sport) CREATEDBY() *User {
	var u User
	db.Model(s).Related(&u, "CreatedBy")
	return &u
}

func (s *Sport) ID() string {
	return s.Name
}

func (_ *resolvers) CreateSport(ctx context.Context, args struct{ In sportInput }) (*Sport, error) {
	log.Println("inside createSport")
	u, err := getCtxUser(ctx)
	if err != nil {
		return nil, err
	}

	s := Sport{
		Name:      args.In.Name,
		TeamSize:  int(args.In.TeamSize),
		CreatedBy: *u,
	}

	err = db.Create(&s).Error

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (_ *resolvers) ListSports(_ context.Context) (*[]*Sport, error) {
	var sports []*Sport
	err := db.Find(&sports).Error

	if err != nil {
		return nil, err
	}

	return &sports, nil
}
