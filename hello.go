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

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Get a greeting message and print it
	// message, err := greetings.Hello("")
	// message, err := greetings.Hello("Tom")

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(message)
	fmt.Println(messages)
}
