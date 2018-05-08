package remoteRPC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"../ethereum"
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

	faucetArgs := &FaucetArgs{
		Account: "0x000f8fdee72ac11b5c542428b35eed5769c409f0",
		Time:    fmt.Sprintf("%d", time.Now().Unix()),
	}
	faucetArgs.Signature, _ = Sign(faucetArgs.Time, faucetArgs.Account)

	faucetReply := &FaucetReply{}

	//create a connection over json rpc to the ethereum client
	rpcClient, _ := ethereum.GetEthereumRPCConn()

	// execute a ftransaction funding the user account with some ether
	var balance string
	err := rpcClient.Call(&balance, "eth_getBalance", faucetArgs.Account, "latest")
	if err != nil {
		log.Fatalf("Failed to get current eth balance: %v", err)

	}
	balance1, _ := strconv.ParseInt(balance[2:], 16, 64)

	client.PatientFaucet(nil, faucetArgs, faucetReply)
	if faucetReply.Error != "" {
		t.Errorf("The Faucet threw an error: %s", faucetReply.Error)
	}

	err = rpcClient.Call(&balance, "eth_getBalance", faucetArgs.Account, "latest")
	if err != nil {
		log.Fatalf("Failed to get current eth balance: %v", err)
	}
	balance2, _ := strconv.ParseInt(balance[2:], 16, 64)
	diff := balance2 - balance1
	if diff != 9223372036854775807 {
		t.Errorf("The faucet did not transfer 9223372036854775807 wei, actual: %d", diff)
	}
}
