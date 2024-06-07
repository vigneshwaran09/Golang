**<aside>♻️ Concurrency vs Parallelism</aside>**

✍🏿 From Offical Document

- Concurrency is not parallelism, although it enables parallelism.
- If you have only one processor, your program can still be concurrent but it cannot be parallel.
- On the other hand, a well-written concurrent program might run efficiently in parallel on a multiprocessor. That property could be important.
</aside>

<aside>
♻️ Dobuts

- If we run a go program ,will it run on single processor and single thread even that processor has multiple thread ,will it work on single thread at the time.
    - By default, when you run a Go program, it starts with a single operating system thread, even if the processor has multiple threads. This single OS thread is used to execute goroutines concurrently. This means that, at any given moment, only one goroutine can be running on the single OS thread. However, the Go scheduler will handle the switching between goroutines, making it appear as if they are running in parallel.
    - While a single OS thread is the default, the Go runtime can dynamically create additional OS threads as needed. For example, if a goroutine blocks on a system call or a blocking operation, the Go scheduler may create a new OS thread to continue executing other goroutines concurrently on a separate thread. This is known as "OS thread spinning."
    - It's important to note that the Go scheduler handles the distribution of goroutines across the available OS threads. This allows the program to fully utilize multiple CPU cores if they are available. As a result, the program can achieve true concurrency across multiple CPU cores even if it starts with a single OS thread.
    
    In summary:
    
    - By default, a Go program starts with a single operating system thread, even if the processor has multiple threads.
    - The Go scheduler multiplexes goroutines onto this single OS thread to allow concurrent execution.
    - The Go runtime can dynamically create additional OS threads if required to handle blocking operations or to optimize concurrency across multiple CPU cores.
    
    The Go runtime and scheduler are designed to make concurrent programming efficient and straightforward, allowing developers to focus on writing concurrent code without having to worry about managing threads explicitly.
    
</aside>

- Dual Processor System
    
    [DUAL CPU SERVER FOR OUR OFFICE | Unboxing Super Micro Server | வடக்கு நண்பரின் Server!](https://www.youtube.com/watch?v=jkNBWShQL40)
    
- concurrency vs Parallelism
    
    [Concurrency vs Parallelism](https://www.youtube.com/watch?v=Y1pgpn2gOSg)
    
- Concurrency is NOT Parallelism
    
    [Concurrency is NOT Parallelism](https://www.ics.uci.edu/~rickl/courses/ics-h197/2014-fq-h197/talk-Wu-Concurrency-is-NOT-parallelism.pdf)
    

**Details**

Here's a simplified overview of how the Go scheduler switches Goroutines:

1. **Goroutines and the Run Queue:**
    - Goroutines are the lightweight, user-space threads in Go. They are functions or methods that run concurrently.
    - The Go runtime maintains a "run queue," which is a data structure that holds a list of runnable Goroutines waiting to be executed.
    - When you create a new Goroutine (using the **`go`** keyword), it is added to the run queue.
2. **OS Threads (M:N Mapping):**
    - The Go runtime manages a small pool of OS threads (also known as M-threads). The number of OS threads is typically equal to the number of CPU cores available on the machine but can be adjusted using the **`GOMAXPROCS`** environment variable or the **`runtime.GOMAXPROCS`** function.
    - Goroutines are not directly bound to OS threads; instead, multiple Goroutines can be multiplexed onto a smaller number of OS threads. This is known as M:N mapping, where M Goroutines are multiplexed onto N OS threads.
3. **Work Stealing:**
    - When an OS thread finishes executing a Goroutine (for example, it hits a blocking operation or yields the CPU voluntarily), it requests more work from the scheduler.
    - If there are Goroutines in its own local run queue, the thread takes one of them and continues execution.
    - If its local run queue is empty, the thread attempts to "steal" work from other OS threads' run queues. This mechanism is known as "work stealing" and helps balance the load among OS threads.
4. **Context Switching:**
    - When the Go scheduler switches from one Goroutine to another (either due to work stealing or a Goroutine yielding the CPU), it performs a context switch.
    - A context switch involves saving the current state of the Goroutine (registers, stack pointer, etc.) and restoring the state of the Goroutine to be executed next.
    - Context switching in Go is designed to be lightweight and efficient, which is crucial for the performance of concurrent programs.

The Go scheduler continuously repeats these steps, efficiently managing the execution of Goroutines across a pool of OS threads. This allows Go programs to take advantage of multiple cores and handle a large number of concurrent tasks without relying on one-to-one thread-to-Goroutine mappings. The details of the scheduler's implementation and optimizations are continuously evolving as the Go language evolves and as newer versions of Go are released.


Thread Vs Goroutine

| Thread | Goroutine |
| --- | --- |
| managed by os | manage by go run time |
| Fixed stack - 1Mb | Fixed stack - 2kb |

### Concurrency and parallelism

[What is the difference between concurrency and parallelism?](https://stackoverflow.com/questions/1050222/what-is-the-difference-between-concurrency-and-parallelism)

[Blob-Master](https://github.com/Criviere/os/blob/master/Chapter4.md)

</aside>

- Summary
    - Concurrency and parallelism are two related but distinct concepts in programming.
        - Concurrency refers to the ability of a program to handle multiple tasks simultaneously, but not necessarily at the same time.
        - Parallelism refers to the ability of a program to perform multiple tasks at the same time.

Examples:

- Concurrency in Go is achieved through the use of goroutines, which are lightweight threads that run in the same address space as the main program.
    - Example: A web server handling multiple client requests concurrently

```
Copy code
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", r.URL.Path)
    })

    http.ListenAndServe(":8080", nil)
}

```

- Parallelism in Go is achieved through the use of the **`go`** keyword, which allows you to run a function in a separate goroutine.
    - Example: A program that calculates the value of pi in parallel using Monte Carlo method

```
Copy code
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func monteCarloPi(n int) float64 {
    inside := 0
    for i := 0; i < n; i++ {
        x := rand.Float64()
        y := rand.Float64()
        if x*x+y*y <= 1 {
            inside++
        }
    }
    return 4 * float64(inside) / float64(n)
}

func main() {
    rand.Seed(time.Now().UnixNano())
    n := 10000000
    pi := monteCarloPi(n)
    fmt.Println("Pi:", pi)
}

```

Note:

- Goroutines are lightweight, and thousands or even millions of goroutines can be run concurrently on a modern machine with limited memory overhead.
- The Go runtime automatically schedules goroutines on available cores, so that your program can take full advantage of parallelism on multi-core systems.

# Just a quick summary :
[reference1](https://youtu.be/w3z4_aksN5s?si=KpVHqNqeIXyFATCW)
[reference2](https://youtu.be/r2__Rw8vu1M?si=KYmpZZjwY1cbgfwc)
[reference3](https://youtu.be/oV9rvDllKEg?si=FvbU8XfEvY0W9Rpf)
[reference4](https://youtu.be/olYdb0DdGtM?si=rmBPlNHOrFgnMTOu)
- We can't achieve parallism in single-processor with single core.
- But we can achieve parallism in single-processor with multicore core.
- We can also achieve parallism in multiple-processorx.