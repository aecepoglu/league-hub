package main

import (
	"log"
	"net/http"

	config "github.com/aecepoglu/league-hub/config"
	league "github.com/aecepoglu/league-hub/graphql"
	"github.com/aecepoglu/league-hub/reactFileServer"
)

func foo() int {
	return 2
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	league.SetConf(conf)

	db, err := connectDb("dev.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	league.SetDb(db)

	redis, err := connectRedis(conf.RedisUri)
	if err != nil {
		log.Fatal(err)
		return
	}
	league.SetRedis(redis)

	h, err := league.Handler()
	if err != nil {
		log.Fatal(err)
		return
	}

	http.Handle("/graphql", h)
	//log.Printf("Serving files in %s\n", assets)
	http.Handle("/", reactFileServer.New("./client/public"))

	log.Println("Running at :8080")
	//TODO handle https
	err = http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}
