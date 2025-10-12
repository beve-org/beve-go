package beve

import (
	"testing"
)

// TestStructTag_Default tests default "beve" tag behavior
func TestStructTag_Default(t *testing.T) {
	type Person struct {
		Name  string `beve:"name"`
		Age   int    `beve:"age"`
		Email string `beve:"email,omitempty"`
	}

	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	
	data, err := Marshal(p)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result Person
	err = Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Name != p.Name || result.Age != p.Age || result.Email != p.Email {
		t.Errorf("Data mismatch: got %+v, want %+v", result, p)
	}
}

// TestStructTag_JSONTag tests using json tags with SetStructTag
func TestStructTag_JSONTag(t *testing.T) {
	// Save original tag
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	// Set to use json tags
	SetStructTag("json")

	type Product struct {
		ID    int     `json:"id"`
		Name  string  `json:"name"`
		Price float64 `json:"price,omitempty"`
	}

	p := Product{ID: 123, Name: "Widget", Price: 19.99}
	
	data, err := Marshal(p)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result Product
	err = Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.ID != p.ID || result.Name != p.Name || result.Price != p.Price {
		t.Errorf("Data mismatch: got %+v, want %+v", result, p)
	}
}

// TestStructTag_CustomTag tests custom tag names
func TestStructTag_CustomTag(t *testing.T) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	SetStructTag("msgpack")

	type Message struct {
		ID      int    `msgpack:"id"`
		Content string `msgpack:"content"`
		Sender  string `msgpack:"sender,omitempty"`
	}

	m := Message{ID: 456, Content: "Hello", Sender: "Bob"}
	
	data, err := Marshal(m)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result Message
	err = Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.ID != m.ID || result.Content != m.Content || result.Sender != m.Sender {
		t.Errorf("Data mismatch: got %+v, want %+v", result, m)
	}
}

// TestStructTag_FallbackToJSON tests fallback behavior
func TestStructTag_FallbackToJSON(t *testing.T) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	// Set to use "proto" tag, but struct only has "json" tags
	SetStructTag("proto")

	type User struct {
		ID       int    `json:"id"`        // proto tag not present, should fallback to json
		Username string `json:"username"`
		Active   bool   `json:"active,omitempty"`
	}

	u := User{ID: 789, Username: "john_doe", Active: true}
	
	data, err := Marshal(u)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result User
	err = Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.ID != u.ID || result.Username != u.Username || result.Active != u.Active {
		t.Errorf("Data mismatch: got %+v, want %+v", result, u)
	}
}

// TestStructTag_OmitEmpty tests omitempty option with different tags
func TestStructTag_OmitEmpty(t *testing.T) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	SetStructTag("json")

	type Config struct {
		Required string `json:"required"`
		Optional string `json:"optional,omitempty"`
	}

	// Test with empty optional field
	c1 := Config{Required: "value"}
	data1, err := Marshal(c1)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result1 Config
	err = Unmarshal(data1, &result1)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result1.Required != c1.Required {
		t.Errorf("Required field mismatch: got %q, want %q", result1.Required, c1.Required)
	}

	// Test with non-empty optional field
	c2 := Config{Required: "value", Optional: "optional"}
	data2, err := Marshal(c2)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result2 Config
	err = Unmarshal(data2, &result2)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result2.Required != c2.Required || result2.Optional != c2.Optional {
		t.Errorf("Data mismatch: got %+v, want %+v", result2, c2)
	}
}

// TestStructTag_FieldNameMapping tests field name remapping
func TestStructTag_FieldNameMapping(t *testing.T) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	SetStructTag("json")

	type APIResponse struct {
		StatusCode int    `json:"status_code"` // Snake case
		Message    string `json:"message"`
		Data       string `json:"data"`
	}

	resp := APIResponse{StatusCode: 200, Message: "OK", Data: "result"}
	
	data, err := Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result APIResponse
	err = Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.StatusCode != resp.StatusCode || result.Message != resp.Message || result.Data != resp.Data {
		t.Errorf("Data mismatch: got %+v, want %+v", result, resp)
	}
}

