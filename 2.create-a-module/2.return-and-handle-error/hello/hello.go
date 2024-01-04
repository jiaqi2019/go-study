package main

import (
	"fmt"
	"log"

	"error/greetings"
)

func main() {
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
		// log.Flags()
	}
	fmt.Println(message)
}
