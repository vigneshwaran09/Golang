- Abstraction in Go allows to hide the implementation details of a program and expose only the necessary information to the user.
- It is achieved by using interfaces which define the methods a struct should have.
- This allows the program to work with any struct that implements the interface, regardless of its specific type.

```go
package main

import "fmt"

// Shape is an interface that defines the methods a shape should have.
type Shape interface {
    area() float64
}

// Hidden Detail start

// Rectangle represents a rectangle shape.
type Rectangle struct {
    width, height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

// Circle represents a circle shape.
type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}


// Hidden Detail End

func NewShape(s Shape) { 
    fmt.Println("Area of :", s.area())
}

func main() {
    var shape Shape
    shape = Rectangle{5, 10}
    NewShape(shape)
    shape = Circle{7}
    NewShape(shape)
}
```

# Summary
**In programming, abstraction is the process of hiding the implementation details of a program** and exposing only the necessary information to the user. It allows for simplified interactions with the program by reducing the complexity of the underlying code.

# Example :
  - Based on this example we exposed only **NewShape** but NewShape doesn't know the Implementation where we call
    the function it don't care about the Implementation at all.