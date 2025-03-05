package commands

import (
	"fmt"
	"net"
)

// Implements GladCommand interface
type AddNodeToClusterCommand struct {
	Id          int
	NodeAddress string
}

func (cmd AddNodeToClusterCommand) GetId() int {
	return cmd.Id
}

func (cmd *AddNodeToClusterCommand) SetId(v int) {
	cmd.Id = v
}

func (cmd AddNodeToClusterCommand) GetCommandName() string {
	return "add-node"
}

func (cmd AddNodeToClusterCommand) ParseParams(params []string) (GladCommand, error) {
	// Parsing the user input
	inp := params[0]
	_, err := net.ResolveTCPAddr("tcp", inp)
	if err != nil {
		fmt.Println(err)
	}

	return &AddNodeToClusterCommand{Id: cmd.Id, NodeAddress: inp}, nil
}

// Implements GladCommand interface
type RemoveNodeToClusterCommand struct {
	Id          int
	NodeAddress string
}

func (cmd RemoveNodeToClusterCommand) GetId() int {
	return cmd.Id
}

func (cmd *RemoveNodeToClusterCommand) SetId(v int) {
	cmd.Id = v
}

func (cmd RemoveNodeToClusterCommand) GetCommandName() string {
	return "remove-node"
}

func (cmd RemoveNodeToClusterCommand) ParseParams(params []string) (GladCommand, error) {
	// Parsing the user input
	inp := params[0]
	_, err := net.ResolveTCPAddr("tcp", inp)
	if err != nil {
		return &cmd, err
	}

	return &RemoveNodeToClusterCommand{Id: cmd.Id, NodeAddress: inp}, nil
}
