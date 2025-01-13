package commands

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
