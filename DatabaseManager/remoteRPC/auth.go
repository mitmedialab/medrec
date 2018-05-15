package remoteRPC

//TODO require that the authentication messsage is within a more stringent time domain
//such as between 0 and 5 minutes, additionally should requre that time signatures
//are monotonically increasing

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
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

	tab := instantiateLookupTable()
	defer tab.Close()

	messageSigner, _ := ECRecover(msg, signature)

	_, err := tab.Get([]byte("uid-"+messageSigner), nil)

	if err == nil {
		return messageSigner, nil
	}

	return "", errors.New("account could not be found")
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
			return messageSigner, nil
		}
	}
	return "", errors.New("account is not a provider")
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
	patientAddress, err := ECRecover(args.Time, args.Signature)
	if err != nil {
		return err
	}

	tab := instantiateLookupTable()
	defer tab.Close()

	err = tab.Put([]byte("uid-"+patientAddress), []byte(args.UniqueID), nil)
	if err != nil {
		log.Println(err)
	}

	return err
}

type AgentContractArgs struct {
	Account   string //the acctount that should be set as the owner of the new AgentContract
	Time      string //unix time encoded into a hex string
	Signature string //signature of the time
}

type AgentContractReply struct {
	ContractAddress string
}

func (client *MedRecRemote) PatientAgentContract(r *http.Request, args *AgentContractArgs, reply *AgentContractReply) error {
	_, err := AuthenticatePatient(args.Time, args.Signature)
	if err != nil {
		return err
	}
	//TODO check if the user already has an agent contract associated with them}

	// sign the current time
	curTime := fmt.Sprintf("%d", time.Now().Unix())
	signature, err := Sign(curTime, args.Account)
	if err != nil {
		log.Fatalf("Failed to Sign: %v", err)
	}

	newAccount, err := exec.Command("node", "../../GolangJSHelpers/generateNewAccount.js", WalletPassword).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to update the Agent Registry: %v", err)
	}

	nextArgs := &AgentContractArgs{string(newAccount), curTime, signature}
	rpcClient, _ := GetMedRecRemoteRPCConn()
	rpcClient.Call(&reply, "MedRecRemote.ProviderAgentContract", nextArgs)

	_, err = exec.Command("node", "../../GolangJSHelpers/addAgentToRegistry.js", reply.ContractAddress).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to update the Agent Registry: %v", err)
	}

	return err
}

func (client *MedRecRemote) ProviderAgentContract(r *http.Request, args *AgentContractArgs, reply *AgentContractReply) error {
	_, err := AuthenticateProvider(args.Time, args.Signature)
	if err != nil {
		return err
	}

	//create the agent contract and set it's owner using a helper script
	contractAddr, err := exec.Command("node", "../../GolangJSHelpers/makeNewAgent.js", args.Account).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to generate a new AgentContract %s for: %v", contractAddr, err)
	}
	reply.ContractAddress = string(contractAddr)

	return nil
}
