package translatornative

import (
	"unsafe"
)

// unsafeString creates a zero-copy string from a byte slice.
// WARNING: The byte slice must not be modified after this call.
//
//go:inline
func unsafeString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}

// unsafeBytes creates a zero-copy byte slice from a string.
// WARNING: The string must not be modified (which is safe since strings are immutable).
//
//go:inline
func unsafeBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// Common JSON key intern pool (top 50 most common keys)
var keyInternPool = map[string]string{
	"id":          "id",
	"name":        "name",
	"email":       "email",
	"age":         "age",
	"created":     "created",
	"updated":     "updated",
	"active":      "active",
	"status":      "status",
	"type":        "type",
	"value":       "value",
	"data":        "data",
	"user":        "user",
	"username":    "username",
	"password":    "password",
	"token":       "token",
	"key":         "key",
	"code":        "code",
	"message":     "message",
	"error":       "error",
	"success":     "success",
	"total":       "total",
	"count":       "count",
	"page":        "page",
	"size":        "size",
	"limit":       "limit",
	"offset":      "offset",
	"sort":        "sort",
	"order":       "order",
	"filter":      "filter",
	"search":      "search",
	"query":       "query",
	"result":      "result",
	"results":     "results",
	"items":       "items",
	"list":        "list",
	"timestamp":   "timestamp",
	"date":        "date",
	"time":        "time",
	"start":       "start",
	"end":         "end",
	"duration":    "duration",
	"title":       "title",
	"description": "description",
	"url":         "url",
	"path":        "path",
	"method":      "method",
	"headers":     "headers",
	"body":        "body",
	"response":    "response",
	"request":     "request",
	"version":     "version",
}

// internString returns an interned version of the string if it's common,
// otherwise returns the original string.
// Optimized: Switch on length + first char (no map lookup!)
//
//go:inline
func internString(s string) string {
	// Fast path: check top 30 most common keys by length
	switch len(s) {
	case 2:
		if s == "id" {
			return "id"
		}
	case 3:
		switch s[0] {
		case 'a':
			if s == "age" {
				return "age"
			}
		case 'u':
			if s == "url" {
				return "url"
			}
		case 'k':
			if s == "key" {
				return "key"
			}
		case 'e':
			if s == "end" {
				return "end"
			}
		}
	case 4:
		switch s[0] {
		case 'n':
			if s == "name" {
				return "name"
			}
		case 'd':
			if s == "date" {
				return "date"
			}
			if s == "data" {
				return "data"
			}
		case 't':
			if s == "type" {
				return "type"
			}
			if s == "time" {
				return "time"
			}
		case 'c':
			if s == "code" {
				return "code"
			}
		case 'p':
			if s == "page" {
				return "page"
			}
			if s == "path" {
				return "path"
			}
		case 's':
			if s == "size" {
				return "size"
			}
			if s == "sort" {
				return "sort"
			}
		case 'b':
			if s == "body" {
				return "body"
			}
		case 'u':
			if s == "user" {
				return "user"
			}
		}
	case 5:
		switch s[0] {
		case 'e':
			if s == "email" {
				return "email"
			}
			if s == "error" {
				return "error"
			}
		case 't':
			if s == "title" {
				return "title"
			}
			if s == "token" {
				return "token"
			}
			if s == "total" {
				return "total"
			}
		case 'v':
			if s == "value" {
				return "value"
			}
		case 's':
			if s == "start" {
				return "start"
			}
		case 'c':
			if s == "count" {
				return "count"
			}
		case 'l':
			if s == "limit" {
				return "limit"
			}
		case 'o':
			if s == "order" {
				return "order"
			}
		case 'q':
			if s == "query" {
				return "query"
			}
		case 'i':
			if s == "items" {
				return "items"
			}
		}
	case 6:
		switch s[0] {
		case 'a':
			if s == "active" {
				return "active"
			}
		case 's':
			if s == "status" {
				return "status"
			}
			if s == "search" {
				return "search"
			}
		case 'f':
			if s == "filter" {
				return "filter"
			}
		case 'o':
			if s == "offset" {
				return "offset"
			}
		case 'r':
			if s == "result" {
				return "result"
			}
		case 'm':
			if s == "method" {
				return "method"
			}
		}
	case 7:
		switch s[0] {
		case 'c':
			if s == "created" {
				return "created"
			}
		case 'u':
			if s == "updated" {
				return "updated"
			}
		case 'm':
			if s == "message" {
				return "message"
			}
		case 'r':
			if s == "results" {
				return "results"
			}
			if s == "request" {
				return "request"
			}
		case 's':
			if s == "success" {
				return "success"
			}
		case 'v':
			if s == "version" {
				return "version"
			}
		}
	case 8:
		switch s[0] {
		case 'u':
			if s == "username" {
				return "username"
			}
		case 'p':
			if s == "password" {
				return "password"
			}
		case 'd':
			if s == "duration" {
				return "duration"
			}
		case 'r':
			if s == "response" {
				return "response"
			}
		}
	}
	return s
}

// Fast inline helpers
//
//go:inline
func isHexDigit(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// needsEscape checks if a string needs JSON escaping (fast path check).
//
//go:inline
func needsEscape(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 0x20 || c == '"' || c == '\\' {
			return true
		}
	}
	return false
}

// estimateJSONSize estimates the output size for buffer pre-allocation.
func estimateJSONSize(v interface{}) int {
	switch val := v.(type) {
	case nil:
		return 4 // "null"
	case bool:
		return 5 // "false" worst case
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return 20 // max digits for int64
	case float32, float64:
		return 24 // max float representation
	case string:
		return len(val) + 2 + len(val)/10 // string + quotes + ~10% escapes
	case []interface{}:
		size := 2 // []
		for _, item := range val {
			size += estimateJSONSize(item) + 1 // item + comma
		}
		return size
	case map[string]interface{}:
		size := 2 // {}
		for key, value := range val {
			size += len(key) + 3 + estimateJSONSize(value) + 1 // "key": value,
		}
		return size
	default:
		return 64 // fallback
	}
}
