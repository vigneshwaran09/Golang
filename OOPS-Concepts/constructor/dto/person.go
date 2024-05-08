package person

/* 
 - This "struct" accessible for all place where ever we import this package. 
 - Because we "Export" this struct.
*/
type Person struct{
	firstName string
	lastName string
	rollNo int
}

/* 
 - This "struct" not accessible for all place where ever we import this package. 
 - Because we didn't "Export" this struct.
 - We can use this "struct" only within this package.
*/

type human struct{
	firstName string
	lastName string
	rollNo int
}

// Constructor
func NewPerson (fn,ln string,rollNo int)Person{
	return Person{
		firstName: fn, 
		lastName: ln,
		rollNo : rollNo,
	}
}

// Constructor
func NewHuman (fn,ln string,rollNo int)human{
	return human{
		firstName: fn, 
		lastName: ln,
		rollNo : rollNo,
	}
}