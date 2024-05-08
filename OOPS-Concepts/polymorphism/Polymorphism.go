package main

import (
	"fmt"
)

type Vehicle interface {
	Car() // Method Signature
}

/*
- "struct" can only implement the "interface" functions.
*/

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

func (v *MustangCar) Car() { // Receiver function or Embedd function
	fmt.Println("Mustang Details :", v.Name, v.Rate, v.HoursPower)
}

func (v *Rolls_royceCar) Car() { // Receiver function or Embedd function
	fmt.Println("Rolls royce Details :", v.Name, v.Rate, v.HoursPower)
}
/*
- Here in this "Vehicles" function getting the data as a "interface"(Vehicles).
- so based on the "object"(struct) it's behaviour would change,that's polymorphism.
*/

func Vehicles(v Vehicle) { 
	v.Car()
}

func main() {
	mustang := MustangCar{"Mustang", 90, 500}
	RollsRoyce := Rolls_royceCar{"Rolls royce", 95, 800}
	Vehicles(&mustang)
	Vehicles(&RollsRoyce)
}
