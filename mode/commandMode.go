package mode

import (
	. "ParkingLot/factory"
	. "ParkingLot/models"
	. "ParkingLot/service"
	"bufio"
	"fmt"
	"os"
)

type CommandMode struct {
	Type string
}

func NewMode(mode string) CommandMode {
	return CommandMode{Type: mode}
}

func (f *CommandMode) ProcessCommand(command Command) {
	var parkingLotService ParkingLotService
	c := NewCommandExecutorFactory(parkingLotService)
	executor, exists := c.GetCommandExecutor(command.CommandName)
	if exists != true {
		fmt.Println("Command not found")
		return
	}
	executor.Execute(command)
}

func (f *CommandMode) Process() {

	fmt.Println("Welcome to Parking Lot")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	command := NewCommand(input)
	command.Process()
	f.ProcessCommand(*command)
	return
}
