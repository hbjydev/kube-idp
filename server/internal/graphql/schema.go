package graphql

import (
	"github.com/graphql-go/graphql"
    "github.com/hbjydev/kube-idp/server/internal/dto"
    "github.com/hbjydev/kube-idp/server/internal/graphql/types"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: `RootQuery`,
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        types.User,
			Description: `Get single user`,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                idQuery, isOK := params.Args["id"].(string)

				if isOK {
					// Search for el with id
					return dto.User{Id: idQuery, Username: `Something`}, nil
				}

				return dto.User{}, nil
			},
		},
	},
})

var IdpSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
