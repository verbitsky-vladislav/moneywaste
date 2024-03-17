package repository

import (
	"database/sql"
	"fmt"
	"log"
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

func (u *User) GetUserByNickname(nickname string) (*models.UserGet, error) {
	var user models.UserGet

	query := fmt.Sprintf(`SELECT * FROM "User" WHERE nickname = '%s'`, nickname)

	rows, err := u.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(rows)
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Nickname, &user.Password)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println(user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &user, nil
}

func (u *User) GetUserById(id string) (*models.UserGet, error) {
	var user models.UserGet

	query := fmt.Sprintf(`SELECT * FROM "User" WHERE id = '%s'`, id)

	rows, err := u.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(rows)
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Nickname, &user.Password)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println(user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &user, nil
}
