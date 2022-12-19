package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	//Call locally created imported function
	message := greetings.Hello("Archer")
	fmt.Println(message)
}
