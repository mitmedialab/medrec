package manager

import (
	"fmt"
	"net/http"

	"./ethereum"
	"./localRPC"
	"./middleware"
	"./remoteRPC"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Init() {
	n := negroni.New()
	router := mux.NewRouter()

	//two different RPC clients are used to prevent accidental leaks of private functions
	remoteRPC.ListenandServe(router)
	localRPC.ListenandServe(router)
	ethereum.Init()

	n.UseFunc(middleware.EnableCORS)
	n.UseFunc(middleware.Logger)
	n.UseHandler(router)

	listenString := fmt.Sprintf("127.0.0.1:%d", 6337)
	http.ListenAndServe(listenString, n)
}
