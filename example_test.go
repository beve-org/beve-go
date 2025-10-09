package beve

import (
	"bytes"
	"fmt"
)

func ExampleMarshal() {
	type Person struct {
		Name string `beve:"name"`
		Age  int    `beve:"age,omitempty"`
	}

	data, err := Marshal(Person{Name: "Ada"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("encoded: %x\n", data)

	var decoded Person
	if err := Unmarshal(data, &decoded); err != nil {
		panic(err)
	}

	fmt.Printf("decoded: %+v\n", decoded)
	// Output:
	// encoded: 0304106e616d65020c416461
	// decoded: {Name:Ada Age:0}
}

func ExampleRawMessage() {
	raw := RawMessage{0x18}

	data, err := Marshal(raw)
	if err != nil {
		panic(err)
	}

	var decoded RawMessage
	if err := Unmarshal(data, &decoded); err != nil {
		panic(err)
	}

	fmt.Printf("raw equal: %v\n", bytes.Equal(data, decoded))
	// Output:
	// raw equal: true
}

func ExampleEncoder_streaming() {
	var buf bytes.Buffer
	enc := NewEncoder(&buf)
	if _, err := enc.Encode([]int32{1, 2, 3}); err != nil {
		panic(err)
	}

	dec := NewDecoder(bytes.NewReader(buf.Bytes()))
	var out []int32
	if err := dec.Decode(&out); err != nil {
		panic(err)
	}

	fmt.Printf("decoded: %v\n", out)
	// Output:
	// decoded: [1 2 3]
}

func ExampleMarshal_inlineStruct() {
	type Address struct {
		City    string `beve:"city"`
		Country string `beve:"country"`
	}

	type Person struct {
		Name    string  `beve:"name"`
		Address Address `beve:",inline"` // Fields will be flattened
		Age     int     `beve:"age"`
	}

	p := Person{
		Name: "Alice",
		Address: Address{
			City:    "Istanbul",
			Country: "Turkey",
		},
		Age: 30,
	}

	data, err := Marshal(p)
	if err != nil {
		panic(err)
	}

	var decoded Person
	if err := Unmarshal(data, &decoded); err != nil {
		panic(err)
	}

	fmt.Printf("%s lives in %s, %s\n", decoded.Name, decoded.Address.City, decoded.Address.Country)
	// Output: Alice lives in Istanbul, Turkey
}

func ExampleMarshal_anonymousStruct() {
	type Base struct {
		ID   int    `beve:"id"`
		Name string `beve:"name"`
	}

	type Extended struct {
		Base // Anonymous field, automatically inlined
		Role string `beve:"role"`
	}

	e := Extended{
		Base: Base{ID: 1, Name: "Admin"},
		Role: "Administrator",
	}

	data, err := Marshal(e)
	if err != nil {
		panic(err)
	}

	var decoded Extended
	if err := Unmarshal(data, &decoded); err != nil {
		panic(err)
	}

	fmt.Printf("ID: %d, Name: %s, Role: %s\n", decoded.ID, decoded.Name, decoded.Role)
	// Output: ID: 1, Name: Admin, Role: Administrator
}
