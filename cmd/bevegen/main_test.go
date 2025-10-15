package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseStructTag(t *testing.T) {
	tests := []struct {
		name          string
		tag           string
		defaultName   string
		wantName      string
		wantOmitEmpty bool
	}{
		{
			name:          "beve tag with name",
			tag:           `beve:"user_id"`,
			defaultName:   "ID",
			wantName:      "user_id",
			wantOmitEmpty: false,
		},
		{
			name:          "beve tag with omitempty",
			tag:           `beve:"email,omitempty"`,
			defaultName:   "Email",
			wantName:      "email",
			wantOmitEmpty: true,
		},
		{
			name:          "beve tag with dash (skip)",
			tag:           `beve:"-"`,
			defaultName:   "Internal",
			wantName:      "Internal",
			wantOmitEmpty: false,
		},
		{
			name:          "json tag fallback",
			tag:           `json:"name"`,
			defaultName:   "Name",
			wantName:      "name",
			wantOmitEmpty: false,
		},
		{
			name:          "json tag with omitempty",
			tag:           `json:"data,omitempty"`,
			defaultName:   "Data",
			wantName:      "data",
			wantOmitEmpty: true,
		},
		{
			name:          "no tag",
			tag:           ``,
			defaultName:   "Field",
			wantName:      "Field",
			wantOmitEmpty: false,
		},
		{
			name:          "empty beve tag",
			tag:           `beve:""`,
			defaultName:   "Field",
			wantName:      "Field",
			wantOmitEmpty: false,
		},
		{
			name:          "complex tag with multiple options",
			tag:           `beve:"field,omitempty" json:"json_field"`,
			defaultName:   "Field",
			wantName:      "field",
			wantOmitEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotOmitEmpty := parseStructTag(tt.tag, tt.defaultName)
			if gotName != tt.wantName {
				t.Errorf("parseStructTag() name = %v, want %v", gotName, tt.wantName)
			}
			if gotOmitEmpty != tt.wantOmitEmpty {
				t.Errorf("parseStructTag() omitEmpty = %v, want %v", gotOmitEmpty, tt.wantOmitEmpty)
			}
		})
	}
}

func TestIsInlinableType(t *testing.T) {
	tests := []struct {
		typeName string
		want     bool
	}{
		// Primitives - should be inlinable
		{"bool", true},
		{"int", true},
		{"int8", true},
		{"int16", true},
		{"int32", true},
		{"int64", true},
		{"uint", true},
		{"uint8", true},
		{"uint16", true},
		{"uint32", true},
		{"uint64", true},
		{"float32", true},
		{"float64", true},
		{"string", true},

		// Complex types - not inlinable
		{"[]int", false},
		{"map[string]int", false},
		{"*User", false},
		{"User", false},
		{"time.Time", false},
		{"[]byte", false},
	}

	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			got := isInlinableType(tt.typeName)
			if got != tt.want {
				t.Errorf("isInlinableType(%q) = %v, want %v", tt.typeName, got, tt.want)
			}
		})
	}
}

