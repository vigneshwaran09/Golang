# Composition

- In golang,there's no **Inheritance**,we can do **composition** through **Embedding** the struct inside the struct.
- It will give the **behaviour** of **Inheritance**.
- By using **Composition**,we make resuable codes.

### We can embedd a function in two ways.
- field-name with embedd-data type.
- only field-name

#### Field-name with embedd-data type
<mark>FileName => pack.go & FolderName => dto</mark>

```go
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
	Engine engine // Embedd
	Wheel wheel // Embedd
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
```
<mark>FileName => main.go & FolderName => composition</mark>

```go
package main

import(
	"fmt"
	dt "composition/dto"
)

func main(){
	c := dt.NewCar("a",600,123)
	fmt.Println(c)
	fmt.Println(c.Engine.HP())
	fmt.Println(c.Wheel.WheelDimension())
    /*
    - Here every time we need to call that "struct" for calling his Recevier function.
    */
}
```
##### Output

```go
go build
./composition
{a {600} {123}}
600
123
```

#### embedd-data type & Without Field-name

<mark>FileName => pack.go & FolderName => dto</mark>

```go
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
```

<mark>FileName => main.go & FolderName => composition</mark>

```go
package main

import(
	"fmt"
	dt "composition-type-2/dto"
)

func main(){
	c := dt.NewCar("a",600,123)
	fmt.Println(c)
	fmt.Println(c.HP())
	fmt.Println(c.WheelDimension())
}
```
##### Output

```go
{a {600} {123}}
600
123
```
- By doing this without **Field-name** the **car** struct can access all of wheel and engine struct `property & methods` directly.
- We can do this embedding using **interface** too.
- We're doing sort of `inheritance` but it's not completly `inheritance`.