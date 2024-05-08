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
	engine // Embedd //  only data-type & without Field-Name
	wheel // Embedd //  only data-type & without Field-Name
}

// Constructor
func NewCar(CarName string,hp,wd int)Car{
	return Car{
		CarName:CarName,
		engine:engine{
			hp:hp,
		},
		wheel:wheel{
			wheelDimensions:wd,
		},
	}
}