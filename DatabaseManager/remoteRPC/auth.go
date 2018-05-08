package remoteRPC

//TODO require that the authentication messsage is within a more stringent time domain
//such as between 0 and 5 minutes, additionally should requre that time signatures
//are monotonically increasing

import (
	"log"
	"strconv"
	"time"

	"../ethereum"

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
	messageSigner += "0x"
	uid, err := tab.Get([]byte("uid"+messageSigner), nil)
	if err != nil {
		log.Println(err)
	}

	if len(uid) == 0 {
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
	rpcClient, err := ethereum.GetEthereumRPCConn()

	// get the current list of signers
	var signers []string
	err = rpcClient.Call(&signers, "clique_getSigners")
	if err != nil {
		log.Fatalf("Failed to get current signers list: %v", err)
	}

	messageSigner, _ := ECRecover(msg, signature)
	messageSigner += "0x"
	for _, signer := range signers {
		if signer == messageSigner {
			return messageSigner
		}
	}
	return ""
}
