package u_sushi

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	return db
}

func SetDB(conn *sqlx.DB) {
	db = conn
}

func ConnectToDB() error {
	wrapErr := func(err error) error {
		return fmt.Errorf("connecting to db: %w", err)
	}

	dbUrl := os.Getenv("DB_CONNECTION_URL")
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		return wrapErr(err)
	}

	SetDB(db)
	return nil
}
