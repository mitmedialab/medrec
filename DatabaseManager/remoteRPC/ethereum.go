package remoteRPC

import (
	"log"
	"net/http"
	"strconv"
	"github.com/ethereum/go-ethereum/rpc"
)

type FaucetArgs struct {
	Account string
}

type FaucetReply struct {
	Error string
	Txid  string
}

type RecoverArgs struct {
	//MsgHash map[string]interface{}
	MsgHex string
	Signature string
}

type RecoverReply struct {
	Msg []byte
	Sender []byte
}

//recover takes a hashed message and returns the address of the sender
func (client *MedRecRemote) Recover(r *http.Request, args *RecoverArgs, reply *RecoverReply) error {

	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}


	log.Println("message is: " + args.MsgHex)
	log.Println("signature is: " + args.Signature)

	type ecRecoverResult struct {
		data []byte
	}

	//var result map[string]interface{}
	var result string
	err = rpcClient.Call(&result, "eth_ecRecover", args.MsgHex, args.Signature)
	if err != nil {
		log.Fatalf("Failed to return ethereum address: %v", err)
	}

	log.Println("return address is " + result)
	return err
}

//Faucet takes an ethereum account and gives it some ether
func (client *MedRecRemote) Faucet(r *http.Request, args *FaucetArgs, reply *FaucetReply) error {
	//create a connection over json rpc to the ethereum client
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	//get the list of accounts open on the client
	var accounts []string
	err = rpcClient.Call(&accounts, "eth_accounts")
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
