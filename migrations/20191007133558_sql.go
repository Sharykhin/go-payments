package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20191007133558, Down20191007133558)
}

func Up20191007133558(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func Down20191007133558(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
