package models

type Slot struct {
	parkedCar  Car
	slotNumber int
}

type ISlot interface {
	GetParkedCar() Car
	GetSlotNumber() int
	Assign(car Car)
	UnAssign()
}

func (s *Slot) GetParkedCar() Car {
	return s.parkedCar
}

func (s *Slot) GetSlotNumber() int {
	return s.slotNumber
}

func (s *Slot) Assign(car Car) {
	s.parkedCar = car
}

func (s *Slot) UnAssign() {
	s.parkedCar = Car{}
}

func NewSlot(slotNumber int) *Slot {
	return &Slot{
		slotNumber: slotNumber,
	}
}
