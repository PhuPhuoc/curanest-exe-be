package techniquerepository

import "github.com/jmoiron/sqlx"

type techniqueStore struct {
	db *sqlx.DB
}

func NewTechniqueStore(db *sqlx.DB) *techniqueStore {
	return &techniqueStore{db: db}
}
