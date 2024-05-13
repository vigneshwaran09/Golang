package main

import(
	"fmt"
	"runtime"
)

type Writer struct {
	queue chan []byte
  }
  
  func NewWriter() *Writer {
	w := &Writer{
	  queue: make(chan []byte, 10),
	}
	go w.process()
	return w
  }
  
  func (w *Writer) Write(message []byte) {
	w.queue <- message
  }
  
  func (w *Writer) process() {
	for {
	  message := <- w.queue
	  fmt.Println("Message :",message)
	  // do something with message
	}
  }


func main() {
	fmt.Println(runtime.NumGoroutine()) // 1
	/*
	- "runtime.NumGoroutine()" this package will show you,how many goroutines are running.
	- Here "main" method also a "goroutine" so it showed `1` as a value. 
	*/
	test()
	fmt.Println(runtime.NumGoroutine()) // 2
	/*
	- Once the "process()" function exposed [process<-NewWriter<-test<-main] would be 
	  running in background because we don't have a thing to stop that.
	- so "process()" goroutine exist as long as the "main" goroutine live.
	- Here "process()" and "main" method also a "goroutine" so it showed `2` as a value. 
	*/
  }
  
  func test() {
	NewWriter()
  }