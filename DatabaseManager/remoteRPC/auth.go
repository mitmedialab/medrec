package remoteRPC

//TODO require that the authentication messsage is within a more stringent time domain
//such as between 0 and 5 minutes, additionally should requre that time signatures
//are monotonically increasing

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"../common"
)

// AuthenticateProvider verifies that the provided message was signed by a provider
// message is current time in seconds formatted as a string
// signature is the signature of the current time
// returns the patient account
func AuthenticatePatient(msg string, signature string) (string, error) {
	//fail authentication if the msg is too old
	msgInt, _ := strconv.ParseInt(msg, 10, 64)
	elapsedTime := time.Now().Sub(time.Unix(msgInt, 0))
	if elapsedTime.Minutes() > 10 {
		return "", errors.New("signature is too old")
	}

	tab := common.InstantiateLookupTable()
	defer tab.Close()

	messageSigner, _ := common.ECRecover(msg, signature)

	ret, _ := tab.Has([]byte("patient-uid-"+messageSigner), nil)

	if ret {
		return messageSigner, nil
	}

	return "", errors.New("account " + messageSigner + " could not be found")
}

// AuthenticateProvider verifies that the provided message was signed by a provider
// message is current time in seconds formatted as a string
// signature is the signature of the current time
// returns provider account
func AuthenticateProvider(msg string, signature string) (string, error) {
	msgInt, _ := strconv.ParseInt(msg, 10, 64)
	elapsedTime := time.Now().Sub(time.Unix(msgInt, 0))
	if elapsedTime.Minutes() > 10 {
		return "", errors.New("signature is too old")
	}

	// get the current list of signers
	var signers []string
	result, err := exec.Command("node", "./GolangJSHelpers/getSigners.js").CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to get current signers list: %v", err)
	}

	json.Unmarshal(result, &signers)

	messageSigner, _ := common.ECRecover(msg, signature)
	log.Print(signers)
	for _, signer := range signers {
		if signer == messageSigner {
			return messageSigner, nil
		}
	}
	return "", errors.New("account is not a provider")
}

type GetProviderAccountArgs struct {
	Time      string //unix time encoded into a hex string
	Signature string //signature of the time
}

type GetProviderAccountReply struct {
	Account string
}

func (client *MedRecRemote) GetProviderAccount(r *http.Request, args *GetProviderAccountArgs, reply *GetProviderAccountReply) error {
	patientAccount, err := AuthenticatePatient(args.Time, args.Signature)
	if err != nil {
		return err
	}

	tab := common.InstantiateLookupTable()
	defer tab.Close()
	account, err := tab.Get([]byte(strings.ToLower("patient-provider-account"+patientAccount)), nil)

	reply.Account = string(account)

	return err
}

type ChangeAccountArgs struct {
	Account   string
	Time      string
	Signature string
}

type ChangeAccountReply struct {
}

//ChangeAccount transfers the mapping from patient account to unique identifier to a different account
func (client *MedRecRemote) ChangeAccount(r *http.Request, args *ChangeAccountArgs, reply *ChangeAccountReply) error {
	patientAccount, err := AuthenticatePatient(args.Time, args.Signature)
	if err != nil {
		return err
	}

	tab := common.InstantiateLookupTable()
	defer tab.Close()

	uniqueID, _ := tab.Get([]byte(strings.ToLower("patient-uid-"+patientAccount)), nil)
	err = tab.Put([]byte(strings.ToLower("patient-uid-"+args.Account)), []byte(uniqueID), nil)
	if err != nil {
		return err
	}
	tab.Delete([]byte(strings.ToLower("patient-uid-"+patientAccount)), nil)

	newAccount, _ := tab.Get([]byte(strings.ToLower("patient-provider-"+patientAccount)), nil)
	err = tab.Put([]byte(strings.ToLower("patient-provider-account"+args.Account)), []byte(newAccount), nil)
	if err != nil {
		return err
	}
	tab.Delete([]byte(strings.ToLower("patient-provider-"+patientAccount)), nil)

	return nil
}
