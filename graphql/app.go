package graphql

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/aecepoglu/league-hub/config"
)

var redisClient *redis.Client
var db *gorm.DB
var conf config.ConfigType

func getAuthUser(ctx context.Context) (*User, error) {
	p := ctx.Value("user")
	if p == nil {
		return nil, errors.New("unauth")
	}

	return p.(*User), nil
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
			gqlHandler.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "user", nil)))
			return
		}

		usernamer := redisClient.Get(t)

		err := usernamer.Err()
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("bad auth"))
			return
		}

		u := User {
			Email: usernamer.Val(),
		}

		gqlHandler.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "user", &u)))
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
