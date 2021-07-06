package repository

import "database/sql"

type TransactionRepositoryDb struct {
	db *sql.DB
}

func new_transaction_repository_db(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}
