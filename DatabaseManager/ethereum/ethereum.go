//go:generate abigen --sol ../../SmartContracts/contracts/AgentRegistry.sol --pkg ethereum --out AgentRegistry.go

package ethereum

import (
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

func Init() {
	AddSigner()
}

func GetEthereumRPCConn() (*rpc.Client, error) {
	//create a connection over json rpc to the ethereum client
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return rpcClient, err
}
