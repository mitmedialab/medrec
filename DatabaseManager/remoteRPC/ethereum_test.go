package remoteRPC

import (
	"log"
	"strconv"
	"testing"
	"encoding/gob"
	"encoding/hex"
	"bytes"

	"github.com/ethereum/go-ethereum/rpc"
)

func TestRecover(t *testing.T){
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	recoverArgs := &RecoverArgs{
		MsgHash: "0x7f6c0e5c497ded52462ec18daeb1c94cefa11cd6949ebdb7074b2a32cac13fba",
      	Signature: "0xd4f70d60a000c5426e10f6e58caf02b7a5c21248dbc0dba47f179d97891a88d12d07d54c69edc24f34ae88722d275f7d198f5de0ffd8897e06979c9b3c36bb6200",
    }

	type ecRecoverResult struct {
		data []byte
	}

	var result map[string]interface{}
	err = rpcClient.Call(&result, "eth_ecRecover", recoverArgs.MsgHash, recoverArgs.Signature)
	if err != nil {
		log.Fatalf("Failed to return ethereum address: %v", err)
	}
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	enc.Encode(result["data"])

	decodedResult := make([]byte, 20)
	buf.Read(decodedResult)

	log.Println(result["data"])
	log.Println(buf.Len())
	log.Println("0x" + hex.EncodeToString(decodedResult))

}

// TestFaucet requires an ethereum client to be running locally
//  with rpc enabled on port 8545,
//  ganache-cli is recommended
func TestFaucet(t *testing.T) {
	client := new(MedRecRemote)

	faucetArgs := &FaucetArgs{
		Account: "0x000f8fdee72ac11b5c542428b35eed5769c409f0",
	}

	faucetReply := &FaucetReply{}

	//create a connection over json rpc to the ethereum client
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// execute a ftransaction funding the user account with some ether
	var balance string
	err = rpcClient.Call(&balance, "eth_getBalance", faucetArgs.Account)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	balance1, _ := strconv.ParseInt(balance[2:], 16, 64)

	client.Faucet(nil, faucetArgs, faucetReply)
	if faucetReply.Error != "" {
		t.Errorf("The Faucet threw an error: %s", faucetReply.Error)
	}

	err = rpcClient.Call(&balance, "eth_getBalance", faucetArgs.Account)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	balance2, _ := strconv.ParseInt(balance[2:], 16, 64)
	diff := balance2 - balance1
	if diff != 9223372036854775807 {
		t.Errorf("The faucet did not transfer 9223372036854775807 wei, actual: %d", diff)
	}
}
