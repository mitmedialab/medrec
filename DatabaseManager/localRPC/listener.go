package localRPC

import (
	"../middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/urfave/negroni"
)

//Scrypt values as recommended here: https://godoc.org/golang.org/x/crypto/scrypt
const (
	ScryptSaltBytes = 8
	ScryptN         = 32768
	ScryptR         = 8
	ScryptP         = 1
	ScryptKeyLen    = 32
	AesKeyLen       = 32
	PrivateKeyLen   = 64
)

// MedRecLocal interface for all the rpc methods
type MedRecLocal struct {
}

// ListenandServe starts the local websocket listener
func ListenandServe(router *mux.Router) {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	s.RegisterService(new(MedRecLocal), "MedRecLocal")

	n := negroni.New()
	n.UseFunc(middleware.Whitelist)
	n.UseHandler(s)

	router.Handle("/localRPC", n)
}
