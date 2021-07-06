package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Transaction struct {
	id             string
	amount         float64
	status         string
	description    string
	store          string
	credit_card_id string
	created_at     time.Time
}

type TransactionRepository interface {
	save_transaction(transaction Transaction, credit_card CreditCard) error
	get_credit_card(credit_card CreditCard) (CreditCard, error)
	new_credit_card(credit_card CreditCard) error
}

func new_transaction() *Transaction {
	transaction := &Transaction{}
	transaction.id = uuid.NewV4().String()
	transaction.created_at = time.Now()
	return transaction
}

func (transaction *Transaction) process_validade(credit_card *CreditCard) {
	if transaction.amount+credit_card.balance > credit_card.limit {
		transaction.status = "rejected"
	} else {
		transaction.status = "approved"
		credit_card.balance = credit_card.balance + transaction.amount
	}
}
