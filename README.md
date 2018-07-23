# MedRec
##### _Patient controlled medical records_

---

https://medrec.media.mit.edu

### Components

- Database Manager - API written in GoLang that provides access to an underlying database. R/W access is governed by permissions stored on the blockchain.
- Ethereum Client - A pointer to the go-ethereum codebase
- SmartContracts - The Solidity contracts and their tests that are used by other MedRec components
- UserClient - A front-facing node app that can be used by any party to interact with the MedRec system.

This project is being developed under the GPLv2 LICENSE.

## For production

#### 2. Install NPM

Install NPM: https://nodejs.org/en/

#### 3. Install npm packages

Setup the UserClient
```
$ cd UserClient
$ npm install
$ npm run build
$ cd ..
```

To resolve an annoying won't fix bug in the bitcoin-mnemonic library you also need to run the following in the UserClient directory.

```
$ rm -r node_modules/bitcore-mnemonic/node_modules/bitcore-lib
```

Setup the Javscript helper files
```
$ cd GolangJSHelpers
$ npm install
$ npm run build
$ cd ..
```

#### 4. Setup your MySQL Database

You need to be running a mysql  database instance locally with username:password root:medrecpassword:
- run query `/scripts/medrec-v1.sql`. It will create a schema called `medrec-v1` for you to store/retrieve information from. It is representing the "remote" DB.
- run query `/scripts/medrecwebapp.sql` for the "local" DB.

If you do not want to be able to look at the example patient records you can skip this section.

#### 5. Start all of MedRec's components

```
$ ./medrec EthereumClient
$ ./medrec DatabaseManager
$ ./medrec UserClient
```

## For Development
#### 1. Setup a PoA blockchain

Use Goethereum to setup a proof of authority blockchain

##### 2. Connect to the blockchain

MedRec has to be modified to connect to the provider nodes of this blockchain. Edit the medrec-genesis.json and startGeth.js in GolangJSHelpers/ so that the parameters match with your network.

#### 3. Install Go and golang libraries

Install Go: https://github.com/golang/go or `brew install go`

and run:
```
$ go get
```

#### 4. Install NPM

Install NPM: https://nodejs.org/en/

#### 5. Install npm packages

```
$ cd UserClient
$ npm install
$ npm run build
```

To resolve an annoying won't fix bug in the bitcoin-mnemonic library you also need to run the following in the UserClient directory.

```
$ rm node_modules/bitcore-mnemonic/node_modules/bitcore-lib
```

#### 6. Run the Database manager

You need to be running a mysql  database instance locally with username:password root:medrecpassword:
- run query `/scripts/medrec-v1.sql`. It will create a schema called `medrec-v1` for you to store/retrieve information from. It is representing the "remote" DB.
- run query `/scripts/medrecwebapp.sql` for the "local" DB.

`$ go run main.go DatabaseManager`

It should start the DatabaseManager running on localhost:6337

#### 7. Deploy the contracts
You will need to install the program truffle using npm to deploy contracts.

In the `SmartContracts` director: `truffle deploy`. The `ganache-cli` should respond to this command showing that the contracts have been deployed.

#### 8. Start the UserClient

```
$ cd UserClient
$ npm start
```

if it throws an error with `web3-requestManager`, do `npm install web3@1.0.0-beta.26`.

#### 9. Building a new production version

```
$ go build
$ cd UserClient
$ npm run build
```


# Getting Started With the MedRec Codebase

## 1. Link MedRec to an existing database
In order to use MedRec as a medical records provider, MedRec's Database Manager needs to be linked to the provider's database. This does not require a change in the database itself, but does involve changing MedRec to fit the form of database used. This might involve more or less work (from changing a few table names, up to completely rewriting queries) depending on the form and structure of database used. In order to link MedRec to your database, the data will need to be directly queryable by an external piece of software -- e.g., using SQL or other style queries.

### SQL-type databases
MedRec uses the Golang mysql driver (github.com/go-sql-driver/mysql) to interface with its test database. There is support for a wide range of common hospital database types, including


- FHIR: https://golanglibs.com/top?q=fhir


### Non-SQL-type databases
There's a <a href="https://github.com/avelino/awesome-go#database-drivers">pretty exhaustive list</a> of Golang drivers for different forms of database. Depending on the database

### Setup Guide
All the files that need to be changed are located in the `DatabaseManager/remoteRPC` directory. This manages remote access to your database by people granted access to view a particular record.

1. Instantiate the database:
In `databse.go`,

- SQL (but not MySQL): replace the database driver for mysql (`import "github.com/go-sql-driver/mysql"`) with the driver for your database format
- NOSQL (/other) -- also get rid of `import "database/sql"`

Once the driver and database type has been set, replace the contents of `instantiateDatabase()` with your own server.

```
func instantiateDatabase() *sql.DB {

  db, err := sql.Open("mysql",
    "<user>:<password>@tcp(127.0.0.1:<port>)/<schema>")
  if err != nil {
    log.Fatal(err)
  }

  return db
}

```

2. Setting up the eth key-value store
The Database Manager uses a LevelDB key-value store to link the unique patient identifier used by MedRec (the patient's ethereum address) to some unique identifier in the record provider's database. This is managed in `lookup.go`.

A unique ID from the database needs to be sent to the prospective patient, so that they can enter it, along with the creation of an ethereum address, when they create an account. This then creates an entry in the LevelDB (`Lookup Table`) that maps the patient's AgentID (some unique patient ID) to their eth address. Depending on the type of UID that is used to identify the patient, you might want to change the type of args.AgentID to match that of the stored value.
```
type AccountArgs struct {
    Account string
    AgentID string
}

...

  err := tab.Put([]byte(args.Account), []byte(args.AgentID), nil)
  if err != nil {
    log.Println(err)
  }
```

NB -- you *don't* need to add the eth address of the agent -- that gets added automatically when an account is made. To test the functionality of this, navigate to the frontend and create a new patient account, with some UID stored in your database. The records returned for that patient should match those associated with that UID.


3. Requesting Documents
The management of this part of MedRec will vary with the kind of information the record provider should make available to each patient. The generic function `PatientDocuments` in `documentRequest.go` provides an outline for receiving a call from a patient's account, checking that their eth address is listed in the key-value store, and replying with the contents of that request (`reply.Documents`).

In order to adapt this to specific instances, the names of specific tables must be changed to align the provider database with the records requested. In `PatientDocuments` these are specified as:

```
  var (
    PatientID   string
    DocumentID int
    DocDateTime string
    PracticeID string
    RecvdDateTime string
  )

```

The names of the result (`params.Document`, defined in `params.go`) remain the same -- the entries in the provider database that correspond to them should be changed to match.

4. The rest of the Database Manager (overview)
Instructions 1, 2, and 3 unpack the 'provider-facing' layers of the Database Manager, which interact directly with the provider database. The rest of this repo (`DatabaseManager/RemoteRPC`) deals with the blockchain interface.

In particular `ethereum.go` handles authentication and connection to the ethereum client.
- Authentication: `Recover(r *http.Request, args *RecoverArgs, reply *RecoverReply)` recovers

## 2. Joining the blockchain as a provider


## 3. Write a custom contract
