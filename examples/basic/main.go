package main

import (
	"fmt"
	"log"

	beve "github.com/beve-org/beve-go"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Create a person
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	// Marshal to BEVE format
	data, err := beve.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Encoded %d bytes\n", len(data))

	// Unmarshal back
	var decoded Person
	if err := beve.Unmarshal(data, &decoded); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decoded: %+v\n", decoded)
}
