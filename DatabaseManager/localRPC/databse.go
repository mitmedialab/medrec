package localRPC

import (
	"log"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/syndtr/goleveldb/leveldb"
)

//instantiates the lookup table
func instantiateLookupTable() *leveldb.DB {
	var home string

	if runtime.GOOS == "windows" {
		home = os.Getenv("APPDATA") + "/MedRec"
	} else if runtime.GOOS == "darwin" {
		log.Println("darwmin")
		home = os.Getenv("HOME") + "/Library/Preferences"
	} else {
		home = os.Getenv("HOME") + "/.medrec"
	}

	tab, err := leveldb.OpenFile(home+"/localLookupTable", nil)
	if err != nil {
		log.Fatal(err)
	}

	return tab
}
