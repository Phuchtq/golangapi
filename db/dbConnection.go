package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//var cnn *sql.DB

func ConnectDB() (*sql.DB, error) {
	if cnn, err := sql.Open("postgres", os.Getenv("CNN_STR")); err != nil {
		log.Print("Fail to get access to database")
		return nil, err
	} else {
		return cnn, nil
	}
}
