package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {
	//Call locally created imported function
//--PT 2--
//Imported the "log" default Go packages for logging, implement them now
log.SetPrefix("greetings: ")
log.SetFlags(0)

//Request a greeting
	message, err := greetings.Hello("Archer")
	//call error handling before attempting to finish successfully
	if err != nil {
		log.Fatal(err)
	}

	//Print function results to term if no errors are returned
	fmt.Println(message)
}
