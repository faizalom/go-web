package models

import (
	"database/sql"
	"log"

	"github.com/faizalom/go-web/config"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq" // PostgreSQL driver
)

func Conn() *sql.DB {
	db, err := sql.Open("mysql", config.DB.URL)
	// db, err := sql.Open("postgres", config.DB.URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Count(db *sql.DB, query string, args ...any) (int, error) {
	var count int
	err := db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	return count, err
}
