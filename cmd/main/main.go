package main

import (
	"IAGE-tcpserver/cmd/server"
	"fmt"
	"log"
)

func main() {
	fmt.Println("App started...")
	TCPServer := server.TCPServer{}
	err := TCPServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error from TCP server", err)
	}
}