package greetings

import "fmt"

//Guess this is how to add comments to GO
//Define "Hello" to return a greeting to the user
func Hello(name string) string { //type declaration happens after variable declaration
	//Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Greetings, %v. Welcome!", name)
//The special ':=' operator is the definition and assignment that happens on the right side of the assignment operator in other languages
	return message
}
