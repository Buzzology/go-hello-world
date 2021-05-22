package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Returns a map of people to greeting messages
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	// Generate a greeting for each name. We don't need the index so it is discarded
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// Associate retrieved message with name
		messages[name] = message
	}

	return messages, nil
}


// Go runs init functions automatically at program startup after global vars initialised.
func init() {
	rand.Seed(time.Now().UnixNano()) // We need to seed rand
}


// randomFormat returns a random greeting message. 
func randomFormat() string {
	// A slice of message formats.
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Randomly select one of the messages to return
	return formats[rand.Intn(len(formats))] // intn returns a random non-negative integer between 0 and n (inclusive)
}