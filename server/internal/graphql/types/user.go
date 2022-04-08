package types

import (
	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: `User`,
	Fields: graphql.Fields{
		`id`: &graphql.Field{
			Type: graphql.String,
		},
		`login`: &graphql.Field{
			Type: graphql.String,
		},
		`displayName`: &graphql.Field{
			Type: NullableString,
		},
		`email`: &graphql.Field{
			Type: graphql.String,
		},
		`createdAt`: &graphql.Field{
			Type: graphql.DateTime,
		},
		`updatedAt`: &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
