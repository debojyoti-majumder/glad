package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	messages "debojyoti.majumder/glad/messages"
	"debojyoti.majumder/glad/server"
)

func processClientCommand(target string, cmdMgr messages.CommandManager, args []string) {
	log.Printf("Sending command to %s\n", target)

	// Process only if there are other parameters are present
	if len(args) != 0 {
		cmdMgr.ProcessCommand(target, args)
	}
}

func main() {
	// The command hanlers are needed both for server and client operations
	commandManager := messages.NewCommandManger()
	commandManager.AddCommands(&messages.AddNodeToClusterMessage{})
	commandManager.AddCommands(&messages.RemoveNodeToClusterMessage{})

	// Defining the commands
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := serverCmd.Int("port", 3000, "Server port")

	clientCmd := flag.NewFlagSet("client", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Can be run in either client or server mode")
		os.Exit(1)
	}

	// Setting log location as standard output location
	log.SetOutput(os.Stdout)

	command := os.Args[1]

	switch command {
	case "server":
		serverCmd.Parse(os.Args[2:])
		server.StartServer(*serverPort)

	case "client":
		remoteSystem := clientCmd.String("server", "127.0.0.1:3001", "Remote address")
		clientCmd.Parse(os.Args[2:])
		commandArguments := clientCmd.Args()

		processClientCommand(*remoteSystem, *commandManager, commandArguments)
	default:
		fmt.Println("Show usage")
	}
}
