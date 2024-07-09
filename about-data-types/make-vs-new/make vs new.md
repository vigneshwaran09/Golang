# Make VS new

<aside>
💡 Add value to the pointer slice

```go
package main

import (
	"fmt"
)

func main()  {
	v := new([]int)
	(*v)[0] = 2
	fmt.Println("v",v) 
}
```

```go
panic: runtime error: index out of range [0] with length 0

goroutine 1 [running]:
main.main()
        /home/netxd-temporary-user/Old_Laptop_Backup/workSpace/Project/NewToGoMySystem/makeVsnew/prt-1/main.go:10 +0x9f
exit status 2
```

- you're using the **`new()`** function to create a pointer to an empty slice of integers. However, the length of the slice is initially 0, so accessing index 0 will result in an out-of-range error.
- Initially, the pointer slice length is zero so we can’t directly give the value to index zero because it’s length is zero.
</aside>

<aside>
💡 Add value to the make slice

- In make function holds three values type, length, and capacity.
- we should give two values at least otherwise we will get an error.

```go
package main

import "fmt"

func main() {
	v := make([]int,0) // Initialize a slice with length 0
	v[0] = 1           // Assign a value to the first element
	fmt.println("v", v)
}
```

- If we give len at Zero, we will get an error while adding data through the index method.
- Because its length is zero we try to add data that index[length] doesn’t exist

```go
panic: runtime error: index out of range [0] with length 0

goroutine 1 [running]:
main.main()
        /home/netxd-temporary-user/Old_Laptop_Backup/workSpace/Project/NewToGoMySystem/makeVsnew/prt-2/main.go:7 +0x2f
exit status 2
```

</aside>

<aside>
💡 Append function

<aside>
<img src="https://www.notion.so/icons/anchor_blue.svg" alt="https://www.notion.so/icons/anchor_blue.svg" width="40px" /> Dynamically change the size by using the append in new

```go
package main

import (
	"fmt"
)

func main()  {
	v := new([]int) // Initialize a slice with length 1
	*v = append(*v, 1)
	fmt.Println("v", v)
}
```

```go
v &[1]
```

</aside>

<aside>
<img src="https://www.notion.so/icons/anchor_blue.svg" alt="https://www.notion.so/icons/anchor_blue.svg" width="40px" /> Dynamically change the size by using the append in make

```go
package main

import(
	"fmt"
)

func main()  {
	v := make([]int,0) 
	v = append(v, 1) // Dynamically memory expanding
	fmt.Println("v", v)
}
```

```go
v [1]
```

- It's important to note that when using **`append`**, if the underlying array capacity is exceeded, a new larger array will be allocated, and the elements will be copied to the new array. This is a mechanism in Go to handle the dynamic growth of slices.
- when you use the **`append`** function, it allows you to add elements to the slice, dynamically increasing its length as needed.
</aside>

</aside>

<aside>
💡 Paraentheses

```go
package main

import "fmt"

func main() {
	// Using make to create and initialize a slice

	s1 := make([]int, 5) // Which returns reference of the value.

	fmt.Println("s1:", s1)

	// Using new to create a pointer to a new and zeroed slice

	s2 := new([]int) // Which return pointer of value

	fmt.Println("s2:", s2)

	// Assigning values to the slice using the pointer

  // We give all the value directly to the s2 through **dereference** 

	*s2 = make([]int, 3)

  /*
    However, to access or modify the value of the slice through the pointer,
    you need to **dereference** it using ***s2**.
  */	

fmt.Println("s2Make:", *s2)

	(*s2)[0] = 1

/*
   # If you want to give a value through the Index way you should do this.
   # In summary, the use of parentheses (*s2) is necessary to dereference a pointer
     and access the underlying value. This is required when assigning values to the
     slice through the pointer or when performing operations on the slice using the
     pointer.
   # After assigning values to the slice using (*s2), you can access and modify
     individual elements of the slice using (*s2)[index].
   # Paraentheses only need when you add a value to the pointer variable by index way.
*/

	(*s2)[1] = 2

	(*s2)[2] = 3

	fmt.Println("s201:", *s2)
   /*
      # But we can also use append like this.
   */ 

  *s2 = append(*s2, 4)

	fmt.Println("s202:", *s2)

// Make a Slice by using new 

	s3 := new([]int)

	*s3 = append(*s3,1)

	fmt.Println("s3",s3)

}
```

<aside>
💡 Reference

- Refer
    - one
        
        [Why would I make() or new()?](https://stackoverflow.com/questions/9320862/why-would-i-make-or-new)
        
    - Two
        
        [Golang Make and New Keywords - Scaler Topics](https://www.scaler.com/topics/golang/golang-make-and-new-keywords/)
        
</aside>

</aside>