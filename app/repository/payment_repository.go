package repository

import (
	"fmt"
	"souzalambdago/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(model model.Transaction) error
	GetTransactionsByFromUser(fromUserId int) ([]model.Transaction, error)
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
	_, err := pr.db.NamedExec(`INSERT INTO transactions (id, from_user_id, to_user_id, amount, type, status, description, created_at)
	VALUES (:id, :from_user_id, :to_user_id, :amount, :type, :status, :description, :created_at)`, model)
	if err != nil {
		fmt.Println("Error inserting transaction:", err.Error())
		return err
	}
	return nil
}

func (pr *TransactionRepository) GetTransactionsByFromUser(fromUserId int) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := pr.db.Select(&transactions, `SELECT * FROM transactions WHERE from_user_id = ?`, fromUserId)
	if err != nil {
		fmt.Printf("Error getting transactions from user %d\n: %s", fromUserId, err.Error())
		return nil, err
	}
	return transactions, nil
}
