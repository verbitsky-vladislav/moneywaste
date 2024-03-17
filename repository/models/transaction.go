package models

import "time"

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type Transaction struct {
	Id          string          `json:"id,omitempty"`
	UserId      string          `json:"userId"`
	Type        TransactionType `json:"type"`
	Amount      float64         `json:"amount"`
	CategoryId  string          `json:"categoryId,omitempty"`
	Date        time.Time       `json:"date"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"createdAt,omitempty"`
	UpdatedAt   time.Time       `json:"updatedAt,omitempty"`
}

type TransactionCategory struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
