package server

import (
	"fmt"
	"log"
	"net"
)

func proccessLDAPClient(client net.Conn) {
	fmt.Println("Processing client connection")
	client.Close()
}

func startLdapService(port int) {
	connectionString := fmt.Sprintf(":%d", port)
	fmt.Println("Starting server on port", port)

	// Opening the port connection
	connection, err := net.Listen("tcp", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Deffering the connection closure
	defer connection.Close()

	for {
		client, err := connection.Accept()
		if err != nil {
			fmt.Println("Not able to process client")
			continue
		}

		// Once the connection is accepted packet parsing is done in
		// aysnc way
		go proccessLDAPClient(client)
	}
}

func startMangementService(port int) {
	fmt.Printf("Started management service on port %d\n", port)
}

func StartServer(port int) {
	go startMangementService(port + 1)
	startLdapService(port)
}
