package appointmentrepository

import "github.com/jmoiron/sqlx"

type appoinmentStore struct {
	db *sqlx.DB
}

func NewAppoinmentStore(db *sqlx.DB) *appoinmentStore {
	return &appoinmentStore{db: db}
}
