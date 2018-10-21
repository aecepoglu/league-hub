package graphql

import (
	"context"
	is "github.com/stretchr/testify/assert"
	"testing"
)

func TestSportFields(t *testing.T) {
	owner := User{
		Email: "sport owner",
	}
	setupTestDb(t)

	is := is.New(t)
	is.Nil(db.Create(&owner).Error)

	s := Sport{
		Name:          "sport name",
		TeamSize:      7,
		CreatedBy:     owner,
		CreatedByUser: owner.ID(),
	}

	is.Equal("sport name", s.NAME())
	is.Equal(s.NAME(), s.ID())
	is.Equal(int32(7), s.TEAMSIZE())
	is.Equal(owner, *s.CREATEDBY())

	is.Nil(db.Delete(&owner).Error)
	is.Nil(db.Delete(&s).Error)

	cleanupTestDb(t)
}

func TestCreateSport(t *testing.T) {
	setupTestDb(t)
	r := &resolvers{}
	ctx := context.Background()

	in := sportInput{
		Name:     "new sport name",
		TeamSize: 7,
	}
	user := User{
		Email: "ahmet emre",
		Phone: "5554441122",
	}
	oldSport := Sport{
		Name:      "old sport name",
		TeamSize:  2,
		CreatedBy: user,
	}

	is.Nil(t, db.Create(&user).Error)
	is.Nil(t, db.Create(&oldSport).Error)

	t.Run("Creates sport that belongs to current user", func(t *testing.T) {
		s, err := r.CreateSport(newCtxWithUser(ctx, &user), struct{ In sportInput }{In: in})
		is.Nil(t, err)
		is.Equal(t, *s, Sport{
			Name:          "new sport name",
			TeamSize:      7,
			CreatedBy:     user,
			CreatedByUser: user.ID(),
		})

		is.Nil(t, db.Delete(s).Error)
	})

	t.Run("Gives error if unauth", func(t *testing.T) {
		s, err := r.CreateSport(newCtxWithUser(ctx, nil), struct{ In sportInput }{In: in})
		is.NotNil(t, err)
		is.Nil(t, s)
		is.Equal(t, "unauth", err.Error())
	})

	t.Run("Gives error if sport name already exists", func(t *testing.T) {
		_, err := r.CreateSport(newCtxWithUser(ctx, &user), struct{ In sportInput }{In: sportInput{Name: "old sport name", TeamSize: 3}})
		is.NotNil(t, err)
	})

	cleanupTestDb(t)
}
