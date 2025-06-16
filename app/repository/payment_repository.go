package repository

import (
	"fmt"
	"souzalambdago/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(model model.Transaction) error
}

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (pr *TransactionRepository) CreateTransaction(model model.Transaction) error {
	_, err := pr.db.NamedExec(`INSERT INTO transactions (id, from_user_id, to_user_id, amount, currency, status, created_at, updated_at)
	VALUES (:id, :from_user_id, :to_user_id, :amount, :currency, :status, :created_at, :updated_at)`, model)
	if err != nil {
		fmt.Println("Error inserting transaction:", err.Error())
		return err
	}
	return nil
}
