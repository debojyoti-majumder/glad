package commands

import (
	"fmt"
	"net"
)

// Implements GladCommand interface
type AddNodeToClusterMessage struct {
	Id          int
	NodeAddress string
}

func (cmd AddNodeToClusterMessage) GetId() int {
	return cmd.Id
}

func (cmd *AddNodeToClusterMessage) SetId(v int) {
	cmd.Id = v
}

func (cmd AddNodeToClusterMessage) GetCommandName() string {
	return "add-node"
}

func (cmd AddNodeToClusterMessage) ParseParams(params []string) (GladCommand, error) {
	// Parsing the user input
	inp := params[0]
	_, err := net.ResolveTCPAddr("tcp", inp)
	if err != nil {
		fmt.Println(err)
	}

	return &AddNodeToClusterMessage{Id: cmd.Id, NodeAddress: inp}, nil
}

// Implements GladCommand interface
type RemoveNodeToClusterMessage struct {
	Id          int
	NodeAddress string
}

func (cmd RemoveNodeToClusterMessage) GetId() int {
	return cmd.Id
}

func (cmd *RemoveNodeToClusterMessage) SetId(v int) {
	cmd.Id = v
}

func (cmd RemoveNodeToClusterMessage) GetCommandName() string {
	return "remove-node"
}

func (cmd RemoveNodeToClusterMessage) ParseParams(params []string) (GladCommand, error) {
	// Parsing the user input
	inp := params[0]
	_, err := net.ResolveTCPAddr("tcp", inp)
	if err != nil {
		return &cmd, err
	}

	return &RemoveNodeToClusterMessage{Id: cmd.Id, NodeAddress: inp}, nil
}
