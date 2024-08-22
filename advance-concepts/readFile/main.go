package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	output := "prodlogfile.txt"          // Define the name of the output file
	input, err := os.Open("logfile.txt") // Open the input file "logfile.txt"
	var inp []string                     // Declare a slice to store the lines read from the input file
	if err != nil {                      // Check if there was an error opening the input file
		log.Println("Getting error on Input file.") // Log an error message if there was an error
	}
	defer input.Close()                // Ensure the input file is closed when the main function finishes
	scanner := bufio.NewScanner(input) // Create a new scanner to read the input file line by line
	for scanner.Scan() {               // Iterate through each line in the input file
		inp = append(inp, scanner.Text()) // Append the line to the slice `inp`
	}
	out, err := os.Create(output) // Create the output file "prodlogfile.txt"
	if err != nil {               // Check if there was an error creating the output file
		log.Println("Getting error on while writing the file.") // Log an error message if there was an error
	}
	defer out.Close()              // Ensure the output file is closed when the main function finishes
	writer := bufio.NewWriter(out) // Create a new buffered writer for the output file
	for _, logData := range inp {  // Iterate through each line stored in the slice `inp`
		_, err := writer.WriteString(logData + "\n") // Write the line to the output file, adding a newline character
		if err != nil {                              // Check if there was an error writing to the output file
			log.Println("Getting error on while writing the file.") // Log an error message if there was an error
		}
	}
	/*
	  - writer.Flush() : When you use a buffered writer, data is written to an in-memory buffer first.
	  - Instead of writing each piece of data immediately to the file, data is stored in a buffer and written in larger chunks.
	  - This reduces the number of system calls, which are relatively expensive operations.
	*/
	writer.Flush() // Flush any buffered data to the output file to ensure all data is written
}
