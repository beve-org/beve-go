package main

import (
	"encoding/hex"
	"fmt"

	beve "github.com/beve-org/beve-go"
)

func main() {
	data := map[string]interface{}{
		"id":     int64(123),
		"name":   "Alice",
		"active": true,
		"metadata": map[string]interface{}{
			"created": "2025-10-12",
			"tags":    []interface{}{"admin", "user"},
		},
	}

	encoded, err := beve.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("encoded (%d bytes): %s\n", len(encoded), hex.EncodeToString(encoded))

	var raw map[string]beve.RawMessage
	if err := beve.Unmarshal(encoded, &raw); err != nil {
		fmt.Printf("unmarshal raw failed: %v\n", err)
		return
	}

	fmt.Printf("decoded keys: %v\n", len(raw))
	for k, v := range raw {
		fmt.Printf("key=%s len=%d bytes\n", k, len(v))
	}

	failingHex := "03101861637469766518206d6574616461746103081c637265617465640228323032352d31302d313210746167738508021461646d696e021075736572086964097b106e616d650214416c696365"
	failingPayload, err := hex.DecodeString(failingHex)
	if err != nil {
		panic(err)
	}
	fmt.Printf("failing payload len=%d\n", len(failingPayload))
	var rawFail map[string]beve.RawMessage
	if err := beve.Unmarshal(failingPayload, &rawFail); err != nil {
		fmt.Printf("decoding failing payload error: %v\n", err)
	} else {
		fmt.Printf("failing payload decoded keys=%d\n", len(rawFail))
	}
}
