package remoteRPC

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mitmedialab/medrec/DatabaseManager/common"
)

type PatientDocumentsArgs struct {
	PatientID int
	Time      string
	Signature string
}

type PatientDocumentsReply struct {
	//restructure as a JSON string
	Documents common.Documents
	Error     string
}

//returns a pointer to an sql database.
func (client *MedRecRemote) PatientDocuments(r *http.Request, args *PatientDocumentsArgs, reply *PatientDocumentsReply) error {
	patientAccount, err := AuthenticatePatient(args.Time, args.Signature)
	if err != nil {
		return err
	}

	tab := common.InstantiateLookupTable()
	defer tab.Close()
	db := instantiateDatabase()
	defer db.Close()

	uid, err := tab.Get([]byte(common.PrefixPatientUID(patientAccount)), nil)
	if err != nil {
		return err
	}

	fmt.Printf("fetching records for patient %s with uid %s \n", patientAccount, string(uid))

	rows, err := db.Query(fmt.Sprintf("SELECT PatientID, DocumentID, DocDateTime, PracticeID, RecvdDateTime FROM document_info WHERE PatientID = %s", uid))
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
		log.Println("have another row")
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
