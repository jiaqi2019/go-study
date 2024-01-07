package main

import (
	"fmt"
	"log"

	"error/greetings"
)

func main() {
	message, err := greetings.Hello("random")
	if err != nil {
		log.Fatal(err)
		// log.Flags()
	}
	fmt.Println(message)
}
