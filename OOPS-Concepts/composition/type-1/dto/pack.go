package dto

type engine struct{
	 hp int // HoursePower
}

func(e engine)HP()int{ // Getter function based on "Encapsulation".
	return e.hp
}

type wheel struct{
	wheelDimensions int 
} 

func(w wheel)WheelDimension()int{ // Getter function based on "Encapsulation".
	return w.wheelDimensions
}

// Composition
type Car struct{
/*
- Having a struct inside another struct is called composition.
*/
	CarName string
	Engine engine // Embedd // Field-Name with data-type
	Wheel wheel // Embedd // Field-Name with data-type
}

// Constructor
func NewCar(CarName string,hp,wd int)Car{
	return Car{
		CarName:CarName,
		Engine:engine{
			hp:hp,
		},
		Wheel:wheel{
			wheelDimensions:wd,
		},
	}
}