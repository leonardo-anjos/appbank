package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreditCard struct {
	id               string
	name             string
	number           string
	expiration_month int32
	expiration_year  int32
	cvv              int32
	balance          float64
	limit            float64
	created_at       time.Time
}

func new_credit_card() *CreditCard {
	creditCard := &CreditCard{}
	creditCard.id = uuid.NewV4().String()
	creditCard.created_at = time.Now()
	return creditCard
}
