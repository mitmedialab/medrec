package remoteRPC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestRecover(t *testing.T) {
	recoverArgs := &RecoverArgs{
		MsgHex:    "0xdeadbeef",
		Signature: "0x5c707e85427d94de23e499b0742dd42f25629b8b33d8dfddee68b50fec59c8bd147354fbd6e66b4c135aaf922491ef4d87e427609e89c70e67c2860c60d7f45a1b",
	}

	type ecRecoverResult struct {
		data []byte
	}

	result, _ := ECRecover(recoverArgs.MsgHex, recoverArgs.Signature)
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	enc.Encode(result)

	decodedResult := make([]byte, 20)
	buf.Read(decodedResult)

	// log.Println(result)
	// log.Println(buf.Len())
	// log.Println("0x" + hex.EncodeToString(decodedResult))
	// TODO figure out how to get the encoding into a result

}

// TestFaucet requires an ethereum client to be running locally
//  with rpc enabled on port 8545,
//  ganache-cli is recommended
func TestFaucet(t *testing.T) {
	client := new(MedRecRemote)
	//create a connection over json rpc to the ethereum client
	rpcClient, _ := GetEthereumRPCConn()

	//get a local ethereum address which can be used for testing
	var accounts []string
	err := rpcClient.Call(&accounts, "eth_accounts")
	if err != nil {
		log.Fatalf("Failed to get an ethereum account: %v", err)
	}

	//make sure the patient exists in the database
	addArgs := &AddAccountArgs{
		Time:     fmt.Sprintf("%d", time.Now().Unix()),
		UniqueID: accounts[0],
	}
	addArgs.Signature, _ = Sign(addArgs.Time, accounts[0])
	addReply := &AddAccountReply{}
	client.AddAccount(nil, addArgs, addReply)

	//create the arguments to the call to the faucet
	faucetArgs := &FaucetArgs{
		Account: accounts[1],
		Time:    fmt.Sprintf("%d", time.Now().Unix()),
	}
	faucetArgs.Signature, _ = Sign(faucetArgs.Time, accounts[0])

	faucetReply := &FaucetReply{}

	//request the faucent send some ether
	client.PatientFaucet(nil, faucetArgs, faucetReply)
	if faucetReply.Error != "" && faucetReply.Txid != "" {
		t.Errorf("The Faucet threw an error: %s", faucetReply.Error)
	}
}
