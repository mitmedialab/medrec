package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/mitmedialab/medrec/DatabaseManager"
	"github.com/mitmedialab/medrec/DatabaseManager/common"
)

var offButton chan bool
var userClient *exec.Cmd
var ethClient *exec.Cmd

func runDatabaseManager() {
	manager.Init()
}

func runEthereumClient() {
	//assign a common group id to the child processes
	ethClient.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	//start the ethereum client
	err := ethClient.Run()
	if err != nil {
		log.Print("Ethereum client Exited")
		log.Println(ethClient.Args)
	}
	offButton <- true
}

func runUserClient() {
	//assign a common group id to the child processes
	userClient.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	//start the front facing app for users
	userClient.Start()

	err := userClient.Wait()
	if err != nil {
		log.Print("User client Exited")
		log.Print(err)
	}
	offButton <- true
}

func main() {
	offButton := make(chan bool)
	ethClient = common.NodeExec("./GolangJSHelpers/startGeth.js")
	ethClient.Stdout = os.Stdout
	ethClient.Stderr = os.Stderr
	userClient = exec.Command("UserClient/node_modules/electron-prebuilt/cli.js", "UserClient/electron-starter.js")

	if len(os.Args) > 1 {
		if os.Args[1] == "DatabaseManager" {
			go runDatabaseManager()
		}
		if os.Args[1] == "UserClient" {
			go runUserClient()
		}
		if os.Args[1] == "EthereumClient" {
			go runEthereumClient()
		}
	} else {
		//start the database manager to provide access management services for the underlying database
		go runDatabaseManager()

		//initialize the eth client
		go runEthereumClient()

		//initialize the user client
		go runUserClient()
	}

	//handle an interrupt via ^C or some other means
	interruptButton := make(chan os.Signal, 1)
	signal.Notify(interruptButton, os.Interrupt)

	//wait for a stop notification then completely stop this and all child processes
	select {
	case <-offButton:
	case <-interruptButton:
	}

	//stop the direct process and the child process group for each
	if ethClient.Process != nil {
		ethClient.Process.Kill()
		syscall.Kill(-ethClient.Process.Pid, syscall.SIGKILL)
	}
	if userClient.Process != nil {
		userClient.Process.Kill()
		syscall.Kill(-userClient.Process.Pid, syscall.SIGKILL)
	}
}
