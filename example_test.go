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

    fmt.Printf("encoded bytes: %v\n", data)

    var decoded Person
    if err := Unmarshal(data, &decoded); err != nil {
        panic(err)
    }

    fmt.Printf("decoded: %+v\n", decoded)
    // Output:
    // encoded bytes: [3 4 16 8 97 100 97 8 8 0]
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
