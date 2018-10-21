package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/aecepoglu/jsonpath"
	conf "github.com/aecepoglu/league-hub/config"
	league "github.com/aecepoglu/league-hub/graphql"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type MyJsonpath struct {
	t *testing.T
	v interface{}
}

func NewMyJsonpath(v interface{}, t *testing.T) MyJsonpath {
	return MyJsonpath{t, v}
}

func (jp MyJsonpath) MustGet(path string) interface{} {
	val, err := jsonpath.Get(jp.v, path)
	if err != nil {
		jp.t.Fail()
	}
	return val
}

type GraphqlQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func startServer(t *testing.T) MyGraphqlClient {
	//set cwd to project root
	if _, fn, _, ok := runtime.Caller(0); ok != true {
		t.Fatalf("couldn't get runtime.Caller")
	} else {
		os.Chdir(filepath.Join(filepath.Dir(fn), "../"))
	}

	h, err := league.Handler()
	if err != nil {
		t.Fatal(err)
	}

	league.SetConf(conf.ConfigType{
		AdminPass: "adminpass",
	})

	if db, err := gorm.Open("sqlite3", "test.db"); err != nil {
		t.Fatal(err)
	} else {
		db.DropTableIfExists("users")
		db.DropTableIfExists("sports")
		db.LogMode(false)
		league.SetDb(db)
	}

	if mr, err := miniredis.Run(); err != nil {
		t.Fatal(err)
	} else {
		r := redis.NewClient(&redis.Options{
			Addr: mr.Addr(),
		})
		err = r.Ping().Err()
		if err != nil {
			t.Fatal(err)
		}
		league.SetRedis(r)
	}

	s := httptest.NewServer(h)
	return MyGraphqlClient{
		server: s,
		client: &http.Client{},
	}
}

type MyGraphqlClient struct {
	server *httptest.Server
	client *http.Client
	token  string
	//t *testing.T
}

func (c *MyGraphqlClient) Sendv(t *testing.T, query string, variables map[string]interface{}) MyJsonpath {
	t.Logf("Query: %s", query)
	bs, err := json.Marshal(GraphqlQuery{
		Query:     query,
		Variables: variables,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Query JSON: %s", bs)

	req, err := http.NewRequest("POST", c.server.URL, bytes.NewReader(bs))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("content-type", "application/json")
	if c.token != "" {
		req.Header.Add("Auth-Token", c.token)
	}

	res, err := c.client.Do(req)
	if err != nil {
		t.Fatal(err)
	} else if res.StatusCode != 200 {
		t.Logf("Response status is expected to be 200 but it is: %d", res.StatusCode)
		t.Fail()
	}
	defer res.Body.Close()

	b := new(bytes.Buffer)
	b.ReadFrom(res.Body)
	body := b.String()

	t.Logf("Response: %s", body)

	var v interface{}
	json.Unmarshal(b.Bytes(), &v)
	return NewMyJsonpath(v, t)
}

func (c *MyGraphqlClient) Send(t *testing.T, query string) MyJsonpath {
	return c.Sendv(t, query, nil)
}

/** logins the user and saves its token **/
func (c *MyGraphqlClient) Login(t *testing.T) {
	jp := c.Send(t, `query {
 login(email: "admin@mail.com", password: "adminpass") {
  token
 }
}`)
	tok := jp.MustGet("$data.login.token")
	c.token = tok.(string)
}
