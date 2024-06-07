# Goroutine vs threads

> [best reference video](https://youtu.be/YHRO5WQGh0k?si=DY4N2dzylzRZAerS)<br>
> [reference video](https://youtu.be/UNtSB-dprIM?si=SK9T7kIGKjJ-pmpx)<br>
> [reference blog](https://github.com/ntk148v/lets-go/blob/master/tips-notes/goroutines.md)

- Goroutine are much little (small) than threads.
- Goroutine consuming less memory than threads.
- Threads are slow.
- Goroutines are fast.
> Why threads are slow ?
> - Threads are take more than 1MP of memory for each thread. 
> - when a thread is created which has to store lot of information like'
[tag-1]
> - A single thread has to store lot of information even before starting.
> - To create or delete or remove a thread,we need to call a **OS** for each and everytime.
[tag-2]
> - Doing **OS** calls is much costlier than you expect this's also a another reason for **threads** are heavy or slow. 

> Why Goroutines are fast ?
> - Goroutines runs inside the **Go-runtime**.
[tag-3]
> - while create or delete the **Goroutine** we're not calling the **OS** instead we create or delete goroutine inside **Go-runtime**.
> Calling **OS** isn't a problem but calling much frequently is a problem.
> Threads has so many data's to store (pointer's) to create a single thread.
> But goroutine stores only few information which cost less than 2KB per goroutine.

> How does golang do this?
> - Go has **scheduler** which take care of creating or deleting or handling **goroutines** efficiently.


# just a summary :

- Goroutines and threads aren't same.
- Goroutines are user-space threads.
- threads are managed by OS (kernel) but Goroutines are managed by **Goruntime**.
- Goroutines are lighter weight and faster than kernel threads.
- operating system only knows how to schedule to put kernel threads on the hardware (on your CPU core)
- Go schedular put goroutines on kernal threads which run on CPU.
- single thread can hold nth number of Goroutines