package types

import (
	"github.com/graphql-go/graphql"
	"github.com/hbjydev/kube-idp/server/internal/db"
)

var NullableString = graphql.NewScalar(graphql.ScalarConfig{
	Name:         "NullableString",
	Description:  "The `NullableString` type repesents a nullable SQL string.",
	Serialize:    db.SerializeNullString,
	ParseValue:   db.ParseNullString,
	ParseLiteral: db.ParseLiteralNullString,
})
