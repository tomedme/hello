package main

import (
	"fmt"

	"github.com/tomedme/greetings"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Tom")
	// message := "Hello Tom"
	fmt.Println(message)
}
