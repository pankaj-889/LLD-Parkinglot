package factory

import (
	. "ParkingLot/models"
	. "ParkingLot/service"
	"fmt"
	"log"
	"strconv"
)

type CreateParkingLotCommandExecutor struct {
	parkingLotService ParkingLotService
}

func (c *CreateParkingLotCommandExecutor) Execute(command Command) {
	capacity := command.CommandParams[0]
	capacityInt, err := strconv.Atoi(capacity)
	if err != nil {
		log.Fatal("Error while converting capacity to int")
	}
	parkingLot := NewParkingLot(capacityInt)
	c.parkingLotService.CreateParkingLot(parkingLot)
	fmt.Println("Created a parking lot with", command.CommandParams[0], "slots")
}

func NewCreateParkingLotCommandExecutor(parkingLotService ParkingLotService) *CreateParkingLotCommandExecutor {
	return &CreateParkingLotCommandExecutor{parkingLotService: parkingLotService}
}
