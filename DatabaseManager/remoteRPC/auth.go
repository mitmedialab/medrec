package remoteRPC

//TODO require that the authentication messsage is within a more stringent time domain
//such as between 0 and 5 minutes, additionally should requre that time signatures
//are monotonically increasing

import (
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type AuthPatientArgs struct {
	UniqueID string
}

// AuthenticateProvider verifies that the provided message was signed by a provider
// message is current time in seconds formatted as a string
// signature is the signature of the current time
// returns the patient account
func AuthenticatePatient(msg string, signature string) string {
	//fail authentication if the msg is too old
	msgInt, _ := strconv.ParseInt(msg, 10, 64)
	elapsedTime := time.Now().Sub(time.Unix(msgInt, 0))
	if elapsedTime.Minutes() > 10 {
		return ""
	}

	tab := instantiateLookupTable()
	defer tab.Close()

	messageSigner, _ := ECRecover(msg, signature)

	_, err := tab.Get([]byte("uid-"+messageSigner), nil)

	if err == nil {
		return messageSigner
	}

	return ""
}

type AuthProviderArgs struct {
	MsgHex    string
	Signature string
}

// AuthenticateProvider verifies that the provided message was signed by a provider
// message is current time in seconds formatted as a string
// signature is the signature of the current time
// returns provider account
func AuthenticateProvider(msg string, signature string) string {
	msgInt, _ := strconv.ParseInt(msg, 10, 64)
	elapsedTime := time.Now().Sub(time.Unix(msgInt, 0))
	if elapsedTime.Minutes() > 10 {
		return ""
	}

	//create a connection over json rpc to the ethereum client
	rpcClient, err := GetEthereumRPCConn()

	// get the current list of signers
	var signers []string
	err = rpcClient.Call(&signers, "clique_getSigners", "latest")
	if err != nil {
		log.Fatalf("Failed to get current signers list: %v", err)
	}
	log.Printf("signers list: %v\n", signers)
	messageSigner, _ := ECRecover(msg, signature)
	log.Println("message sign" + msg + " " + messageSigner)
	for _, signer := range signers {
		if signer == messageSigner {
			return messageSigner
		}
	}
	return ""
}

func lookupPatient(address []byte) []byte {

	tab := instantiateLookupTable()
	defer tab.Close()

	log.Println("instantiated, in lookup, address is", address)
	data, err := tab.Get(address, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(data)
	return data
}

type AddAccountArgs struct {
	Account   string
	Time      string
	Signature string
	UniqueID  string
}

type AddAccountReply struct {
	Error string
}

func getUniqueID(account string) []byte {
	tab := instantiateLookupTable()
	defer tab.Close()

	data, err := tab.Get([]byte("uid"+account), nil)
	if err != nil {
		log.Println(err)
	}

	return data
}

//should add test to check that:
//unique ID is not a duplicate
//unique id matches an entry in the database
func (client *MedRecRemote) AddAccount(r *http.Request, args *AddAccountArgs, reply *AddAccountReply) error {
	patientAddress, _ := ECRecover(args.Time, args.Signature)

	tab := instantiateLookupTable()
	defer tab.Close()

	err := tab.Put([]byte("uid-"+patientAddress), []byte(args.UniqueID), nil)
	if err != nil {
		log.Println(err)
	}

	return err
}
