package database

import (
	"fmt"
	"go-api/pkg/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Postgres(conf *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.DatabaseUser,
		conf.DatabasePass,
		conf.DatabaseHost,
		conf.DatabasePort,
		conf.DatabaseName,
	)

	return sqlx.Open("postgres", dataSourceName)
}
