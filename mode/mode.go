package mode

import (
	"ParkingLot/models"
)

type Mode interface {
	Process() models.Command
	ProcessCommand(command models.Command)
}
