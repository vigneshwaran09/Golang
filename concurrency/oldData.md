# Old Data

<aside>

### What are WaitGroups?

You can use WaitGroups to wait for multiple goroutines to finish. A WaitGroup blocks the execution of a function until its internal counter becomes 0.

Let's see a simple code snippet:

```go
package main

import (
	"fmt"
)

func main() {
	go helloworld()
	go goodbye()
}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("Good Bye!")
}

```

Output

```go
$ go run HelloWorld.go

```

If we run the above program, it doesn't print anything. This is because the main function got terminated as soon as those two goroutines started executing. So, we can use `Sleep` which pauses the execution of the main function. It looks like this:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go helloworld()
	go goodbye()
	time.Sleep(2 * time.Second)
}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("Good Bye!")
}

```

Here's the output:

```go
$ go run HelloWorld.go
Good Bye!
Hello World!

```

Here, the `main` function was blocked for 2 seconds and all the goroutines were executed successfully.

Blocking the method for 2 seconds might not create any problems. But at the production level, where each millisecond is vital, millions of concurrent requests can create a huge problem.

You can solve this problem using **sync.WaitGroup** like this:

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go helloworld(&wg)
	go goodbye(&wg)
	wg.Wait()
}

func helloworld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World!")
}

func goodbye(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good Bye!")
}

```

Output

```go
$ go run HelloWorld.go
Good Bye!
Hello World!

```

The output is the same as the previous one, but it doesn't block the `main` for 2 seconds.

1. `wg.Add(int)`: This method indicates the number of goroutines to wait. In the above code, I have provided 2 for 2 different goroutines. Hence the internal counter wait becomes 2.
2. `wg.Wait()`: This method blocks the execution of code until the internal counter becomes 0.
3. `wg.Done()`: This will reduce the internal counter value by 1.

**NOTE**: If a WaitGroup is explicitly passed into functions, it should be added by a pointer.

