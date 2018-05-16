package common

import (
	"os/exec"
	"strings"
)

type Document struct {
	DocumentID    int    `json:"DocumentID"`
	PatientID     string `json:"patientid"`
	PracticeID    string `json:"practiceid"`
	RecvdDateTime string `json:"recvddatetime"`
	DocDateTime   string `json:"docdatetime"`
}

type Documents []Document

type Patient struct {
	PatientID  int    `json:"patientid"`
	EthAddress string `json:"ethaddress"`
}

type Patients []Patient

type NoArgs struct {
	// nothin
}

// NodeExec calls node and abstracts away figuring out where the node binary is located
func NodeExec(args ...string) *exec.Cmd {
	cmd := []string{"node"}
	cmd = append(cmd, args...)
	cmd = append(cmd, "||")
	cmd = append(cmd, "node")
	cmd = append(cmd, args...)
	return exec.Command("/bin/bash", []string{"-c", strings.Join(cmd, " ")}...)
}
