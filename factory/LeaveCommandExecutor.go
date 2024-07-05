package factory

import (
	. "ParkingLot/models"
	. "ParkingLot/service"
	"fmt"
)

type LeaveCommandExecutor struct {
	parkingLotService ParkingLotService
}

func (c *LeaveCommandExecutor) Execute(command Command) {
	c.parkingLotService.CreateParkingLot(command)
	fmt.Println("Created a parking lot with", command.CommandParams[0], "slots")
}

func NewLeaveCommandExecutor(parkingLotService ParkingLotService) *CreateParkingLotCommandExecutor {
	return &CreateParkingLotCommandExecutor{parkingLotService: parkingLotService}
}
