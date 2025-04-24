package commands

import (
	"encoding/json"
	"log"
	"net"
)

type GladCommand interface {
	GetId() int
	SetId(int)

	// This is used for mapping commands from the command line
	// arguments
	GetCommandName() string
	ParseParams(params []string) (GladCommand, error)
}

/*
Every action from client is modeled as command and each command
need to register itself with command manager.
*/
type CommandManager struct {
	commandCounter int
	commandList    []GladCommand
}

func NewCommandManger() *CommandManager {
	mgr := &CommandManager{
		commandCounter: 0,
		commandList:    make([]GladCommand, 0),
	}

	return mgr
}

func (cmdManeger *CommandManager) AddCommands(newCmd GladCommand) {
	cmdManeger.commandCounter++

	newCmd.SetId(cmdManeger.commandCounter)
	cmdManeger.commandList = append(cmdManeger.commandList, newCmd)
}

func (cmdManager CommandManager) ProcessCommand(target string, userCommand []string) {
	command := userCommand[0]

	for _, gladCommand := range cmdManager.commandList {
		if gladCommand.GetCommandName() == command {
			parsedCommand, err := gladCommand.ParseParams(userCommand[1:])

			// Sending the to remote target given there no parsing error
			if err == nil {
				sendCommand(parsedCommand, target)
			}

			break
		}
	}
}

func sendCommand(cmd GladCommand, target string) {
	// Connecting to the remote
	remote, err := net.Dial("tcp", target)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}
	defer remote.Close()

	obj, err := json.Marshal(cmd)

	// Sending the data to the server
	if err == nil {
		remote.Write(obj)
	}
}
