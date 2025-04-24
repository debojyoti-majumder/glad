package server

import (
	"fmt"
	"log"
	"net"
)

func proccessLDAPClient(client net.Conn) {
	fmt.Println("Processing LDAP Request")
	client.Close()
}

func processMangementCommands(client net.Conn) {
	log.Println("Processing management command")
	client.Close()
}

func startLdapService(port int) {
	connectionString := fmt.Sprintf(":%d", port)
	log.Println("Starting server on port", connectionString)

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
	connectionString := fmt.Sprintf(":%d", port)
	log.Printf("Started management service on %s\n", connectionString)

	connection, err := net.Listen("tcp", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Closing the connection upon function end
	defer connection.Close()

	for {
		client, err := connection.Accept()
		if err == nil {
			go processMangementCommands(client)
		}
	}
}

func StartServer(port int) {
	go startMangementService(port + 1)
	startLdapService(port)
}
