package remoteRPC

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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

type AddAccountArgs struct {
	Account  string
	UniqueID string
}

type AddAccountReply struct {
	Error string
}

func getUniqueID(account string) []byte {
	tab := instantiateLookupTable()
	defer tab.Close()

	data, err := tab.Get([]byte("uid"+account), nil)
	if err != nil {
		log.Println(err)
	}

	return data
}

//should add test to check that:
//unique ID is not a duplicate
//unique id matches an entry in the database
func (client *MedRecRemote) AddAccount(r *http.Request, args *AddAccountArgs, reply *AddAccountReply) error {
	log.Println("adding new patient account", args.UniqueID, "     ", args.Account)

	tab := instantiateLookupTable()
	defer tab.Close()

	err := tab.Put([]byte("uid-"+args.Account), []byte(args.UniqueID), nil)
	if err != nil {
		log.Println(err)
	}

	return err
}
