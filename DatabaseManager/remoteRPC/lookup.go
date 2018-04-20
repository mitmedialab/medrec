package remoteRPC

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type AccountArgs struct {
	Account string
	AgentID string
}

type AccountReply struct {
	Error string
}

 func lookupPatient(address []byte) []byte {

 	tab := instantiateLookupTable()
	defer tab.Close()

	log.Println("instantiated, in lookup, address is", address)
	data, err := tab.Get(address, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(data)
	return data
 }

//should add test to check that:
//unique ID is not a duplicate
//unique id matches an entry in the database
func (client *MedRecRemote) AddAccount(r *http.Request, args *AccountArgs, reply *AccountReply) error{
	log.Println("adding new patient account", args.AgentID, "     ", args.Account)

 	tab := instantiateLookupTable()
	defer tab.Close()

	err := tab.Put([]byte(args.Account), []byte(args.AgentID), nil)
	if err != nil {
		log.Println(err)
	}

	data, err := tab.Get([]byte(args.Account), nil)
	if err != nil {
		log.Println(err)
	}

	dataString := string(data[:])
	log.Println("account created for patient: ", dataString)
	return err
}
