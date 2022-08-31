package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DBConnection *sql.DB

func init() {
	// Open a connection to connect mysql database
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_curd")
	// Handler connect error
	if err != nil {
		log.Printf("Open mysql database happended an error: %v\n", err)
	}
	DBConnection = db
}
