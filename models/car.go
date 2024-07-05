package models

type Car struct {
	number int
	color  string
}

type ICar interface {
	GetNumber() int
	GetColor() string
}

func (c *Car) GetNumber() int {
	return c.number
}

func (c *Car) GetColor() string {
	return c.color
}

func NewCar(number int, color string) *Car {
	return &Car{
		number: number,
		color:  color,
	}
}
