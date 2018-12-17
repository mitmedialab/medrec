package localRPC

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/mitmedialab/medrec/DatabaseManager/common"
)

type AddAccountArgs struct {
	UniqueID string
	Account  string
	Username string
	Password string
}

type AddAccountReply struct {
}

//should add test to check that:
//unique ID is not a duplicate
//unique id matches an entry in the database
func (client *MedRecLocal) AddAccount(r *http.Request, args *AddAccountArgs, reply *AddAccountReply) error {

	tab := common.InstantiateLookupTable()
	defer tab.Close()

	err := tab.Put([]byte(common.PrefixPatientUID(args.Account)), []byte(args.UniqueID), nil)
	if err != nil {
		return err
	}

	newAccount, err := exec.Command("node", "./GolangJSHelpers/generateNewAccount.js", common.GetKeystorePath(args.Username), args.Password).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to generate a new account for this patient: %v", err)
	}

	err = tab.Put([]byte(strings.ToLower("patient-provider-account"+args.Account)), []byte(newAccount), nil)
	if err != nil {
		return err
	}

	return nil
}
