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

func (cmd AddNodeToClusterCommand) Serialize() []byte {
	return nil
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

func (cmd RemoveNodeToClusterCommand) Serialize() []byte {
	return nil
}
