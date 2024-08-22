# Generator

[Sample](/concurrency/patternForInterviewAt19August/Generator/Generator.go)

- **Generator Pattern** is used to <mark>generate a sequence of values</mark> which is used to produce some output.

- In our example, we have a generator function that simply returns a channel from which we can read the values.

- This works on the fact that sends and receives block until both the sender and receiver are ready. This property allowed us to wait until the next value is requested.
- If we run this, we'll notice that we can consume values that were produced <mark>on demand</mark>.
- This is a similar behavior as <mark>yield</mark> in JavaScript and Python.

