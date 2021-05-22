package main

import (
	"fmt"
	"log"
	"github.com/buzzology/greetings"
)

func main() {
	// Set properties of the logger to disable printing time, source file and line number
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	// Get a greeting message and print it
	message, err := greetings.Hello("Jenny")
	if err != nil {
		log.Fatal(err)
	}

	// Display a slice of greetings this time
	messages, err := greetings.Hellos([]string {
		"Batman",
		"Flash",
		"Superman",
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, greeting := range messages {
		fmt.Println(greeting)
	}

	// No errors, safe to print
	fmt.Println(message)
}