package models

type Command struct {
	Command       string
	CommandName   string
	CommandParams []string
}

func NewCommand(command string) *Command {
	return &Command{Command: command, CommandName: "", CommandParams: []string{}}
}

func (command *Command) Process() {

	var commandName string
	var commandParams string

	input := command.Command
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			commandName = input[:i]
			commandParams = input[i+1:]
			break
		}
	}

	var list []string
	for i := 0; i < len(commandParams); i++ {
		if commandParams[i] == ' ' || i == len(commandParams)-1 {
			list = append(list, commandParams[:i])
			commandParams = commandParams[i+1:]
			i = -1
		}
	}

	command.CommandName = commandName
	command.CommandParams = list
	return
}
