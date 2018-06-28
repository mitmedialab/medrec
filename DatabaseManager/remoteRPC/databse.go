package remoteRPC

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func instantiateDatabase() *sql.DB {

	db, err := sql.Open("mysql",
		"root:medrecpassword@tcp(127.0.0.1:3306)/medrec-v1")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
