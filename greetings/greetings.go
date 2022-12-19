package greetings

import (
"fmt"
"errors"
"math/rand"
"time"
)

//Guess this is how to add comments to GO
//Define "Hello" to return a greeting to the user
func Hello(name string) (string, error) { //type declaration happens after variable declaration
//--PT 2--
//Add error handling, the error return type is specified on the function definition
	if name == "" {
		return "", errors.New("empty name")
	}
	//Return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	//The special ':=' operator is the definition and assignment that happens on the right side of the assignment operator in other languages
	return message, nil
	//Interesting, but looks like adding the error handling syntax
	//necessitates a return value for the error object even on a success. Probably more C type syntax.
}
//--PT 3--
//Init sets initial values for variables used in the function 
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat returns a random greeting from the slice defined herein
func randomFormat() string {
	//A 'slice' of message formats
	formats := []string {
		"Hi %v, Welcome!",
		"Hail and well met, %v!",
		"Sup, %v?",
	}

//Return the randomly selected message
	return formats[rand.Intn(len(formats))]
}


//###############################################
//Documentation makes a good point; if there's ever
//a point where you publish modules for public use,
//avoid changing method signatures, instead build a
//new function that may have diff signatures/return
//types
//###############################################
//--PT 4--

//Hellos returns a map that associates each of the named people
//with a greeting message
func Hellos(names []string) (map[string]string, error) {
	//A map to associate names with messages
	messages := make(map[string]string)
	//Loop through the received slice of names, calling the
	//Hello function to get a message for each name.
	//^^ That's kinda genius ^^
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//In the map, associate the retrieved message with
		//the name.
		messages[name] = message
	}
	return messages, nil
}
