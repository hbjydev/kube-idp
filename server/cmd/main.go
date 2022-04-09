package main

import (
	"log"
	_http "net/http"

	"github.com/hbjydev/kube-idp/server/internal/db"
	"github.com/hbjydev/kube-idp/server/internal/http"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Db.Close()

	_http.HandleFunc(`/graphql`, http.GraphQLHandler)

	log.Println(`server listening on *:8080`)

	log.Fatal(_http.ListenAndServe(`:8080`, nil))
}
