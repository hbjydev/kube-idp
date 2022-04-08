package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/hbjydev/kube-idp/server/internal/graphql/resolvers"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: `RootQuery`,
	Fields: graphql.Fields{
		"user": resolvers.GetUser,
	},
})

var IdpSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
