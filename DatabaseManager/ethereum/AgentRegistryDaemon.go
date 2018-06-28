package ethereum

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// AddSigner watches the AgentRegistry for the AddSigner event and upon detecting one votes
// in the signer via the clique protocol
func AddSigner() {
	//sha3 (keccack-256) hash of the "AddSigner(address)" event
	const addSignerHash = "0x637c77a2d598a51d085d4a2413332c45a235a25ee855bf3dfcdc2c8fcf02860f"

	rpcClient, _ := GetEthereumRPCConn()

	var events []map[string]interface{}
	var lastBlockNumber int64

	lastBlockNumber = 1
	hexBlockNumber := "0x" + fmt.Sprintf("%x", lastBlockNumber)
	txObject := map[string]string{"fromBlock": hexBlockNumber, "toBlock": "latest", "topic": addSignerHash}
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			err := rpcClient.Call(&events, "eth_getLogs", txObject)
			if err != nil {
				log.Printf("Failed to get the filter logs: %v", err)
				continue
			}

			for _, event := range events {
				addSignerTopic := event["topics"].([]interface{})
				hexBlockNumber := event["blockNumber"].(string)
				blockNumber, _ := strconv.ParseInt(hexBlockNumber, 0, 64)
				newSigner := "0x" + addSignerTopic[1].(string)[26:]
				log.Printf("New Signer added %s at block %d\n", newSigner, blockNumber)

				err = rpcClient.Call(&events, "clique_propose", newSigner, true)
				if err != nil {
					log.Printf("Non fatal err when proposing new signer to the clique, likely due to using ganache-cli")
				}

				if blockNumber > lastBlockNumber {
					lastBlockNumber = blockNumber
					hexBlockNumber := "0x" + fmt.Sprintf("%x", lastBlockNumber+1)
					txObject["fromBlock"] = hexBlockNumber
				}
			}
		}
	}()
}

// RemoveSigner watches the AgentRegistry for the RemoveSigner event and upon detecting one votes
// to kick a signer via the clique protocol
func RemoveSigner() {
	//sha3 (keccack-256) hash of the "RemoveSigner(address)" event
	const addSignerHash = "0x1803740ef72fc16e647c10fe2d31cf61a1578081960c2e3fb7f5aa957e82f550"

	rpcClient, _ := GetEthereumRPCConn()

	var events []map[string]interface{}
	var lastBlockNumber int64

	lastBlockNumber = 1
	hexBlockNumber := "0x" + fmt.Sprintf("%x", lastBlockNumber)
	txObject := map[string]string{"fromBlock": hexBlockNumber, "toBlock": "latest", "topic": addSignerHash}
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			err := rpcClient.Call(&events, "eth_getLogs", txObject)
			if err != nil {
				log.Printf("Failed to get the filter logs: %v", err)
				continue
			}

			for _, event := range events {
				removeSignerTopic := event["topics"].([]interface{})
				hexBlockNumber := event["blockNumber"].(string)
				blockNumber, _ := strconv.ParseInt(hexBlockNumber, 0, 64)
				kickedSigner := "0x" + removeSignerTopic[1].(string)[26:]
				log.Printf("Signer removed %s at block %d\n", kickedSigner, blockNumber)

				err = rpcClient.Call(&events, "clique_propose", kickedSigner, false)
				if err != nil {
					log.Printf("Non fatal err when proposing new signer to the clique, likely due to using ganache-cli")
				}

				if blockNumber > lastBlockNumber {
					lastBlockNumber = blockNumber
					hexBlockNumber := "0x" + fmt.Sprintf("%x", lastBlockNumber+1)
					txObject["fromBlock"] = hexBlockNumber
				}
			}
		}
	}()
}
