package ethereum

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func AddSigner() {
	//sha3 hash of the "AddSigner(address)" event
	const addSignerHash = "0x637c77a2d598a51d085d4a2413332c45a235a25ee855bf3dfcdc2c8fcf02860f"

	rpcClient, _ := GetEthereumRPCConn()

	var events []map[string]interface{}
	var lastBlockNumber int64

	lastBlockNumber = 1
	hexBlockNumber := "0x" + fmt.Sprintf("%x", lastBlockNumber)
	txObject := map[string]string{"fromBlock": hexBlockNumber, "toBlock": "latest", "topic": addSignerHash}
	ticker := time.NewTicker(time.Millisecond * 500)
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
