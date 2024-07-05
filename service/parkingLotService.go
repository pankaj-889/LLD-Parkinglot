package service

import (
	. "ParkingLot/models"
	. "ParkingLot/models/strategy"
	"fmt"
)

type ParkingLotService struct {
	ParkingLot      *ParkingLot
	ParkingStrategy *ParkingStrategy
}

type IParkingLotService interface {
	CreateParkingLot(lot *ParkingLot)
	GetParkingLot() *ParkingLot
	Park(car Car)
	Leave(slotNumber int)
}

func (p *ParkingLotService) CreateParkingLot(lot *ParkingLot) {
	p.ParkingLot = lot
	for i := 0; i < p.ParkingLot.GetCapacity(); i++ {
		p.ParkingStrategy.AddSlot(i)
	}
}

func (p *ParkingLotService) GetParkingLot() *ParkingLot {
	return p.ParkingLot
}

func (p *ParkingLotService) Park(car Car) {

	slotNumber := p.ParkingStrategy.GetNextSlot()
	parkSlotNumber, err := p.ParkingLot.Park(car, slotNumber)
	if err != nil {
		return
	}
	fmt.Printf("Car is parked %d\n", parkSlotNumber.GetSlotNumber())
	p.ParkingStrategy.RemoveSlot(slotNumber)
}

func (p *ParkingLotService) Leave(slotNumber int) {
	p.ParkingLot.Leave(slotNumber)
	p.ParkingStrategy.AddSlot(slotNumber)
	fmt.Printf("Car is removed from %d\n", slotNumber)
}
