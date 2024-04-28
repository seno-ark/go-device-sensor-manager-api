package postgres

import (
	"go-api/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repositories.IRepository {
	return &repository{
		db: db,
	}
}
