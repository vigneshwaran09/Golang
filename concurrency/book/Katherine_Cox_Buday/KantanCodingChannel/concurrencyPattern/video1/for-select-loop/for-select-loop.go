package main

import(
	"fmt"
)

func main(){
	charchannel := make(chan string,3) // Unbuffered Channel.
	chars := []string{"a","b","c"}

	for _,s := range chars{
         select {
		 case charchannel <- s:
			 fmt.Println("Case-1")
		 }
	}

	close(charchannel)

	/*
	  # we can loop over the closed channel.
	 
	*/
	 
    for result := range charchannel {
		/*
		 # Once we close the channel,there's a flag internally it will tell the channel's open or close
		   and where that channel closed (range)
		 # So when it reached the close channel it will know,it should end the loop.
		*/
		fmt.Println(result)
	}

}