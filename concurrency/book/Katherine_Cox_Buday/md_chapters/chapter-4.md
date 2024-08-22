# Chapter - 4

><mark> **Why do we using channels that pass empty interfaces in this chapter ?**</mark>
> - **Conciseness:** Using empty interfaces simplifies code examples, making them easier to read and understand.
> - **Generality:** Empty interfaces can hold values of any type, allowing for more flexible and adaptable code patterns.
> - **Representativeness:** In some cases, using empty interfaces better reflects the core concept of the pattern being explained.

## Confinement

When working with concurrent code, there are a few different options for safe operation. We’ve gone over two of them:
-  Synchronization primitives for sharing memory (e.g., sync.Mutex)
-  Synchronization via communicating (e.g., channels)

There are a couple of other options that are implicitly safe within multiple concurrent processes:
- Immutable data
- Data protected by confinement