package remoteRPC

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestRecover(t *testing.T) {
	signerAccount := "0xc5b2fe6f6bc85d71f4ae9a335896c9308ec8977c"
	recoverArgs := &RecoverArgs{
		Time:      "1526422718",
		Signature: "0xe27f440e8520bd9ea447505a3ff42f5220e29f285ac542806aff78ab66f8a95f03d3c5edb809c0b39e4233a3dc2ce71bf6ba61e0783e8abafce8d7fdd815bb711c",
	}

	type ecRecoverResult struct {
		data []byte
	}

	result, _ := ECRecover(recoverArgs.Time, recoverArgs.Signature)

	if result != signerAccount {
		t.Errorf("ECRecover failed")
	}
}

// TestFaucet requires an ethereum client to be running locally
//  with rpc enabled on port 8545,
//  ganache-cli is recommended
func TestFaucet(t *testing.T) {
	os.Chdir("../../")

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
	err = client.PatientFaucet(nil, faucetArgs, faucetReply)
	if faucetReply.Error != "" && faucetReply.Txid != "" {
		t.Errorf("The Faucet threw an error: %s", faucetReply.Error)
	}
	if err != nil {
		t.Errorf("Faucet failed with error: %v", err)
	}
}
