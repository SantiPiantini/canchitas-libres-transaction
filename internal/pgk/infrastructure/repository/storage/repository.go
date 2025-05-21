package storage

import (
	"canchitas-libres-transaction/internal/configuration"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	Configuration *configuration.Configuration
	*sqlx.DB
}

func NewPostgresStorage(config *configuration.Configuration, db *sqlx.DB) *Postgres {
	return &Postgres{
		Configuration: config,
		DB:            db,
	}
}
