package remoteRPC

import (
	"../params"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type AuthArgs struct {
	PatientID int
}

type AuthReply struct {
	//restructure as a JSON string
	Patients params.Patients
	Error    string
}

func (client *MedRecRemote) Authenticate(db *sql.DB, args *AuthArgs, reply *AuthReply) error {

	rows, err := db.Query(fmt.Sprintf("SELECT PatientID FROM patient_info WHERE PatientID = %d", 1))
	if err != nil {
		log.Fatal(err)
	}

	var (
		PatientID int
	)

	for rows.Next() {
		err = rows.Scan(&PatientID)
		if err != nil {
			log.Fatal(err)
		}
		result := *new(params.Patient)

		if err != nil {
			log.Fatal(err)
			reply.Error = err.Error()
		}
		result.PatientID = PatientID

		reply.Patients = append(reply.Patients, result)
	}

	return err
}
