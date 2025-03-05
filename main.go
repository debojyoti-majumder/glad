package main

import (
	"flag"
	"fmt"
	"os"

	"debojyoti.majumder/glad/commands"
	"debojyoti.majumder/glad/server"
)

func processClientCommand(target string, cmdMgr commands.CommandManager, args []string) {
	fmt.Printf("Should be sending it to %s\n", target)
	cmdMgr.ProcessCommand(args)
}

func main() {
	// The command hanlers are needed both for server and client operations
	commandManager := commands.NewCommandManger()
	commandManager.AddCommands(&commands.AddNodeToClusterCommand{})
	commandManager.AddCommands(&commands.RemoveNodeToClusterCommand{})

	// Defining the commands
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := serverCmd.Int("port", 3000, "Server port")

	clientCmd := flag.NewFlagSet("client", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Can be run in either client or server mode")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "server":
		serverCmd.Parse(os.Args[2:])
		server.StartServer(*serverPort)

	case "client":
		remoteSystem := clientCmd.String("server", "127.0.0.1:3000", "Remote address")
		clientCmd.Parse(os.Args[2:])
		commandArguments := clientCmd.Args()

		processClientCommand(*remoteSystem, *commandManager, commandArguments)
	default:
		fmt.Println("Show usage")
	}
}
