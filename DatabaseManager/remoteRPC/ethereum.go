package remoteRPC

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/mitmedialab/medrec/DatabaseManager/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// GetMedRecRemoteRPCConn returns a connection to a random provider
func GetMedRecRemoteRPCConn() (*rpc.Client, string, error) {
	//create a connection over json rpc to the ethereum client
	gethClient, _ := common.GetEthereumRPCConn()

	// get the current list of signers
	var signers []string
	err := gethClient.Call(&signers, "clique_getSigners", "latest")
	if err != nil {
		log.Fatalf("Failed to get signers: %v", err)
	}
	rand.Seed(time.Now().Unix())
	nextProvider := signers[rand.Intn(len(signers))]
	//get the host info of the provider who should fufil the faucet request
	//using a js helper script
	host, err := exec.Command("node", "./GolangJSHelpers/getProviderHost.js", nextProvider).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to get the next provider's hostname: %v", err)
	}

	//create a connection over json rpc to the ethereum client
	rpcClient, err := rpc.Dial("http://" + string(host) + ":6337/remoteRPC")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return rpcClient, nextProvider, err
}

type PatientFaucetArgs struct {
	Account   string // the provider account to provide the refund
	Time      string // the current time
	Signature string // signature of the current time
}

type PatientFaucetReply struct {
	Error string
	Txid  string
}

//PatientFaucet takes an ethereum account and gives it some ether
//Message and Signature should be from the patient
//Account should refer to the Provider's account that money should be sent from
func (client *MedRecRemote) PatientFaucet(r *http.Request, args *PatientFaucetArgs, reply *PatientFaucetReply) error {
	patientAddress, err := AuthenticatePatient(args.Time, args.Signature)
	if err != nil {
		return err
	}

	// sign the current time
	curTime := fmt.Sprintf("%d", time.Now().Unix())
	signature, err := common.Sign(curTime, args.Account)
	if err != nil {
		log.Fatalf("Failed to Sign: %v", err)
	}

	rpcClient, providerAddress, _ := GetMedRecRemoteRPCConn()
	nextArgs := &ProviderFaucetArgs{patientAddress, providerAddress, curTime, signature}
	err = rpcClient.Call(&reply, "MedRecRemote.ProviderFaucet", nextArgs)
	return err
}

type ProviderFaucetArgs struct {
	RecipientAccount string // the patient account to recieve the funds
	SendingAccount   string // the provider account providing the funds
	Time             string // the current time
	Signature        string // signature of the current time
}

type ProviderFaucetReply struct {
	Error string
	Txid  string
}

//ProviderFaucet takes an ethereum account and gives it some ether
// The Message and Signature should be from the requesting provider
// The Account should be of the patient to whom funds should be sent
func (client *MedRecRemote) ProviderFaucet(r *http.Request, args *ProviderFaucetArgs, reply *ProviderFaucetReply) error {
	log.Printf("provider faucet: %v", args)
	_, err := AuthenticateProvider(args.Time, args.Signature)
	if err != nil {
		log.Printf("there was auth error: %v", err)
		return err
	}

	//create a connection over json rpc to the ethereum client
	rpcClient, _ := common.GetEthereumRPCConn()

	//get the list of accounts open on the client
	var accounts []string
	err = rpcClient.Call(&accounts, "eth_accounts")
	if err != nil {
		log.Fatalf("Failed to get the ethereum accounts: %v", err)
	}

	// execute a ftransaction funding the user account with some ether
	var txid string
	value := "0x" + strconv.FormatInt(1000000000000000000, 16)
	txObject := map[string]string{"from": args.SendingAccount, "to": args.RecipientAccount, "value": value}
	err = rpcClient.Call(&txid, "eth_sendTransaction", txObject)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	//reply with the transaction id
	reply.Txid = txid

	return err
}
