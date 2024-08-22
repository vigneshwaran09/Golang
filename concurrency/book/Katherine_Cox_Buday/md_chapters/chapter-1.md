# Chapter - 1

## Why Is Concurrency Hard?
 1) you can use the **go** keyword to run a function concurrently. Doing so creates what’s called a **goroutine**.

### Race Conditions :
1) **Data Race**
> <mark>Data Race</mark><br>
> [**sample-1**](/concurrency/book/Katherine_Cox_Buday/chapter-1/Why_Is_Concurrency_Hard/Race_Conditions/data_race.go)
> - What’s called a data race, where one concurrent
operation attempts to read a variable while at same time another concurrent operation is attempting to write to the same variable.
>
> - It shows up as a **data race**, when two threads are competing for the same variable or resource.

> <mark>time.Sleep</mark><br>
> [**sample-2**](/concurrency/book/Katherine_Cox_Buday/chapter-1/Why_Is_Concurrency_Hard/Race_Conditions/data_race_with_timeSleep.go)
> - **time.sleep** is not a proper way to prevent the **race condition** because we don'know concurrency function how long will it take so **time.sleep** a temporary solution.

### Atomicity :
- In Go atomic operations,are operations that are desgined to be exectuted without interruption and without interference from other concurrenct operation.
> [Reference blog](https://medium.com/@hatronix/go-go-lockless-techniques-for-managing-state-and-synchronization-5398370c379b#:~:text=Atomic%20operations%20are%20a%20crucial,efficient%20and%20reliable%20concurrent%20code.) 
>
> ```go
> package main
>
>import (
> "fmt"
> "sync/atomic"
> "time"
>)
>
>func main() {
> var counter int32
>
> // Create a goroutine to increment the counter.
> go func() {
>  for i := 0; i < 5; i++ {
>   atomic.AddInt32(&counter, 1) // To add (AddInt32)
>   fmt.Printf("Incremented: %d\n", atomic.LoadInt32(&counter)) // To read (AddInt32)
>   time.Sleep(time.Millisecond)
>  }
> }()
>
> // Create a goroutine to decrement the counter.
> go func() {
>  for i := 0; i < 5; i++ {
>   atomic.AddInt32(&counter, -1)
>   fmt.Printf("Decremented: %d\n", atomic.LoadInt32(&counter))
>   time.Sleep(time.Millisecond)
>  }
> }()
>
> // Wait for the goroutines to finish.
> time.Sleep(2 * time.Second)
>
> fmt.Printf("Final Value: %d\n", atomic.LoadInt32(&counter))
>}
>
>// Decremented: -1
>// Incremented: 0
>// Incremented: 1
>// Decremented: 0
>// Decremented: -1
>// Incremented: 0
>// Incremented: 1
>// Decremented: 0
>// Decremented: -1
>// Incremented: 0
>// Final Value: 0
>```
> - Atomic operation use for only when we have a shareable memory (Variable).
> - It ensure the safety of shared data without requiring a mutex for synchronization.
> - We can do atomic operations in integer and pointers only.
> - The key advantage of using atomic operations is that they are typically more efficient than using traditional mutexes,especially for simple
 operations on primitive data type like integer and pointers.
> - Atomic opertions operate on the hardware level,ensuring that the operations are performed atomically,preventing data races and
avoiding the need for traditional looking mechanisms.
> - It's important to note that atomic operations have their limitations,they are best suite for simple,low-level on primitive data types.if
you need to perform more complex operations that involve multiple variable or require more complex synchronization,you may still need to
use mutexes or other synchronizations primitives.
> ##### Summary
> 1) Atomic also same as mutex but for simple opeartions atomic is more efficient than mutex.
> 2) When one goroutine doing process on shared memory (variable) another one is wait for the process done
then it will do his process on shared variable.

