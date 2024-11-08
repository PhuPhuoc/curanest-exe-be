package customerrepository

import "github.com/jmoiron/sqlx"

type customerStore struct {
	db *sqlx.DB
}

func NewCustomerStore(db *sqlx.DB) *customerStore {
	return &customerStore{db: db}
}
