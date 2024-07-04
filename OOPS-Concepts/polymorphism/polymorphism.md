# polymorphism

```go
package main

import (
	"fmt"
)


type Human interface { // we Declare it
	spices()
// what
// how
}


type person1 struct { 
	name string
	rank string
}
type person2 struct { 
	name string
	rank string
}
type person3 struct { 
	name string
	rank string
}

func (e person1) spices() { // Method
   // embedd or recevier
	fmt.Println("Name - ", e.name, " Rank - ", e.rank)
}
func (e person2) spices() {
	fmt.Println("Name - ", e.name, " Rank - ", e.rank)
}
func (e person3) spices() {
	fmt.Println("Name - ", e.name, " Rank - ", e.rank)
}

func displayEveryHuman(h Human) { // we use that interface as parameter
	h.spices()                      // Now the struct decide the behavior of the function.
}

func main() {
	p1 := person1{name: "vicky", rank: "1"}  // initiate it
	p2 := person2{name: "venkat", rank: "2"} 
	p3 := person3{name: "karthi", rank: "3"} 
	displayEveryHuman(p1)
	displayEveryHuman(p2)
	displayEveryHuman(p3)

}
```

<aside>
👁️ struct Details

- Struct is a user-define data type.
- It can hold different type of data-types.
- It can hold only data types.
- It's similar to classes in other object-oriented programming languages.
</aside>

<aside>
👁️ interface Details

- Interface can only Implement  "Method signature"
- Decoupling or loose coupling
- **Example** -> function without parameter and return data,function with parameter and return data
</aside>

<aside>
👁️ what

- Any type implement this funtion that type is the part of "Human" interface after that implementation.
</aside>

<aside>
👁️ how

- How it's a polymorphism?
    - poly-"many",morp-"forms".
    - we Implement the "method signature" which type is implement that method that can change his behaviour based on the **struct** implementation so that's why.
</aside>

<aside>
👁️ method

If we Embedd or put struct as a receiver into some funtion that function is belong to that struct after that's called as a Method.

</aside>