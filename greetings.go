package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message .
	message := fmt.Sprintf(randGreeting(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randGreeting() string {
	greetings := []string{"Hello %v", "Hi %v", "Greetings %v", "Salutations %v", "Good day %v", "Good evening %v", "Good morning %v"}
	return greetings[rand.Intn(len(greetings))]
}
