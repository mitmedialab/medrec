package remoteRPC

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

type FaucetArgs struct {
	Account   string // the provider account to provide the refund
	Time      string // the current time
	Signature string // signature of the current time
}

type FaucetReply struct {
	Error string
	Txid  string
}

type RecoverArgs struct {
	MsgHex    string
	Signature string
}

type RecoverReply struct {
	Account string
}

func GetEthereumRPCConn() (*rpc.Client, error) {
	//create a connection over json rpc to the ethereum client
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return rpcClient, err
}

func GetMedRecRemoteRPCConn(host string) (*rpc.Client, error) {
	//create a connection over json rpc to the ethereum client
	rpcClient, err := rpc.Dial("http://" + host + ":6337/remoteRPC")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return rpcClient, err
}

//recover takes a hashed message and returns the address of the sender
func (client *MedRecRemote) Recover(r *http.Request, args *RecoverArgs, reply *RecoverReply) error {

	log.Println("message is: " + args.MsgHex)
	log.Println("signature is: " + args.Signature)

	result, err := ECRecover(args.MsgHex, args.Signature)
	reply.Account = result
	return err
}

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

//PatientFaucet takes an ethereum account and gives it some ether
//Message and Signature should be from the patient
//Account should refer to the Provider's account that money should be sent from
func (client *MedRecRemote) PatientFaucet(r *http.Request, args *FaucetArgs, reply *FaucetReply) error {
	patientAddress := AuthenticatePatient(args.Time, args.Signature)
	if patientAddress == "0x" {
		//TODO test, I don't think the code ever goes to this case
		return errors.New("patient does not exist for this provider")
	}
	//create a connection over json rpc to the ethereum client
	gethClient, _ := GetEthereumRPCConn()

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
	host, err := exec.Command("node", "../../GolangJSHelpers/getProviderHost.js", nextProvider).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to get the next provider's hostname: %v", err)
	}

	// sign the current time
	curTime := fmt.Sprintf("%d", time.Now().Unix())
	signature, err := Sign(curTime, args.Account)
	if err != nil {
		log.Fatalf("Failed to Sign: %v", err)
	}

	nextArgs := &FaucetArgs{patientAddress, curTime, signature}
	rpcClient, _ := GetMedRecRemoteRPCConn(string(host))
	rpcClient.Call(&reply, "MedRecRemote.ProviderFaucet", nextArgs)
	return err
}

//ProviderFaucet takes an ethereum account and gives it some ether
// The Message and Signature should be from the requesting provider
// The Account should be of the patient to whom funds should be sent
func (client *MedRecRemote) ProviderFaucet(r *http.Request, args *FaucetArgs, reply *FaucetReply) error {
	providerAddress := AuthenticateProvider(args.Time, args.Signature)
	//TODO check that this actually gets called
	if providerAddress == "" {
		return errors.New("provider check failed")
	}

	//create a connection over json rpc to the ethereum client
	rpcClient, _ := GetEthereumRPCConn()

	//get the list of accounts open on the client
	var accounts []string
	err := rpcClient.Call(&accounts, "eth_accounts")
	if err != nil {
		log.Fatalf("Failed to get the ethereum accounts: %v", err)
	}

	// execute a ftransaction funding the user account with some ether
	var txid string
	value := "0x" + strconv.FormatInt(1000000000000000000, 16)
	txObject := map[string]string{"from": accounts[0], "to": args.Account, "value": value}
	err = rpcClient.Call(&txid, "eth_sendTransaction", txObject)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	//reply with the transaction id
	reply.Txid = txid

	return nil
}
