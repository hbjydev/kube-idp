package user

import (
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/hbjydev/kube-idp/server/internal/db"
)

type UserRepository struct{}

func CreateUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetUserById(id string) (*User, error) {
	user := User{}

	users := db.Qb.Select("*").From("users").Where(sq.Eq{"id": id}).Limit(1)

	rows := users.RunWith(db.Db).QueryRow()
	if err := rows.Scan(&user.Id, &user.Login, &user.DisplayName, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		log.Printf("error trying to query user by id: %v", err)
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByLogin(login string) (*User, error) {
	user := User{}

	users := db.Qb.Select("*").From("users").Where(sq.Eq{`login`: login}).Limit(1)

	rows := users.RunWith(db.Db).QueryRow()
	if err := rows.Scan(&user.Id, &user.Login, &user.DisplayName, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		log.Printf("error trying to query user by id: %v", err)
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) CreateUser(user User) (*User, error) {
	query := db.Qb.
		Insert("users").
		Columns("login", "email", "password").
		Values(user.Login, user.Email, user.Password).
		Suffix(`returning "id", "createdat", "updatedat"`).
		RunWith(db.Db)

	if err := query.QueryRow().Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
