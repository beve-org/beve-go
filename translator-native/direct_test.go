package translatornative

import (
	"testing"
)

func TestDirectEncoder_MediumJSON(t *testing.T) {
	json := []byte(`{
		"users": [
			{"id":1,"name":"Alice","age":30,"email":"alice@example.com","active":true},
			{"id":2,"name":"Bob","age":25,"email":"bob@example.com","active":false},
			{"id":3,"name":"Charlie","age":35,"email":"charlie@example.com","active":true}
		],
		"total": 3,
		"page": 1
	}`)

	enc := NewDirectEncoder(json)
	result, err := enc.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v\nPosition: %d\nContext: %q", err, enc.pos, string(json[max(0, enc.pos-20):min(len(json), enc.pos+20)]))
	}

	t.Logf("Encoded %d bytes → %d bytes BEVE", len(json), len(result))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
