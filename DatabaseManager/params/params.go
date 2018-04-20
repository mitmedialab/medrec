package params

type Document struct {
	DocumentID    int    `json:"DocumentID"`
	PatientID     string `json:"patientid"`
	PracticeID    string `json:"practiceid"`
	RecvdDateTime string `json:"recvddatetime"`
	DocDateTime   string `json:"docdatetime"`
}

type Documents []Document

type Patient struct{
	PatientID 	 int  `json:"patientid"`
	EthAddress 	 string `json:"ethaddress"`
}

type Patients []Patient

type NoArgs struct {
	// nothin
}
