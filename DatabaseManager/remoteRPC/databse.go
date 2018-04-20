package remoteRPC

import (
	"database/sql"
	"log"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/syndtr/goleveldb/leveldb"
)

func instantiateDatabase() *sql.DB {

	db, err := sql.Open("mysql",
		"root:medrecpassword@tcp(127.0.0.1:3306)/medrec-v1")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

//instantiates the lookup table
func instantiateLookupTable() *leveldb.DB {
	var home string

	if runtime.GOOS == "windows" {
		home = os.Getenv("APPDATA") + "/MedRec"
	} else if runtime.GOOS != "darwin" {
		home = os.Getenv("HOME") + "/Library/Preferences"
	} else {
		home = os.Getenv("HOME") + "/.medrec"
	}

	tab, err := leveldb.OpenFile(home+"/remoteLookupTable", nil)
	if err != nil {
		log.Fatal(err)
	}

	return tab
}
