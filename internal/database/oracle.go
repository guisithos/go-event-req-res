package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/godror/godror"
)

func Connect() (*sql.DB, error) {
	dsn := os.Getenv("ORACLE_DSN")
	if dsn == "" {
		dsn = "oracle:1521/ORCLPDB1"
	}
	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to Oracle database")
	return db, nil
}
