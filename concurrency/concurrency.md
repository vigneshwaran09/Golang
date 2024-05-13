**If you need to get better at Concurrency patter you should get well at these three :**
- go routine
- channel
- select

# Which is the entry point of GO program .
   <mark>main</mark> function

<aside>
<h3>Goroutine life span</h3>
    - In Go <code>main</code> function also a <code>goroutine</code> which is a parent Goroutine once the parent goroutine finish all other goroutine also Exit.
</aside>

---

<aside>

[https://www.freecodecamp.org/news/concurrent-programming-in-go/](https://www.freecodecamp.org/news/concurrent-programming-in-go/)

![https://www.freecodecamp.org/news/content/images/size/w2000/2022/12/2-1.png](https://www.freecodecamp.org/news/content/images/size/w2000/2022/12/2-1.png)

**Concurrency** refers to a programming language's ability to deal with lots of things at once.

A good way to understand concurrency is by imagining multiple cars traveling on two lanes. Sometimes the cars overtake each other, and sometimes they stop and let others pass by.

Another good example is when your computer runs multiple background tasks like messaging, downloading movies, running the operating system, and so on – all at once.

**Parallelism** means doing lots of things simultaneously and independently. It might sound similar to concurrency, but it’s actually quite different.

Let's understand it better with the same traffic example. In this case, cars travel on their own road without intersecting each other. Each task is isolated from all other tasks. Concurrent tasks can be executed in any given order.

This is a non-deterministic way to achieve multiple things at once. True parallel events require multiple CPUs.

![https://www.freecodecamp.org/news/content/images/2022/12/1-1.png](https://www.freecodecamp.org/news/content/images/2022/12/1-1.png)

Illustration showing difference between parallelism and concurrency


#### Example1
```go
package main

import (
	"fmt"
)

func main(){
    someFunc("01")
	fmt.Println("a.... HI")
}

func someFunc(num string){
	fmt.Println(num)
}
```
> If you want to call this `someFunc` we should call that into the
 `main` function because that's entry point of the go program.
> 

<dl>
<dt>What happen!</dt>
  <dd>When we call this <code>someFunc</code> it runs synchronously that means our <code>main</code> function is going to be block until the completion of this <code>someFunc</code> function.  </dd>
  <dd>So below <code>print</code> statement won't run until after <code>someFunc</code> is finished running <dd>
  <dd>But with help of <mark>go routine</mark> we can actually make <code>someFunc</code> <mark>fork off</mark> of our
      <code>main</code> function <mark>asynchronously</mark>  and that would mean that <code>main</code> would not need to wait 
      for <code>someFunc</code> to finish before it could continue with its work </dd>
</dl>

![](diagrams/synchronous.drawio.svg)

**Golang follows a model of concurrency called `fork-Join`** 

## What is a Goroutine?

A goroutine is an independent function that executes simultaneously in some separate lightweight threads managed by Go. GoLang provides it to support concurrency in Go.

```go
package main

import (
	"fmt"
)

func main() {
	go helloworld()
	goodbye()
}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("Good Bye!")
}

```
![](diagrams/fork_join.drawio.svg)

Without time.Sleep():

```
$ go run HelloWorld.go
Good Bye!

```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go helloworld()
	time.Sleep(1 * time.Second)
	goodbye()
}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("Good Bye!")
}

```
![](diagrams/sleep.drawio.svg)

After adding time.Sleep(), the `helloworld` goroutine is able to finish its execution before main exits:

```
$ go run HelloWorld.go
Hello World!
Good Bye!

```
- But this way is not a proper way for sync our goroutine with main goroutine because we couldn't tell accurate processing time for all goroutine function in time.Sleep,some will take too long or some will take too small.

### What are Channels?

![**As shown in the image**](img/Untitled-Diagram46.jpg)

- Go provides **channels** that you can use for bidirectional communication between goroutines.
- Bidirectional communication means that one goroutine will send a message and the other will read it. <mark>Sends and receives are blocking</mark>. Code execution will be stopped until the write and read are done successfully.

There are a couple different types of channels:

- **Unbuffered channel**: Unbuffered channels require both the sender and receiver to be present to be successful operations. It requires a goroutine to read the data, otherwise, it will lead to deadlock. By default, channels are unbuffered.

- **Buffered channel**: Buffered channels have the capacity to store values for future processing. The sender is not blocked until it becomes full and it doesn't necessarily need a reader to complete the synchronization with every operation.

  - If a space in the array is available, the sender can send its value to the channel and complete its send operation immediately.

  - After its execution, if a receiver comes, the channel will start sending values to the receiver and it will start its operation once it receives the values. As the sender and receiver are operating at different times, this is called `asynchronous communication`.

### Creating a Channel

- A channel is created using **chan** keyword
- It can only transfer data of the same type, different types of data are not allowed to transport from the same channel.

```go
Syntax to declare a channel
ch := make(chan Type) // Bidirectional default.

```

```go
Declaration of channels based on directions
1. Bidirectional channel : chan T
2. Send only channel: chan <- T
3. Receive only channel: <- chan T

```

# Difference between declaration of variable and make

```go
// how to create a channel
package main

import "fmt"

func main() {

	// Creating a channel

	// Using var keyword

	var mychannel chan int
 /*
 # a channel of type int is declared using the var keyword without initializing it.         By default, when a channel is declared without initialization, its value is nil.
*/
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	// Creating a channel using make() function

	mychannel1 := make(chan int)
/*
# a channel of type int is created using the make() function. The make() function initializes and returns a new, initialized channel. When you use make(chan int), it creates a new channel and assigns its memory address to the variable mychannel1.
*/
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)

}
```

**Output:**

```
Value of the channel:
Type of the channel: chan int
Value of the channel1:  0x432080
Type of the channel1: chan int

