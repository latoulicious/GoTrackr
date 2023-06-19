package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was give , return error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// create message using random format
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of named people
// with a greetings message
func Hellos(name []string) (map[string]string, error) {
	// a map associate name with a message
	messages := make(map[string]string)
	// loop through the received slice of names, calling
	// the hello function to get messages for each name
	for _, name := range name {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with
		// the name
		messages[name] = message
	}
	return messages, nil
}

// RandomFormat returns one of a set of greetings message. The returned
// message is selected at random
func randomFormat() string {
	// a slice of message format
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format  by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
