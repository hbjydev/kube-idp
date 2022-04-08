package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/hbjydev/kube-idp/server/internal/db"
	"github.com/hbjydev/kube-idp/server/internal/dto"
	schema "github.com/hbjydev/kube-idp/server/internal/graphql"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Db.Close()

	http.HandleFunc(`/graphql`, func(w http.ResponseWriter, req *http.Request) {
		var p dto.PostData
		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			w.WriteHeader(400)
			return
		}

		result := graphql.Do(graphql.Params{
			Context:        req.Context(),
			Schema:         schema.IdpSchema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})

		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Printf(`could not write result to response: %v`, err)
		}
	})

	log.Println(`server listening on *:8080`)

	log.Fatal(http.ListenAndServe(`:8080`, nil))
}