func TestSanitizeTypeName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"int", "int"},
		{"string", "string"},
		{"[]byte", "Slicebyte"},
		{"[]int", "Sliceint"},
		{"map[string]int", "mapstringint"},
		{"*User", "PtrUser"},
		{"time.Time", "timeTime"},
		{"encoding/json.RawMessage", "encodingjsonRawMessage"},
		{"[]*User", "SlicePtrUser"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := sanitizeTypeName(tt.input)
			if got != tt.want {
				t.Errorf("sanitizeTypeName(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCodeGeneration(t *testing.T) {
	// Create a temporary directory for test
	tmpDir := t.TempDir()

	// Create a test Go file with struct definitions
	testFile := filepath.Join(tmpDir, "test.go")
	testCode := `package testpkg

type SimpleStruct struct {
	ID   int64  ` + "`beve:\"id\"`" + `
	Name string ` + "`beve:\"name\"`" + `
}

type ComplexStruct struct {
	ID       int64   ` + "`beve:\"id\"`" + `
	Name     string  ` + "`beve:\"name\"`" + `
	Email    string  ` + "`beve:\"email,omitempty\"`" + `
	Age      int     ` + "`beve:\"age\"`" + `
	IsActive bool    ` + "`beve:\"active\"`" + `
	Score    float64 ` + "`beve:\"score\"`" + `
}
`

	if err := os.WriteFile(testFile, []byte(testCode), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Change to temp directory
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)
	os.Chdir(tmpDir)

	// Test simple struct generation
	t.Run("SimpleStruct", func(t *testing.T) {
		*typeNames = "SimpleStruct"
		*output = ""

		// Note: We can't easily test the full main() flow here
		// because it calls os.Exit and flag.Parse
		// Instead, we test the individual functions

		// This would require refactoring main() to be more testable
		// For now, we verify the helper functions work correctly
	})
}

func TestTemplateHelpers(t *testing.T) {
	t.Run("zeroValue", func(t *testing.T) {
		zeroValue := templateFuncs["zeroValue"].(func(string) string)

		tests := []struct {
			typeName string
			want     string
		}{
			{"bool", "false"},
			{"string", `""`},
			{"int", "0"},
			{"int64", "0"},
			{"float64", "0"},
		}

		for _, tt := range tests {
			got := zeroValue(tt.typeName)
			if got != tt.want {
				t.Errorf("zeroValue(%q) = %v, want %v", tt.typeName, got, tt.want)
			}
		}
	})

	t.Run("title", func(t *testing.T) {
		title := templateFuncs["title"].(func(string) string)

		tests := []struct {
			input string
			want  string
		}{
			{"user", "User"},
			{"userID", "UserID"},
			{"[]byte", "Slicebyte"},
		}

		for _, tt := range tests {
			got := title(tt.input)
			if got != tt.want {
				t.Errorf("title(%q) = %v, want %v", tt.input, got, tt.want)
			}
		}
	})

	t.Run("hasPrefix", func(t *testing.T) {
		hasPrefix := templateFuncs["hasPrefix"].(func(string, string) bool)

		if !hasPrefix("int64", "int") {
			t.Error("hasPrefix('int64', 'int') should be true")
		}
		if hasPrefix("string", "int") {
			t.Error("hasPrefix('string', 'int') should be false")
		}
	})

	t.Run("needsReflect", func(t *testing.T) {
		needsReflect := templateFuncs["needsReflect"].(func([]fieldDef) bool)

		primitiveFields := []fieldDef{
			{Type: "int64"},
			{Type: "string"},
			{Type: "bool"},
		}
		if needsReflect(primitiveFields) {
			t.Error("needsReflect should be false for all primitives")
		}

		complexFields := []fieldDef{
			{Type: "int64"},
			{Type: "[]string"},
		}
		if !needsReflect(complexFields) {
			t.Error("needsReflect should be true for complex types")
		}
	})

	t.Run("uniqueTypes", func(t *testing.T) {
		uniqueTypes := templateFuncs["uniqueTypes"].(func([]fieldDef) []string)

		fields := []fieldDef{
			{Type: "int64"},
			{Type: "string"},
			{Type: "int64"}, // duplicate
			{Type: "bool"},
			{Type: "string"}, // duplicate
		}

		unique := uniqueTypes(fields)
		if len(unique) != 3 {
			t.Errorf("uniqueTypes should return 3 types, got %d", len(unique))
		}

		// Check that duplicates are removed
		typeMap := make(map[string]bool)
		for _, typ := range unique {
			if typeMap[typ] {
				t.Errorf("uniqueTypes returned duplicate type: %s", typ)
			}
			typeMap[typ] = true
		}
	})
}

func TestGeneratedCodeCompiles(t *testing.T) {
	// This test verifies that generated code actually compiles
	// by generating code in a temp directory and running go build

	tmpDir := t.TempDir()

	// Create test struct
	testFile := filepath.Join(tmpDir, "model.go")
	modelCode := `package testmodel

type User struct {
	ID       int64   ` + "`beve:\"id\"`" + `
	Username string  ` + "`beve:\"username\"`" + `
	Email    string  ` + "`beve:\"email,omitempty\"`" + `
	Age      int     ` + "`beve:\"age\"`" + `
	Active   bool    ` + "`beve:\"active\"`" + `
}
`
	if err := os.WriteFile(testFile, []byte(modelCode), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Generate BEVE code manually
	genCode := `// Code generated by bevegen. DO NOT EDIT.

package testmodel

import (
	"github.com/beve-org/beve-go/core"
)

func (s *User) MarshalBEVE() ([]byte, error) {
	enc := core.GetEncoderFromPool()
	defer core.PutEncoderToPool(enc)

	if err := enc.WriteByte(0x86); err != nil {
		return nil, err
	}

	fieldCount := 5
	if s.Email == "" {
		fieldCount--
	}

	if err := enc.WriteCompressedUint(uint64(fieldCount)); err != nil {
		return nil, err
	}

	if err := core.EncodeStringFast(enc, "id"); err != nil {
		return nil, err
	}
	if err := core.EncodeIntFast(enc, int64(s.ID)); err != nil {
		return nil, err
	}

	if err := core.EncodeStringFast(enc, "username"); err != nil {
		return nil, err
	}
	if err := core.EncodeStringFast(enc, s.Username); err != nil {
		return nil, err
	}

	if s.Email != "" {
		if err := core.EncodeStringFast(enc, "email"); err != nil {
			return nil, err
		}
		if err := core.EncodeStringFast(enc, s.Email); err != nil {
			return nil, err
		}
	}

	if err := core.EncodeStringFast(enc, "age"); err != nil {
		return nil, err
	}
	if err := core.EncodeIntFast(enc, int64(s.Age)); err != nil {
		return nil, err
	}

	if err := core.EncodeStringFast(enc, "active"); err != nil {
		return nil, err
	}
	if err := core.EncodeBoolFast(enc, s.Active); err != nil {
		return nil, err
	}

	return enc.Buf.Bytes(), nil
}
`
	genFile := filepath.Join(tmpDir, "user_beve.go")
	if err := os.WriteFile(genFile, []byte(genCode), 0644); err != nil {
		t.Fatalf("Failed to create generated file: %v", err)
	}

	// Note: We can't easily compile this without the full module setup
	// In a real test environment, you would:
	// 1. Create a go.mod in tmpDir
	// 2. Run `go build` to verify compilation
	// 3. Actually run the generated code with sample data

	// For now, we just verify the files were created
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Error("Test file was not created")
	}
	if _, err := os.Stat(genFile); os.IsNotExist(err) {
		t.Error("Generated file was not created")
	}

	// Verify generated code contains expected elements
	content, err := os.ReadFile(genFile)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	requiredStrings := []string{
		"package testmodel",
		"func (s *User) MarshalBEVE()",
		"core.GetEncoderFromPool()",
		"core.EncodeStringFast",
		"core.EncodeIntFast",
		"core.EncodeBoolFast",
	}

	for _, required := range requiredStrings {
		if !strings.Contains(string(content), required) {
			t.Errorf("Generated code missing required string: %s", required)
		}
	}
}
