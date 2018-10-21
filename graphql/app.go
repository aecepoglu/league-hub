package graphql

import (
	"context"
	"errors"
	"github.com/aecepoglu/league-hub/config"
	"github.com/go-redis/redis"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
)

var redisClient *redis.Client
var db *gorm.DB
var conf config.ConfigType

func getCtxUser(ctx context.Context) (*User, error) {
	p := ctx.Value("user")

	if p != nil {
		u := p.(*User)
		if u != nil {
			return u, nil
		}
	}

	return nil, errors.New("unauth")
}

func newCtxWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, "user", user)
}

func Handler() (http.Handler, error) {
	b, err := ioutil.ReadFile("graphql/schema.graphql")
	if err != nil {
		return nil, err
	}

	schema := graphql.MustParseSchema(string(b), &resolvers{})
	gqlHandler := relay.Handler{Schema: schema}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		t := r.Header.Get("Auth-Token")
		if t == "" {
			gqlHandler.ServeHTTP(w, r.WithContext(newCtxWithUser(ctx, nil)))
			return
		}

		usernamer := redisClient.Get(t)

		err := usernamer.Err()
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("bad auth"))
			return
		}

		u := User{
			Email: usernamer.Val(),
		}

		gqlHandler.ServeHTTP(w, r.WithContext(newCtxWithUser(ctx, &u)))
	}), nil
}

func SetRedis(r *redis.Client) {
	redisClient = r
}

func SetDb(d *gorm.DB) {
	db = d
	migrate()
}

func SetConf(c config.ConfigType) {
	conf = c
}
