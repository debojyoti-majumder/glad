package commands

import "fmt"

// Implements GladCommand interface
type AddNodeToClusterCommand struct {
	id          int
	NodeAddress string
}

func (cmd AddNodeToClusterCommand) GetId() int {
	return cmd.id
}

func (cmd *AddNodeToClusterCommand) SetId(v int) {
	cmd.id = v
}

func (cmd AddNodeToClusterCommand) GetCommandName() string {
	return "add-node"
}

func (cmd AddNodeToClusterCommand) ProcessCommand(params []string) {
	fmt.Println("Adding node to the cluster")
	fmt.Println(params)
}

// Implements GladCommand interface
type RemoveNodeToClusterCommand struct {
	id int
}

func (cmd RemoveNodeToClusterCommand) GetId() int {
	return cmd.id
}

func (cmd *RemoveNodeToClusterCommand) SetId(v int) {
	cmd.id = v
}

func (cmd RemoveNodeToClusterCommand) GetCommandName() string {
	return "remove-node"
}

func (cmd RemoveNodeToClusterCommand) ProcessCommand(params []string) {
	fmt.Println("Removing node from the cluster")
	fmt.Println(params)
}
