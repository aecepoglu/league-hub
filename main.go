package main

import (
	"log"
	"net/http"
	//"path/filepath"
	//"os"

	league "github.com/aecepoglu/league-hub/graphql"
	config "github.com/aecepoglu/league-hub/config"
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
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	log.Println("Running at :8080");
	err = http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}
