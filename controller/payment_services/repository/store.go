package paymentrepository

import "github.com/jmoiron/sqlx"

type paymentStore struct {
	db *sqlx.DB
}

func NewPaymentStore(db *sqlx.DB) *paymentStore {
	return &paymentStore{db: db}
}
