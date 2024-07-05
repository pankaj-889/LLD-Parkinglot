package factory

import (
	. "ParkingLot/service"
)

type CommandExecutorFactory struct {
	commands map[string]CommandExecutor
}

func NewCommandExecutorFactory(parkingLotService ParkingLotService) *CommandExecutorFactory {
	factory := make(map[string]CommandExecutor)

	factory["create_parking_lot"] = NewCreateParkingLotCommandExecutor(parkingLotService)
	factory["park"] = NewParkingCommandExecutor(parkingLotService)
	factory["leave"] = NewLeaveCommandExecutor(parkingLotService)
	return &CommandExecutorFactory{commands: factory}
}

func (f *CommandExecutorFactory) GetCommandExecutor(command string) (CommandExecutor, bool) {
	executor, exists := f.commands[command]
	if exists != true {
		return nil, false
	}
	return executor, exists
}