### Memory Access Synchronization :
>  ```go
>   package main
>
>    import (
>        "fmt"
>        "sync"
>    )
>
>    func main() {
>        var memoryAccess sync.Mutex // <1>
>        var value int
>        go func() {
>            memoryAccess.Lock() // <2>
>            value++
>            memoryAccess.Unlock() // <3>
>        }()
>
>        memoryAccess.Lock() // <4>
>        if value == 0 {
>            fmt.Printf("the value is %v.\n", value)
>        } else {
>            fmt.Printf("the value is %v.\n", value)
>        }
>        memoryAccess.Unlock() // <5>
>    }
>```
> - There’s a name for a section of your program that needs <mark>exclusive access to a shared resource</mark>. This is called a **critical section**.
> - We have three critical sections in this program:
>   1) Our goroutine, which is incrementing the data variables.
>   2) Our if statement, which checks whether the value of data is 0.
>   3) Our fmt.Printf statement, which retrieves the value of data for output.
> - One of the way to solve this problem is to <mark>synchronize access to the memory</mark> between your critical sections.
> Code(Lock and Unlock) between those two statements can then assume it has exclusive access to data, we have successfully synchronized access to the memory.
> - You may have noticed that while we have solved our data race but we haven’t actually solved our race condition! The order of operations in this program is still nondeterministic.
> Synchronizing access to the memory in this manner also has make a poor performance.
> But the calls to Lock and Unlock you see can make our program slow. Every time we perform one of these operations, our program pauses for a period
of time.<br>
> <mark>This brings up two questions ( We should ask to ourselfs ) : </mark><br>
> 1) Are my critical sections entered and exited repeatedly? ( If yes then we should considered for performance issue in our code. ) <br>
> 2) What size should my critical sections be? ( The more will grow,it will do less performance. )<br>
> - Answering these two questions in the context of your program is an art, and this adds to the difficulty in synchronizing access to the memory.
> ##### Summary
> - When there is a data race, we can wrap that critical section (when we need exclusive access to a shared resource) with a mutex, locking and unlocking the mutex. When we do this, we are synchronising memory access.
> - This fixes the data race, but it doesn't fix the race condition (threads still compete for first write/read).
> - It can create maintenance and performance problems. We should try to keep them only to critical sections. If it gets inefficient, make them broader.
> - We should use `Lock` and `UnLock` where we try to read and write on shared memory (Variable). 

### Deadlocks, Livelocks, and Starvation :

#### Deadlock : 
>```go
>package main
>
>import (
>	"fmt"
>	"sync"
>	"time"
>)
>
>type value struct {
>	mu    sync.Mutex
>	value int
>}
>
>func main() {
>	var wg sync.WaitGroup
>	printSum := func(v1, v2 *value) {
>		defer wg.Done()
>		v1.mu.Lock() // 1
>		defer v1.mu.Unlock() // 2
>		time.Sleep(2 * time.Second) // 3
>		v2.mu.Lock()
>		defer v2.mu.Unlock()
>		fmt.Printf("sum=%v\n", v1.value+v2.value)
>	}
>	var a, b value
>	wg.Add(2)
>	go printSum(&a, &b) // 4
>	go printSum(&b, &a) // 5
>	wg.Wait()
>}
>
>```
> **Output :**
>```go
>fatal error: all goroutines are asleep - deadlock!
>
>goroutine 1 [semacquire]:
>sync.runtime_Semacquire(0xc000078040?)
>        /usr/local/go/src/runtime/sema.go:62 +0x25
>sync.(*WaitGroup).Wait(0x60?)
>        /usr/local/go/src/sync/waitgroup.go:116 +0x48
>main.main()
>        /home/vigneswaran/Documents/software/git/advance_golang_concepts/> >concurrency/book/Katherine_Cox_Buday/chapter-1/Why_Is_Concurrency_Hard/Deadlocks/>Deadlocks.go:29 +0x15c
>
>goroutine 6 [sync.Mutex.Lock]:
>sync.runtime_SemacquireMutex(0xc00005af08?, 0x95?, 0x8365497a3e?)
>        /usr/local/go/src/runtime/sema.go:77 +0x25
>sync.(*Mutex).lockSlow(0xc0000120c0)
>        /usr/local/go/src/sync/mutex.go:171 +0x15d
>sync.(*Mutex).Lock(...)
>        /usr/local/go/src/sync/mutex.go:90
>main.main.func1(0xc0000120b0, 0xc0000120c0)
>        /home/vigneswaran/Documents/software/git/advance_golang_concepts/>concurrency/book/Katherine_Cox_Buday/chapter-1/Why_Is_Concurrency_Hard/Deadlocks/>Deadlocks.go:21 +0xec
>created by main.main in goroutine 1
>        /home/vigneswaran/Documents/software/git/advance_golang_concepts/>concurrency/book/Katherine_Cox_Buday/chapter-1/Why_Is_Concurrency_Hard/Deadlocks/>Deadlocks.go:27 +0xf2
>
>goroutine 7 [sync.Mutex.Lock]:
>sync.runtime_SemacquireMutex(0xc00005b708?, 0x95?, 0x83654967b1?)
>        /usr/local/go/src/runtime/sema.go:77 +0x25
>sync.(*Mutex).lockSlow(0xc0000120b0)
>        /usr/local/go/src/sync/mutex.go:171 +0x15d
>sync.(*Mutex).Lock(...)
>        /usr/local/go/src/sync/mutex.go:90
>main.main.func1(0xc0000120c0, 0xc0000120b0)
>        /home/vigneswaran/Documents/software/git/advance_golang_concepts/>concurrency/book/Katherine_Cox_Buday/chapter-1/Why_Is_Concurrency_Hard/Deadlocks/>Deadlocks.go:21 +0xec
>created by main.main in goroutine 1
>        /home/vigneswaran/Documents/software/git/advance_golang_concepts/>concurrency/book/Katherine_Cox_Buday/chapter-1/Why_Is_Concurrency_Hard/Deadlocks/>Deadlocks.go:28 +0x152
>exit status 2
>```
>- A deadlocked program is one in which all concurrent processes are waiting on one another. In this state, the program will never recover without outside intervention.
>- The Go runtime attempts to do its part and will detect some deadlocks (all goroutines must be blocked, or “asleep”1).
>- (3) Here we sleep for a period of time to simulate work (and trigger a deadlock).
>- Why?(deadlock happens) If you look carefully, you’ll see a timing issue in this code.
> ![](Katherine_Cox_Buday/img/Demonstration%20of%20a%20timing%20issue%20giving%20rise%20to%20a%20deadlock.png)
>  **Demonstration of a timing issue giving rise to a deadlock**
>- The boxes represent functions, the horizontal lines calls to these functions, and the vertical bars lifetimes of the function at the head of the graphic
> ![Demonstration of a timing issue giving rise to a deadlock](Katherine_Cox_Buday/diagram/deadlock.drawio.svg)
> - mutex normally wait for another theard complete the process on the shared memory.
> - Here both are wait forever (infinity) so we get deadlock.
> - The `go runtime` would find all thread waiting so long somehow then gave a deadlock.
> ##### Summary
> - Happens when a shared mutex is locked twice. by two different routines.
> - **Coffman Conditions help identity when we have a deadlock:**
>> 1) Mutual exclusions - A concurrent process holds exclusive rights to a resource at any one time.
>> 2) Wait for condition - When a process holds a resource and simultaneously is waiting for another resource
>> 3) No preemption - when a resource held by the process can only be released by that process
>> 4) Circular wait - a routine waiting for a routine that is waiting for that routine (Example : A concurrent process (P1) must
 be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1)).<br>
