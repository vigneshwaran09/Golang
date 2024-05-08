package person

/*
# Note:
   - For Every Encapsulation class should have only struct,get,set.
   - We shouldn't use "Private" values over sample class or package,if we did there's no use for get,set method
     so avoid to use.
   - The main purpose of get and set is avoiding the datas to get dirrectly access from other class or packges.
*/

/* 
 - This "struct" not accessible for all place where ever we import this package. 
 - Because we didn't "Export" this struct.
 - We can use this "struct" only within this package.
*/

type person struct{ // Private or Not Exported
	firstName string // Private or Not Exported
	lastName string // Private or Not Exported
	rollNo int // Private or Not Exported
}


// Constructor
func NewPerson (fn,ln string,rollNo int)*person{
	return &person{
		firstName: fn, 
		lastName: ln,
		rollNo : rollNo,
	}
}

func(p *person)GetfirstName()string{ // For Get the Private data.
	return p.firstName
}

func(p *person)GetlastName()string{ // For Get the Private data.
	return p.lastName
}

func(p *person)GetrollNo()int{ // For Get the Private data.
	return p.rollNo
}

func(p *person)SetfirstName(fn string){ // For Set the Private data.
	p.firstName = fn
}

func(p *person)SetlastName(ln string){ // For Set the Private data.
	p.lastName = ln
}

func(p *person)SetrollNo(no int){ // For Set the Private data.
	p.rollNo = no
}