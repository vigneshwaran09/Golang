## About First Class Functions

In Go, functions are first-class values. This means that you can do with functions the same things you can do with all other values:

- Assign functions to variables.

## Issue of this Method

The value of an uninitialized variable of function type is nil. Therefore, calling a nil function value causes a panic.
```golang
var dutchGreeting func(string) string
dutchGreeting("Alice") // panic: call of nil function
```
Function values can be compared with nil. This can be useful to avoid unnecessary program panics.

```golang
var dutchGreeting func(string) string
if dutchGreeting != nil {
	dutchGreeting("Alice") // safe to call dutchGreeting
}
```
A function type denotes the set of all functions with the same sequence of parameter types and the same sequence of result types. User-defined types can be declared on top of function types.

```golang
type greetingFunc func(string) string

func dialog(name string, f greetingFunc) {
	fmt.Println(f(name))
	fmt.Println("I'm a dialog bot.")
}
```
**Function literal or Anonymous function** : Anonymous function means without function name.
- A function literal is a function declared inline without a name,as well as known as a anonymous function.
```golang
package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("Naruto")
	}
	a()
	fmt.Printf("%T", a)
}

```

Self-invoking function

```golang
package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("hello world first class function")
	}()
}

```

Assigning Function to Variable

```golang
package main

import "fmt"

func test(){
  fmt.Println("Tested...")
}

func main() {
	t := test // this line...
	t()
}

```
> This line creates a variable named <mark> t </mark>  
 and we're not calling the <mark> test function </mark> instead of we assigns that to <mark> t </mark>  . In simpler terms, the variable <mark> t </mark> now holds a <mark> reference </mark> to the <mark> test function </mark>. 

 > If we have parentheses <mark> () </mark> only mean we're calling the function.
 
What we do here is call the function immediately after we got the return value.

```golang
package main

import "fmt"

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	returnFunc("return")()
}

```

# Differences between First-Order Function and Higher-Order Function

| First-Order Function                                                | Higher-Order Function                                                |
|----------------------------------------------------------------------|------------------------------------------------------------------------|
| Function is treated as a variable that can be assigned to any other variable or passed as an argument.                           | The function receives another function as an argument or returns a new function or both. |

