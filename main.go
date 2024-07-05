package main

import (
	model "ParkingLot/mode"
)

func main() {

	mode := model.NewMode("command")
	mode.Process()
}
