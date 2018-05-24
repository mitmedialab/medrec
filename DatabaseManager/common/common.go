package common

import (
	"encoding/hex"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/syndtr/goleveldb/leveldb"
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

var WalletPassword string

//recover takes a message and returns the address of the sender
func ECRecover(msg string, signature string) (string, error) {
	msgHex := "0x" + hex.EncodeToString([]byte(msg))
	rpcClient, _ := GetEthereumRPCConn()

	var result string
	err := rpcClient.Call(&result, "personal_ecRecover", msgHex, signature)
	return result, err
}

//Sign takes a hashed message and signs it
func Sign(msg string, account string) (string, error) {
	msgHex := "0x" + hex.EncodeToString([]byte(msg))

	rpcClient, _ := GetEthereumRPCConn()

	var result string
	err := rpcClient.Call(&result, "personal_sign", msgHex, account, "")
	return result, err
}

func GetEthereumRPCConn() (*rpc.Client, error) {
	//create a connection over json rpc to the ethereum client
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return rpcClient, err
}

//instantiates the lookup table
func InstantiateLookupTable() *leveldb.DB {
	var home string

	if runtime.GOOS == "windows" {
		home = os.Getenv("APPDATA") + "/MedRec"
	} else if runtime.GOOS == "darwin" {
		log.Println("darwmin")
		home = os.Getenv("HOME") + "/Library/Preferences"
	} else {
		home = os.Getenv("HOME") + "/.medrec"
	}

	tab, err := leveldb.OpenFile(home+"/lookupTable", nil)
	if err != nil {
		log.Fatal(err)
	}

	return tab
}
