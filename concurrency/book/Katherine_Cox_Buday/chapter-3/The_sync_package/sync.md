# Sync-Package

- [Reference](https://dev.to/karanpratapsingh/go-course-sync-package-5c3m)


## WaitGroup

[**Example**](/sync/waitGroup.go)

- There're three components :
   - **Add(i int)** takes in an integer value which is essentially the number of goroutines that the waitgroup has to wait for. This must be called before we execute a goroutine.
   - **Done()** is called within the goroutine to signal that the goroutine has successfully executed.
   - **Wait()** blocks the program until all the goroutines specified by Add() have invoked Done() from within.
- Important to know that a waitgroup must not be copied after first use. And if it's explicitly passed into functions, it should be done by a pointer. This is because it can affect our counter which will disrupt the logic of our program.

---

## Mutex

- A Mutex is a mutual exclusion lock that prevents other processes use the **shared data** while using other thread to prevent **race conditions** from happening.
- A **Shared data** not be run by **multiple threads** at once because the code contains shared resources.

[**Example - 1**](/sync/Mutex.go)

- Mutex has three methods
    - `Lock()` acquires or holds the lock.
	- `Unlock()` releases the lock.
    - `TryLock()` which will use like a whether their we can lock or UnLock.

[**Example - 2**](/sync/MutexTryLock.go)

- **Note:** Similar to WaitGroup a Mutex must not be copied after first use.
- **TryLock** Give a bool about which locked (Give False) if it's not locked yet (Give True)

## Once

- `Do(f func())` calls the function **f** only once,
   If `Do` is called multiple times, only the first call will invoke the function **f**.

[**Example**](/sync/Once.go)

- A `Once` must not be copied after first use.
- As we can see, even when we ran 100 goroutines, the count only got incremented once.

## Pool

[**Example-1**](/sync/pool.go)

[**Example-2**](/sync/pool2.go)


## Atomic

[**Example**](/sync/atomic.go)

- **Atomic** Using for managing state.
- We’ll use an atomic integer type to represent our (always-positive) counter.
- Simply if you have a simple **shared data** like a integer we can use **Atomic varible**,when we use 
  atomic variable we don't need a mutex lock but it will work like it has a lock in shared data.
