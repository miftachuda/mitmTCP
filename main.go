package main

import (
	"log"
)

var (
	listener GameConnection
)

func main() {
	remoteHost = "127.0.0.1"
	remotePort = "109"
	initConnections()

}

func initConnections() {
	listener = GameConnection{}
	log.Println("Starting listener...")
	InitListener(&listener)
}
