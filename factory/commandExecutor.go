package factory

import (
	. "ParkingLot/models"
)

type CommandExecutor interface {
	Execute(command Command)
}
