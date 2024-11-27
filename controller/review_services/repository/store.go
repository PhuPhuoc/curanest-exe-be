package reviewrepository

import "github.com/jmoiron/sqlx"

type reviewStore struct {
	db *sqlx.DB
}

func NewReviewStore(db *sqlx.DB) *reviewStore {
	return &reviewStore{db: db}
}
