package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/hbjydev/kube-idp/server/internal/graphql/types"
	"github.com/hbjydev/kube-idp/server/internal/user"
)

var GetUser = &graphql.Field{
	Type:        types.User,
	Description: `Get single user`,

	Args: graphql.FieldConfigArgument{
		`login`: &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		loginQuery, isOK := params.Args[`login`].(string)

		if isOK {
			userRepo := user.CreateUserRepository()
			return userRepo.GetUserByLogin(loginQuery)
		}

		return user.User{}, nil
	},
}
