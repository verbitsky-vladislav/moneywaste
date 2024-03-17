package repository

import (
	"database/sql"
	"errors"
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

func (u *User) CreateUser(user models.User) (string, error) {
	var id string

	// Проверяем, существует ли уже пользователь с данным email
	err := u.db.QueryRow(`SELECT id FROM "User" WHERE email = $1`, user.Email).Scan(&id)
	if err == nil {
		// Пользователь найден, возвращаем существующий ID
		return id, nil
	} else if !errors.Is(sql.ErrNoRows, err) {
		// Произошла ошибка, отличная от отсутствия записи
		return "", err
	}

	// Пользователь не найден, создаем нового
	err = u.db.QueryRow(`INSERT INTO "User" (fio, email, passwordhash) VALUES ($1, $2, $3) RETURNING id`, user.Fio, user.Email, user.Password).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *User) UpdateUser() {

}

func (u *User) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf(`SELECT * FROM "User" WHERE email = '%s'`, email)

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
		err := rows.Scan(&user.Id, &user.Fio, &user.Email, &user.Password)
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

func (u *User) GetUserById(id string) (*models.User, error) {
	var user models.User

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
		err := rows.Scan(&user.Id, &user.Fio, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
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
