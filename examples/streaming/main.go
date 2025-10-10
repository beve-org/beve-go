package main

import (
	"bytes"
	"fmt"
	"log"

	beve "github.com/beve-org/beve-go"
)

type Message struct {
	ID        int
	Type      string
	Payload   []byte
	Timestamp int64
}

func main() {
	messages := []Message{
		{ID: 1, Type: "info", Payload: []byte("Hello"), Timestamp: 1234567890},
		{ID: 2, Type: "warn", Payload: []byte("Warning"), Timestamp: 1234567891},
		{ID: 3, Type: "error", Payload: []byte("Error"), Timestamp: 1234567892},
	}

	var buf bytes.Buffer

	// Create encoder
	encoder := beve.NewEncoder(&buf)

	// Encode multiple messages
	for _, msg := range messages {
		_, err := encoder.Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Encoded %d messages in %d bytes\n", len(messages), buf.Len())

	// Create decoder
	decoder := beve.NewDecoder(&buf)

	// Decode multiple messages
	var decoded []Message
	for i := 0; i < len(messages); i++ {
		var msg Message
		if err := decoder.Decode(&msg); err != nil {
			log.Fatal(err)
		}
		decoded = append(decoded, msg)
	}

	fmt.Printf("Decoded %d messages:\n", len(decoded))
	for _, msg := range decoded {
		fmt.Printf("  [%d] %s: %s (timestamp: %d)\n",
			msg.ID, msg.Type, string(msg.Payload), msg.Timestamp)
	}
}
