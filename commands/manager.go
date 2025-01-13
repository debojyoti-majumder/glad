package commands

type GladCommand interface {
	GetId() int
	SetId(int)

	Serialize() []byte
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
