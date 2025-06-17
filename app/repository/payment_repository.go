package repository

import (
	"database/sql"
	"fmt"
	"souzalambdago/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryInterface interface {
	GetTransactionById(transactionId string) (model.Transaction, error)
	CreateTransaction(model model.Transaction) error
	GetTransactionsByFromUser(fromUserId int) ([]model.Transaction, error)
	UpdateTransactionStatus(status string, transactionId string) error
}

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (pr *TransactionRepository) GetTransactionById(transactionId string) (model.Transaction, error) {
	var transaction model.Transaction
	err := pr.db.Get(&transaction, `SELECT * FROM transactions WHERE id = ?`, transactionId)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Transaction{}, nil // Return empty transaction if not found
		}
		fmt.Printf("Error getting transaction by id: %s", err.Error())
		return model.Transaction{}, err
	}
	return transaction, nil
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

func (pr *TransactionRepository) UpdateTransactionStatus(status string, transactionId string) error {
	_, err := pr.db.Exec(`UPDATE transactions
						  SET status = ?
						  WHERE id = ?`, status, transactionId)
	if err != nil {
		fmt.Println("Error updating transaction status:", err.Error())
		return err
	}
	return nil
}
