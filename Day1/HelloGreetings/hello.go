package main

import (
	"fmt"
	"log"

	greetings "github.com/latoulicious/GoTrackr/Day1/Greetings"
)

func main() {
	// Set the properties of the predifined Logger , Including
	// the log entry prefix and flag to disable printing
	// the time, source profile, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greetings messages for the names
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned , print the returned map of
	// message to the console
	fmt.Println(messages)
}
