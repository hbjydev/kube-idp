package user

import (
	"time"

	"github.com/hbjydev/kube-idp/server/internal/db"
)

type User struct {
	Id          string         `json:"id"`
	Login       string         `json:"login"`
	DisplayName *db.NullString `json:"display_name"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