// TestStructTag_SkipField tests "-" tag to skip fields
func TestStructTag_SkipField(t *testing.T) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	SetStructTag("json")

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"-"` // Should be skipped
		Token    string `json:"token"`
	}

	cred := Credentials{Username: "admin", Password: "secret123", Token: "abc"}
	
	data, err := Marshal(cred)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result Credentials
	err = Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Username != cred.Username || result.Token != cred.Token {
		t.Errorf("Data mismatch: got %+v, want %+v", result, cred)
	}

	// Password should not be encoded/decoded
	if result.Password != "" {
		t.Errorf("Password should be empty, got %q", result.Password)
	}
}

// TestStructTag_RuntimeChange tests changing tag at runtime
func TestStructTag_RuntimeChange(t *testing.T) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	type MultiTagStruct struct {
		Field1 string `beve:"beve_field1" json:"json_field1"`
		Field2 int    `beve:"beve_field2" json:"json_field2"`
	}

	m := MultiTagStruct{Field1: "value1", Field2: 42}

	// Test with beve tag
	SetStructTag("beve")
	dataBeve, err := Marshal(m)
	if err != nil {
		t.Fatalf("Marshal with beve tag failed: %v", err)
	}

	var resultBeve MultiTagStruct
	err = Unmarshal(dataBeve, &resultBeve)
	if err != nil {
		t.Fatalf("Unmarshal with beve tag failed: %v", err)
	}

	if resultBeve.Field1 != m.Field1 || resultBeve.Field2 != m.Field2 {
		t.Errorf("Data mismatch with beve tag: got %+v, want %+v", resultBeve, m)
	}

	// Test with json tag
	SetStructTag("json")
	dataJSON, err := Marshal(m)
	if err != nil {
		t.Fatalf("Marshal with json tag failed: %v", err)
	}

	var resultJSON MultiTagStruct
	err = Unmarshal(dataJSON, &resultJSON)
	if err != nil {
		t.Fatalf("Unmarshal with json tag failed: %v", err)
	}

	if resultJSON.Field1 != m.Field1 || resultJSON.Field2 != m.Field2 {
		t.Errorf("Data mismatch with json tag: got %+v, want %+v", resultJSON, m)
	}
}

// TestStructTag_NestedStructs tests nested structures with different tags
func TestStructTag_NestedStructs(t *testing.T) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	SetStructTag("json")

	type Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
		Zip    string `json:"zip,omitempty"`
	}

	type Person struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address Address `json:"address"`
	}

	p := Person{
		Name: "Charlie",
		Age:  35,
		Address: Address{
			Street: "123 Main St",
			City:   "Boston",
			Zip:    "02101",
		},
	}

	data, err := Marshal(p)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result Person
	err = Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Name != p.Name || result.Age != p.Age {
		t.Errorf("Person data mismatch: got %+v, want %+v", result, p)
	}

	if result.Address.Street != p.Address.Street || 
	   result.Address.City != p.Address.City || 
	   result.Address.Zip != p.Address.Zip {
		t.Errorf("Address data mismatch: got %+v, want %+v", result.Address, p.Address)
	}
}

// TestStructTag_EmptyTagName tests empty tag name defaults to "beve"
func TestStructTag_EmptyTagName(t *testing.T) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)

	// Set empty tag (should default to "beve")
	SetStructTag("")

	currentTag := GetStructTag()
	if currentTag != "beve" {
		t.Errorf("Expected default tag 'beve', got %q", currentTag)
	}

	type Simple struct {
		Value string `beve:"value"`
	}

	s := Simple{Value: "test"}
	data, err := Marshal(s)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result Simple
	err = Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Value != s.Value {
		t.Errorf("Value mismatch: got %q, want %q", result.Value, s.Value)
	}
}

// BenchmarkStructTag_BeveTag benchmarks with default beve tag
func BenchmarkStructTag_BeveTag(b *testing.B) {
	type Data struct {
		ID   int    `beve:"id"`
		Name string `beve:"name"`
		Age  int    `beve:"age"`
	}

	d := Data{ID: 123, Name: "Benchmark", Age: 30}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data, _ := Marshal(d)
		var result Data
		_ = Unmarshal(data, &result)
	}
}

// BenchmarkStructTag_JSONTag benchmarks with json tag
func BenchmarkStructTag_JSONTag(b *testing.B) {
	originalTag := GetStructTag()
	defer SetStructTag(originalTag)
	SetStructTag("json")

	type Data struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	d := Data{ID: 123, Name: "Benchmark", Age: 30}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data, _ := Marshal(d)
		var result Data
		_ = Unmarshal(data, &result)
	}
}
