package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/hbjydev/kube-idp/server/internal/graphql/types"
	"github.com/hbjydev/kube-idp/server/internal/user"
)

var CreateUser = &graphql.Field{
	Type:        types.User,
	Description: `Create a new User`,
	Args: graphql.FieldConfigArgument{
		`login`: &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		`password`: &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		`email`: &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		`displayName`: &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		login, _ := p.Args[`login`].(string)
		password, _ := p.Args[`password`].(string)
		email, _ := p.Args[`email`].(string)

		newUser := user.User{
			Login:    login,
			Password: password,
			Email:    email,
		}

		userRepo := user.CreateUserRepository()
		user, err := userRepo.CreateUser(newUser)
		if err != nil {
			return nil, err
		}

		return user, nil
	},
}
