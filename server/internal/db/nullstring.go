package db

import (
	"database/sql"
	"encoding/json"

	"github.com/graphql-go/graphql/language/ast"
)

// NullString to be used in place of sql.NullString
type NullString struct {
	sql.NullString
}

// MarshalJSON from the json.Marshaler interface
func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON from the json.Unmarshaler interface
func (v *NullString) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.String = *x
		v.Valid = true
	} else {
		v.Valid = false
	}
	return nil
}

// NewNullString create a new null string. Empty string evaluates to an
// "invalid" NullString
func NewNullString(value string) *NullString {
	var null NullString
	if value != "" {
		null.String = value
		null.Valid = true
		return &null
	}
	null.Valid = false
	return &null
}

// SerializeNullString serializes `NullString` to a string
func SerializeNullString(value interface{}) interface{} {
	switch value := value.(type) {
	case NullString:
		return value.String
	case *NullString:
		v := *value
		return v.String
	default:
		return nil
	}
}

// ParseNullString parses GraphQL variables from `string` to `CustomID`
func ParseNullString(value interface{}) interface{} {
	switch value := value.(type) {
	case string:
		return NewNullString(value)
	case *string:
		return NewNullString(*value)
	default:
		return nil
	}
}

// ParseLiteralNullString parses GraphQL AST value to `NullString`.
func ParseLiteralNullString(valueAST ast.Value) interface{} {
	switch valueAST := valueAST.(type) {
	case *ast.StringValue:
		return NewNullString(valueAST.Value)
	default:
		return nil
	}
}
