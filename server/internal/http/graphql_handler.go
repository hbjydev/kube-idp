package http

import (
	"encoding/json"
	"log"
	_http "net/http"

	"github.com/graphql-go/graphql"
	"github.com/hbjydev/kube-idp/server/internal/dto"
	schema "github.com/hbjydev/kube-idp/server/internal/graphql"
)

func GraphQLHandler(w _http.ResponseWriter, req *_http.Request) {
	addCors(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

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
}
