package main

import (
	"fmt"
	"log"

	"github.com/tomedme/greetings"
)

func main() {
	// Set properties of the predefined logger, [...]
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it
	message, err := greetings.Hello("")
	// message := "Hello Tom"

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
