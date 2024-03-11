package repository

import (
	"database/sql"
	"fmt"
	"moneywaste/repository/models"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) CreateUser(user models.UserCreate) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO "User" (nickname, password) VALUES ('%s', '%s') RETURNING id`, user.Nickname, user.Password)

	row := u.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *User) UpdateUser() {

}

func (u *User) GetUser() {

}
