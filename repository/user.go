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

func (u *User) CreateUser(user models.UserCreate) {
	res, err := u.db.Exec(
		fmt.Sprintf(`"INSERT INTO "User" * VALUES %s, %s"`, user.Nickname, user.Password),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func (u *User) UpdateUser() {

}

func (u *User) GetUser() {

}
