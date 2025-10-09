package beve

import (
	"testing"
)

func TestInlineStruct(t *testing.T) {
	type Address struct {
		City    string `beve:"city"`
		Country string `beve:"country"`
	}

	type Person struct {
		Name    string  `beve:"name"`
		Address Address `beve:",inline"`
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
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded Person
	err = Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.Name != p.Name {
		t.Errorf("Name mismatch: got %q, want %q", decoded.Name, p.Name)
	}
	if decoded.Address.City != p.Address.City {
		t.Errorf("City mismatch: got %q, want %q", decoded.Address.City, p.Address.City)
	}
	if decoded.Address.Country != p.Address.Country {
		t.Errorf("Country mismatch: got %q, want %q", decoded.Address.Country, p.Address.Country)
	}
	if decoded.Age != p.Age {
		t.Errorf("Age mismatch: got %d, want %d", decoded.Age, p.Age)
	}
}

func TestAnonymousStruct(t *testing.T) {
	type Base struct {
		ID   int    `beve:"id"`
		Name string `beve:"name"`
	}

	type Extended struct {
		Base
		Extra string `beve:"extra"`
	}

	e := Extended{
		Base: Base{
			ID:   123,
			Name: "Test",
		},
		Extra: "Additional",
	}

	data, err := Marshal(e)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded Extended
	err = Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.ID != e.ID {
		t.Errorf("ID mismatch: got %d, want %d", decoded.ID, e.ID)
	}
	if decoded.Name != e.Name {
		t.Errorf("Name mismatch: got %q, want %q", decoded.Name, e.Name)
	}
	if decoded.Extra != e.Extra {
		t.Errorf("Extra mismatch: got %q, want %q", decoded.Extra, e.Extra)
	}
}

func TestNestedInline(t *testing.T) {
	type Level3 struct {
		Z string `beve:"z"`
	}

	type Level2 struct {
		Y      string `beve:"y"`
		Level3 `beve:",inline"`
	}

	type Level1 struct {
		X      string `beve:"x"`
		Level2 `beve:",inline"`
	}

	l := Level1{
		X: "x-value",
		Level2: Level2{
			Y: "y-value",
			Level3: Level3{
				Z: "z-value",
			},
		},
	}

	data, err := Marshal(l)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded Level1
	err = Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.X != l.X {
		t.Errorf("X mismatch: got %q, want %q", decoded.X, l.X)
	}
	if decoded.Y != l.Y {
		t.Errorf("Y mismatch: got %q, want %q", decoded.Y, l.Y)
	}
	if decoded.Z != l.Z {
		t.Errorf("Z mismatch: got %q, want %q", decoded.Z, l.Z)
	}
}

func TestInlineWithOmitEmpty(t *testing.T) {
	type Optional struct {
		Field1 string `beve:"field1,omitempty"`
		Field2 int    `beve:"field2,omitempty"`
	}

	type Container struct {
		Name     string   `beve:"name"`
		Optional Optional `beve:",inline"`
	}

	// Test with values
	c1 := Container{
		Name: "Test",
		Optional: Optional{
			Field1: "value1",
			Field2: 42,
		},
	}

	data1, err := Marshal(c1)
	if err != nil {
		t.Fatalf("Marshal c1 failed: %v", err)
	}

	var decoded1 Container
	err = Unmarshal(data1, &decoded1)
	if err != nil {
		t.Fatalf("Unmarshal c1 failed: %v", err)
	}

	if decoded1.Name != c1.Name || decoded1.Optional.Field1 != c1.Optional.Field1 || decoded1.Optional.Field2 != c1.Optional.Field2 {
		t.Errorf("c1 mismatch: got %+v, want %+v", decoded1, c1)
	}

	// Test with empty values
	c2 := Container{
		Name: "Test2",
	}

	data2, err := Marshal(c2)
	if err != nil {
		t.Fatalf("Marshal c2 failed: %v", err)
	}

	var decoded2 Container
	err = Unmarshal(data2, &decoded2)
	if err != nil {
		t.Fatalf("Unmarshal c2 failed: %v", err)
	}

	if decoded2.Name != c2.Name {
		t.Errorf("c2 Name mismatch: got %q, want %q", decoded2.Name, c2.Name)
	}
}
