package db

import (
	"database/sql"
	"fmt"
	"lxdAssessmentServer/ent"
	"lxdAssessmentServer/pkg"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/lib/pq"
)

func NewClient(config pkg.PostgresConfig) (*ent.Client, *sql.DB, error) {
	connector, err := pq.NewConnector(config.ConnectionString())

	if err != nil {
		return nil, nil, fmt.Errorf("failed to open a connection to postgres: %w", err)
	}

	db := sql.OpenDB(connector)

	drv := entsql.OpenDB("postgres", db)

	return ent.NewClient(ent.Driver(drv)), db, nil

}