>> ![](Katherine_Cox_Buday/img/sharedMemory.png)
>> ![](Katherine_Cox_Buday/img//deadlockcircle.png)

#### Livelocks :
##### time.Tick()
> **Syntax :**
 >> `func Tick(d Duration) <- chan Time` 
 >> - <mark>d</mark> is the duration of time for ticker to tick.
 >> - <mark>chan</mark> is the ticking channel (We will get return as a channel
 => time.Time(type))
 >> - <mark>return value :</mark> It return current date and actual time after regualr
   interval of time.
>```go
>package main
>
>import(
>	"fmt"
>	"time"
>	"reflect"
>)
>
>func main(){
>	var count int
>	for ch := range time.Tick(1 * time.Millisecond) {
>		if count == 5 {
>			break
>		}
>		fmt.Println("Hellooo...")
>		fmt.Println("Hellooo... ch value:", ch, "ch type:", reflect.TypeOf(ch))
>		count++
>	}
>	fmt.Println("Process done...")
>}
>```
> **Output :**
>```go
>Hellooo...
>Hellooo... ch value: 2024-05-28 07:59:41.047197184 +0530 IST m=+0.001099531 ch type: time.Time
>Hellooo...
>Hellooo... ch value: 2024-05-28 07:59:41.048325763 +0530 IST m=+0.002228110 ch type: time.Time
>Hellooo...
>Hellooo... ch value: 2024-05-28 07:59:41.049429194 +0530 IST m=+0.003331542 ch type: time.Time
>Hellooo...
>Hellooo... ch value: 2024-05-28 07:59:41.050505726 +0530 IST m=+0.004408070 ch type: time.Time
>Hellooo...
>Hellooo... ch value: 2024-05-28 07:59:41.051580067 +0530 IST m=+0.005482416 ch type: time.Time
>Process done...
>```
> - Return `nil` if d <=0 .
> - So we get a error here.
>```go
>package main
>
>import(
>	"fmt"
>	"time"
>	"reflect"
>)
>
>func main(){
>	var count int
>	for ch := range time.Tick(-1 * time.Millisecond) {
>		if count == 5 {
>			break
>		}
>		fmt.Println("Hellooo...")
>		fmt.Println("Hellooo... ch value:", ch, "ch type:", reflect.TypeOf(ch))
>		count++
>	}
>	fmt.Println("Process done...")
>}
>```
>```go
>fatal error: all goroutines are asleep - deadlock!
>
>goroutine 1 [chan receive (nil chan)]:
>main.main()
>        /home/vigneswaran/Documents/software/git/advance_golang_concepts/concurrency/book/Katherine_Cox_Buday/chapter-1/Why_Is_Concurrency_Hard/Livelock/timeTick.go:11 +0x8a
>exit status 2
>```