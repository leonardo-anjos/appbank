package repository

import (
	"database/sql"
	"github.com/leonardo-anjos/appbank/domain"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func new_transaction_repository_db(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (transaction *TransactionRepositoryDb) save_transaction(
	transaction domain.Transaction, credit_card domain.CreditCard) error {
		
		stmt, err := transaction.db.Prepare(`
			insert into transactions(id, credit_card_id, amount, status, description, store, created_at)
			values($1, $2, $3, $4, $5, $6, $7)
		`)

		if err != nil {
			return err
		}

		_, err = stmt.Exec(
			transaction.id, 
			transaction.credit_card_id, 
			transaction.amount, 
			transaction.status, 
			transaction.description, 
			transaction.store, 
			transaction.created_at
		)

		if err != nil {
			return err
		}

		if transaction.status == "approved" {
			
		}

		err = stmt.Close()

		if err != nil {
			return err
		}

		return nil
}

func (transaction *TransactionRepositoryDb) update_balance(credit_card domain.CreditCard) error {
	_, err := transaction.db.Exec("update credit_card set balance = $1 where id = $2", credit_card.balance, credit_card.id)
	
	if err != nil {
		return err
	}
	
	return nil
}

