package models

type IParkingStrategy interface {
	AddSlot(slotNumber int)
	RemoveSlot(slotNumber int)
	GetNextSlot() int
}

type ParkingStrategy struct {
	Slots []int
}

func (p *ParkingStrategy) AddSlot(slotNumber int) {
	p.Slots = append(p.Slots, slotNumber)
}

func (p *ParkingStrategy) RemoveSlot(slotNumber int) {
	for i, slot := range p.Slots {
		if slot == slotNumber {
			p.Slots = append(p.Slots[:i], p.Slots[i+1:]...)
			break
		}
	}
}

func (p *ParkingStrategy) GetNextSlot() int {
	if len(p.Slots) == 0 {
		return -1
	}
	return p.Slots[0]
}
