package config

import (
	"database/sql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/yourdb?parseTime=true")
	if err != nil {
		return nil
	}
	return db
}
