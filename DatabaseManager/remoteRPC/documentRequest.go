package remoteRPC

import (
	"fmt"
	"log"
	"net/http"

	"../common"
	_ "github.com/go-sql-driver/mysql"
)

type PatientDocumentsArgs struct {
	PatientID int
}

type PatientDocumentsReply struct {
	//restructure as a JSON string
	Documents common.Documents
	Error     string
}

//returns a pointer to an sql database.
func (client *MedRecRemote) PatientDocuments(r *http.Request, args *PatientDocumentsArgs, reply *PatientDocumentsReply) error {

	db := instantiateDatabase()
	defer db.Close()

	newID := lookupPatient([]byte("address1"))
	s := string(newID[:])

	log.Println("got id from level db", s)

	fmt.Printf("fetching records for patient %d \n", args.PatientID)

	rows, err := db.Query(fmt.Sprintf("SELECT PatientID, DocumentID, DocDateTime, PracticeID, RecvdDateTime FROM document_info WHERE PatientID = %d", 1))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var (
		PatientID     string
		DocumentID    int
		DocDateTime   string
		PracticeID    string
		RecvdDateTime string
	)
	for rows.Next() {
		err = rows.Scan(&PatientID, &DocumentID, &DocDateTime, &PracticeID, &RecvdDateTime)
		if err != nil {
			log.Fatal(err)
		}
		result := *new(common.Document)

		if err != nil {
			log.Fatal(err)
			reply.Error = err.Error()
		}
		result.PatientID = PatientID
		result.DocumentID = DocumentID
		result.DocDateTime = DocDateTime
		result.PracticeID = PracticeID
		result.RecvdDateTime = RecvdDateTime

		reply.Documents = append(reply.Documents, result)
	}
	err = rows.Err()
	return err
}
