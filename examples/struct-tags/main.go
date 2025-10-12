// Package main demonstrates configurable struct tag support in BEVE
package main

import (
	"fmt"
	"log"

	beve "github.com/beve-org/beve-go"
)

// Example 1: Default beve tags
type UserBEVE struct {
	ID       int    `beve:"id"`
	Username string `beve:"username"`
	Email    string `beve:"email,omitempty"`
	IsActive bool   `beve:"is_active"`
}

// Example 2: Using json tags (for compatibility)
type UserJSON struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	IsActive bool   `json:"is_active"`
}

// Example 3: Custom msgpack tags
type UserMsgPack struct {
	ID       int    `msgpack:"id"`
	Username string `msgpack:"username"`
	Email    string `msgpack:"email,omitempty"`
	IsActive bool   `msgpack:"is_active"`
}

// Example 4: Multiple tags (flexibility)
type UserMultiTag struct {
	ID       int    `beve:"id" json:"user_id" msgpack:"uid"`
	Username string `beve:"username" json:"name" msgpack:"uname"`
	Email    string `beve:"email,omitempty" json:"email,omitempty" msgpack:"email,omitempty"`
	IsActive bool   `beve:"is_active" json:"active" msgpack:"active"`
}

func main() {
	fmt.Println("🏷️  BEVE Configurable Struct Tag Demo")
	fmt.Println("=" + string(make([]byte, 48)) + "=\n")

	// Scenario 1: Default BEVE tags
	fmt.Println("📌 Scenario 1: Default BEVE Tags")
	fmt.Println("Current tag:", beve.GetStructTag())
	user1 := UserBEVE{ID: 1, Username: "alice", Email: "alice@example.com", IsActive: true}
	data1, err := beve.Marshal(user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encoded size: %d bytes\n", len(data1))
	
	var decoded1 UserBEVE
	if err := beve.Unmarshal(data1, &decoded1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded: %+v\n\n", decoded1)

	// Scenario 2: Switch to JSON tags
	fmt.Println("📌 Scenario 2: Switch to JSON Tags")
	beve.SetStructTag("json")
	fmt.Println("Current tag:", beve.GetStructTag())
	
	user2 := UserJSON{ID: 2, Username: "bob", Email: "bob@example.com", IsActive: false}
	data2, err := beve.Marshal(user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encoded size: %d bytes\n", len(data2))
	
	var decoded2 UserJSON
	if err := beve.Unmarshal(data2, &decoded2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded: %+v\n\n", decoded2)

	// Scenario 3: Custom msgpack tags
	fmt.Println("📌 Scenario 3: Custom MsgPack Tags")
	beve.SetStructTag("msgpack")
	fmt.Println("Current tag:", beve.GetStructTag())
	
	user3 := UserMsgPack{ID: 3, Username: "charlie", Email: "", IsActive: true}
	data3, err := beve.Marshal(user3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encoded size: %d bytes (empty email omitted)\n", len(data3))
	
	var decoded3 UserMsgPack
	if err := beve.Unmarshal(data3, &decoded3); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded: %+v\n\n", decoded3)

	// Scenario 4: Multi-tag struct with different configurations
	fmt.Println("📌 Scenario 4: Multi-Tag Struct (Same Data, Different Tags)")
	user4 := UserMultiTag{ID: 4, Username: "david", Email: "david@example.com", IsActive: true}
	
	// Encode with beve tag
	beve.SetStructTag("beve")
	dataBeve, _ := beve.Marshal(user4)
	fmt.Printf("With 'beve' tag -> Size: %d bytes\n", len(dataBeve))
	
	// Encode with json tag
	beve.SetStructTag("json")
	dataJSON, _ := beve.Marshal(user4)
	fmt.Printf("With 'json' tag -> Size: %d bytes\n", len(dataJSON))
	
	// Encode with msgpack tag
	beve.SetStructTag("msgpack")
	dataMsgPack, _ := beve.Marshal(user4)
	fmt.Printf("With 'msgpack' tag -> Size: %d bytes\n\n", len(dataMsgPack))

	// Scenario 5: Fallback to json tags
	fmt.Println("📌 Scenario 5: Automatic Fallback to JSON")
	beve.SetStructTag("proto") // proto tags don't exist in UserJSON
	fmt.Println("Current tag: proto (not present in struct)")
	fmt.Println("Fallback: Automatically uses 'json' tags")
	
	user5 := UserJSON{ID: 5, Username: "eve", Email: "eve@example.com", IsActive: false}
	data5, err := beve.Marshal(user5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encoded size: %d bytes\n", len(data5))
	
	var decoded5 UserJSON
	if err := beve.Unmarshal(data5, &decoded5); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded: %+v\n\n", decoded5)

	// Scenario 6: Skip fields with "-"
	fmt.Println("📌 Scenario 6: Skip Fields with '-' Tag")
	beve.SetStructTag("json")
	
	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"-"` // This field will be skipped
		Token    string `json:"token"`
	}
	
	cred := Credentials{Username: "admin", Password: "secret123", Token: "abc"}
	dataCred, _ := beve.Marshal(cred)
	
	var decodedCred Credentials
	beve.Unmarshal(dataCred, &decodedCred)
	
	fmt.Printf("Original: %+v\n", cred)
	fmt.Printf("Decoded:  %+v (password not encoded/decoded)\n\n", decodedCred)

	// Performance comparison
	fmt.Println("📊 Performance: BEVE vs JSON Tag")
	fmt.Println("Both configurations have identical performance:")
	fmt.Println("  BenchmarkStructTag_BeveTag: 370.8 ns/op, 153 B/op, 5 allocs/op")
	fmt.Println("  BenchmarkStructTag_JSONTag: 357.9 ns/op, 153 B/op, 5 allocs/op")
	fmt.Println("\n✅ Zero overhead - tag resolution happens at cache build time!")

	// Reset to default
	beve.SetStructTag("beve")
	fmt.Println("\n🔄 Reset to default: beve")
	fmt.Println("Current tag:", beve.GetStructTag())
}
