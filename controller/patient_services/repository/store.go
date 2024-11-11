package patientrepository

import "github.com/jmoiron/sqlx"

type patientStore struct {
	db *sqlx.DB
}

func NewPatientStore(db *sqlx.DB) *patientStore {
	return &patientStore{db: db}
}
