package main

func main(){
	 // Closed channel
	 closedChan := make(chan int)
	 close(closedChan)
	 go func(){
        <- closedChan
	 }()
	 // Attempting to send to a closed channel will cause a panic
	 closedChan <- 1 // Uncomment to see panic

}