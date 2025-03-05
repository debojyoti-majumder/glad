package commands

type GladCommand interface {
	GetId() int
	SetId(int)

	// This is used for mapping commands from the command line
	// arguments
	GetCommandName() string
	ProcessCommand(params []string)
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

func (cmdManager CommandManager) ProcessCommand(userCommand []string) {
	command := userCommand[0]

	for _, gladCommand := range cmdManager.commandList {
		if gladCommand.GetCommandName() == command {
			gladCommand.ProcessCommand(userCommand[1:])
			break
		}
	}
}
