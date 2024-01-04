package main

import (
	"fmt"
	"log"

	"error/greetings"
)

func main() {
	names := []string{
		"Alice",
		"Bob",
		"Joker",
		"Jack",
	}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
		// log.Flags()
	}
	fmt.Println(message)
}
