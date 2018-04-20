we can now launch requests to the go backend through a web socket! This uses https://github.com/gorilla/websocket to communicate with an electron (currently plain html) frontend.

```
$ go get github.com/googollee/go-socket.io
$ go get github.com/go-sql-driver/mysql
$ go build
$ ./build
```
everything should come up on localhost:3000.

Modules:
- remoteRPC accepts/denies requests for access to a provider database of PPI
- localRPC connects to a local cache and user login store
- assets folder -- houses a demo test web app (index.html, jquery, etc.)
- ethereum -- code for making queries to smart contracts on the blockchain
- params - contains constants and struct definitions
