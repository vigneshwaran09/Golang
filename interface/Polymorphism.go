package main

import (
	"fmt"
)

type Vehicle interface {
	Car()
}

type MustangCar struct {
	Name       string
	Rate       int
	HoursPower int
}

type Rolls_royceCar struct {
	Name       string
	Rate       int
	HoursPower int
}

func (v *MustangCar) Car() {
	fmt.Println("Mustang Details :", v.Name, v.Rate, v.HoursPower)
}

func (v *Rolls_royceCar) Car() {
	fmt.Println("Rolls royce Details :", v.Name, v.Rate, v.HoursPower)
}

func Vehicles(v Vehicle) {
	v.Car()
}

func main() {
	mustang := MustangCar{"Mustang", 90, 500}
	RollsRoyce := Rolls_royceCar{"Rolls royce", 95, 800}
	Vehicles(&mustang)
	Vehicles(&RollsRoyce)
}
