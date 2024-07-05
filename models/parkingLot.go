package models

import "errors"

type ParkingLot struct {
	capacity int
	slots    map[int]Slot
}

type IParkingLot interface {
	GetCapacity() int
	GetSlots() map[int]Slot
	GetSlot(slotNumber int) (Slot, error)
	Park(car Car, slotNumber int) (Slot, error)
	Leave(slotNumber int)
}

func (p *ParkingLot) GetCapacity() int {
	return p.capacity
}

func (p *ParkingLot) GetSlots() map[int]Slot {
	return p.slots
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		slots:    make(map[int]Slot),
	}
}

func (p *ParkingLot) GetSlot(slotNumber int) (Slot, error) {
	if slotNumber > p.capacity || slotNumber <= 0 {
		return Slot{}, errors.New("slot number is invalid")
	}

	slot, exists := p.slots[slotNumber]
	if exists != true {
		p.slots[slotNumber] = *(NewSlot(slotNumber))
	}
	return slot, nil
}

func (p *ParkingLot) Park(car Car, slotNumber int) (Slot, error) {
	slot, err := p.GetSlot(slotNumber)
	if err != nil {
		return Slot{}, errors.New("slot is already occupied")
	}
	slot.Assign(car)
	return slot, nil
}

func (p *ParkingLot) Leave(slotNumber int) {
	slot, err := p.GetSlot(slotNumber)
	if err != nil {
		return
	}
	slot.UnAssign()
}
