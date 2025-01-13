package server

import (
	"fmt"

	"debojyoti.majumder/glad/commands"
)

func initCommands() {
	fmt.Println("Registering commands with command manager")
	cmdMgr := commands.NewCommandManger()

	cmdMgr.AddCommands(&commands.AddNodeToClusterCommand{})
	cmdMgr.AddCommands(&commands.RemoveNodeToClusterCommand{})
}

func StartServer(port int) {
	initCommands()
	fmt.Println("Starting server on port", port)
}
