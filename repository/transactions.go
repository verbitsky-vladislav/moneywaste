package repository

import (
	"database/sql"
	"moneywaste/repository/models"
)

type Transactions struct {
	db *sql.DB
}

func NewTransactions(db *sql.DB) *Transactions {
	return &Transactions{
		db: db,
	}
}

func (s *Transactions) Create(transaction models.Transaction) (models.Transaction, error) {
	var newTrans models.Transaction
	err := s.db.QueryRow(`INSERT INTO "Transaction" 
		(userid, type, amount, categoryid, date, description)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		transaction.UserId, transaction.Type, transaction.Amount, transaction.CategoryId, transaction.Date, transaction.Description).Scan(&newTrans)
	if err != nil {
		return newTrans, err
	}

	return newTrans, nil
}

func (s *Transactions) Update() {

}

func (s *Transactions) Delete() {

}

func (s *Transactions) GetOneById() {

}

func (s *Transactions) GetAll() {
	var transactions []models.Transaction

}
