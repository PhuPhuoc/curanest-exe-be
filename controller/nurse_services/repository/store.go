package nurserepository

import "github.com/jmoiron/sqlx"

type nurseStore struct {
	db *sqlx.DB
}

func NewNurseStore(db *sqlx.DB) *nurseStore {
	return &nurseStore{db: db}
}
