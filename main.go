package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Defining the commands
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := serverCmd.Int("port", 3000, "Server port")

	clientCmd := flag.NewFlagSet("client", flag.ExitOnError)
	remoteSystem := clientCmd.String("server", "127.0.0.1:3000", "Remote address")

	if len(os.Args) < 2 {
		fmt.Println("Can be run in either client or server mode")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "server":
		serverCmd.Parse(os.Args[2:])
		fmt.Println("Running in server mode on", *serverPort)

	case "client":
		clientCmd.Parse(os.Args[2:])
		fmt.Println("Running on client mode", *remoteSystem)

	default:
		fmt.Println("Show usage")
	}
}
