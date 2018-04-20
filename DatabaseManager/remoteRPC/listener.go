package remoteRPC

import (
	// "./ethereum

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

// MedRecRemote interface for all the rpc methods
type MedRecRemote struct {
}

// ListenandServe remote connections
func ListenandServe(router *mux.Router) {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	s.RegisterService(new(MedRecRemote), "MedRecRemote")

	router.Handle("/remoteRPC", s)
}
