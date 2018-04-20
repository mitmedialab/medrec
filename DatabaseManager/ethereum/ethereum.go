package ethereum

import (
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

func Init() {
	AddSigner()
}

func GetRPCConn() *rpc.Client {
	//create a connection over json rpc to the ethereum client
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return rpcClient
}
