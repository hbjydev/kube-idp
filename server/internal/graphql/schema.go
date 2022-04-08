package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/hbjydev/kube-idp/server/internal/graphql/resolvers"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: `RootQuery`,
	Fields: graphql.Fields{
		`user`: resolvers.GetUser,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: `RootMutation`,
	Fields: graphql.Fields{
		`createUser`: resolvers.CreateUser,
	},
})

var IdpSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
