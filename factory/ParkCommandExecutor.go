package factory

import (
	. "ParkingLot/models"
	. "ParkingLot/service"
	"fmt"
)

type ParkingCommandExecutor struct {
	parkingLotService ParkingLotService
}

func (c *ParkingCommandExecutor) Execute(command Command) {
	c.parkingLotService.CreateParkingLot(command)
	fmt.Println("Created a parking lot with", command.CommandParams[0], "slots")
}

func NewParkingCommandExecutor(parkingLotService ParkingLotService) *CreateParkingLotCommandExecutor {
	return &CreateParkingLotCommandExecutor{parkingLotService: parkingLotService}
}