```

### Bidirectional Channel

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string) // declaring a channel.
	go greet(msg) // Goroutine
	greeting := <-msg // We read here,It will block until someOne send data here.
	fmt.Println("Greeting received")
	fmt.Println(greeting)
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")
	ch <- "Hello vignesh"  // We write here.
	fmt.Println("Greeter completed")
}

```

```go
$ go run main.go
Greeter waiting to send greeting!
Greeter completed
Greeting received
Hello vignesh

```

<aside>
👁️ UniDirectional

```go
package main

import "fmt"

// Function to send values into a send-only channel
func sendData(ch chan<- int, value int) {
    ch <- value
    close(ch)
}

// Function to receive values from a receive-only channel
func receiveData(ch <-chan int, done chan<- bool) {
    for num := range ch {
        fmt.Println("Received:", num)
    }
    done <- true
}

func main() {
    // Creating a send-only channel
    sendCh := make(chan<- int)

    // Creating a receive-only channel
    receiveCh := make(<-chan int)

    // Creating a bidirectional channel
    ch := make(chan int)

    // Assigning the send-only channel to the bidirectional channel
    sendCh = ch

    // Assigning the receive-only channel to the bidirectional channel
    receiveCh = ch

    // Creating a channel to signal the completion of receiving
    done := make(chan bool)

    // Sending a value into the send-only channel
    go sendData(sendCh, 42)

    // Receiving values from the receive-only channel
    go receiveData(receiveCh, done)

    // Waiting for receiving to complete
    <-done
}
```

- An unidirectional channel refers to a channel that is restricted to either sending or receiving values. It means that the channel can only be used in one direction, either for sending or for receiving, but not both.
</aside>

**Closing the channel**: 
- Closing the channel indicates that no more values should be sent through this channel.
  - Because we want to show that the work has been completed so there is no need to keep a channel open.

```go
package main

import (
	"fmt"
)

func main() {
	msg := make(chan string)
	go greet(msg)
	/*
	 - By using below syntax,we can know the channel is open or close by using second variable.
	 - First variable give value.
	*/
	greeting, ok := <-msg
	if ok {
		fmt.Println("result :",greeting)
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello vignesh"
	close(ch)
	/*
	  - In above,we Close the channel means there's no more value can send through this channel
	    but we can read from this channel even after closing also too.
	*/

	fmt.Println("Greeter completed")
}

```

We close a channel by using `close()` like `close(ch)` on the above code snippet.

```go
$ go run main.go
Greeter waiting to send greeting!
Greeter completed
result : Hello vignesh
Channel is open!
```

<aside>
♻️ Here we make a Unbuffered channel then use the channel like a sync.WaitGroup

```go
package main

import (
	"fmt"
)

var count int

func PrintData(i int, ch chan int) {
	fmt.Println(i)
	ch <- 1 // Send
}

func main() {
	ch := make(chan int) // Unbuffer
	/*
	 - Defaultly channels are Unbuffer.
	*/
	for i := 1; i <= 5; i++ {
		go PrintData(i, ch)
	}
	<-ch // We read here,It will block until someOne send data here.This act like a sync.WaitGroup
	<-ch
	<-ch
	<-ch
	<-ch
}
```

- Channnel are synchronised one when we put a channel which will waiting for other end either it would be receiver or sender.
- Here we put a Receiver channel in `main` function make main goroutine to wait until their respective channel come,Here we make five receiver so it wait for five sender to send a value.
</aside>

### Channel without goroutine

> Causing error because of use `channel` without goroutine in "Unbuffered Channel”
> 

```go
    package main
    
    import "fmt"
    
    func WithoutGoroutineInUnBufferedChan() {
    	c := make(chan int) // Unbuffered Channel
    	c <- 17
    	fmt.Println(<-c)
    }
```    
<aside>
    ♻️ Reason
    
- But the reason we got deadlock is Go runtime or go scheduler detect somehow when we don’t have a respective sender or receiver in our program that will be terminated by go runtime.
- The main goroutine wait for the receiver receive a value on `c <- 17` but there’s no receiver at the time because the receiver we made next line once the sender opertion complete then only program execute the next line.
</aside>

# Select

```go

package main

import (
	"fmt"
)

func SendData(ch chan string){
      ch <- "Hello"
}

func main(){
     ch1 := make(chan string)
	 ch2 := make(chan string)

	 go SendData(ch1)

	 go func(){
         ch2 <- "Hiiii"
	 }()

	 select {
	 case data1 := <-ch1:
		fmt.Println("Getting data from channel-1 :",data1)
     case data2 := <-ch2:
		fmt.Println("Getting data from channel-2 :",data2)
	 }
}

```

- We can wait multiple goroutine in select statement if one of the case success then it end the process like a switch statement.