#!/bin/bash

#setup the build directory
rm -r build
mkdir build

#get node
cp "$(which node)" build

#build the Database Manager and entry point to the app
go build -o medrec-amd64 main.go
cp medrec-amd64 build

#copy the ethereum client starter app
mkdir build/EthereumClient
cp EthereumClient/main.js build/EthereumClient/
cp EthereumClient/clientBinaries.json build/EthereumClient/
cp -r EthereumClient/node_modules build/EthereumClient/

#build the UserClient
# (cd UserClient && npm run build)
mkdir build/UserClient
cp -r UserClient/build build/UserClient
cp UserClient/electron-starter.js build/UserClient
cp UserClient/node_modules/electron-prebuilt/cli.js build/UserClient/electron
