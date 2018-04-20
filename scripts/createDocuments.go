package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//returns a pointer to an sql database.
func main() {
	//doesn't open db, but prepares abstraction
	db, err := sql.Open("mysql",
		"root:medrecpassword@tcp(127.0.0.1:3306)/medrec-v1")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO document_info (DocumentID, PatientID, PracticeID, RecvdDateTime)
                          VALUES (445534, 4, 1234567, '2017-01-01 00:00:00')`)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("connection error")
	}

	defer db.Close()
}
