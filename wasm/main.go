//go:build wasm
// +build wasm

package main

import (
	"syscall/js"

	beve "github.com/beve-org/beve-go"
)

// BEVE WASM API
// Provides JavaScript bindings for BEVE serialization

// marshal encodes a JavaScript object to BEVE format
func marshal(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return map[string]interface{}{
			"error": "marshal requires exactly 1 argument",
		}
	}

	// Convert JS value to Go interface
	data := jsToGo(args[0])

	// Debug: print what we're trying to marshal
	println("[BEVE WASM] Marshal input type:", getTypeName(data))
	if m, ok := data.(map[string]interface{}); ok {
		println("[BEVE WASM] Map keys:", len(m))
		for k, v := range m {
			println("  -", k, ":", getTypeName(v))
		}
	}

	// Marshal with BEVE
	encoded, err := beve.Marshal(data)
	if err != nil {
		println("[BEVE WASM] Marshal ERROR:", err.Error())
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	println("[BEVE WASM] Marshal SUCCESS: encoded", len(encoded), "bytes")

	// Return as Uint8Array for JavaScript
	dst := js.Global().Get("Uint8Array").New(len(encoded))
	js.CopyBytesToJS(dst, encoded)

	return map[string]interface{}{
		"data": dst,
	}
}

// unmarshal decodes BEVE format to a JavaScript object
func unmarshal(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return map[string]interface{}{
			"error": "unmarshal requires exactly 1 argument",
		}
	}

	// Get byte array from JavaScript
	src := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(src, args[0])

	println("[BEVE WASM] Unmarshal: received", len(src), "bytes")

	// Use RawMessage to preserve the original bytes, then decode manually
	result, err := unmarshalDynamic(src)
	if err != nil {
		println("[BEVE WASM] Unmarshal ERROR:", err.Error())
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	println("[BEVE WASM] Unmarshal SUCCESS: type:", getTypeName(result))

	// Convert Go result to JS value
	jsResult := goToJs(result)

	return map[string]interface{}{
		"data": jsResult,
	}
}

// version returns BEVE library version
func version(this js.Value, args []js.Value) interface{} {
	return "1.2.0-wasm"
}

// benchmark runs a simple marshal/unmarshal benchmark
func benchmark(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return map[string]interface{}{
			"error": "benchmark requires 2 arguments: data, iterations",
		}
	}

	data := jsToGo(args[0])
	iterations := args[1].Int()

	// Marshal benchmark
	marshalStart := js.Global().Get("performance").Call("now")
	for i := 0; i < iterations; i++ {
		_, _ = beve.Marshal(data)
	}
	marshalEnd := js.Global().Get("performance").Call("now")
	marshalTime := marshalEnd.Float() - marshalStart.Float()

	// Unmarshal benchmark
	encoded, _ := beve.Marshal(data)
	unmarshalStart := js.Global().Get("performance").Call("now")
	for i := 0; i < iterations; i++ {
		var result interface{}
		_ = beve.Unmarshal(encoded, &result)
	}
	unmarshalEnd := js.Global().Get("performance").Call("now")
	unmarshalTime := unmarshalEnd.Float() - unmarshalStart.Float()

	return map[string]interface{}{
		"marshal": map[string]interface{}{
			"totalMs":   marshalTime,
			"avgMs":     marshalTime / float64(iterations),
			"opsPerSec": float64(iterations) / (marshalTime / 1000.0),
		},
		"unmarshal": map[string]interface{}{
			"totalMs":   unmarshalTime,
			"avgMs":     unmarshalTime / float64(iterations),
			"opsPerSec": float64(iterations) / (unmarshalTime / 1000.0),
		},
		"payloadSize": len(encoded),
	}
}

// unmarshalDynamic decodes BEVE bytes into a dynamic Go value
// Now that BEVE supports interface{} unmarshaling, we can do it directly!
func unmarshalDynamic(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, nil
	}

	// Direct unmarshal to interface{} - BEVE now supports this!
	var result interface{}
	if err := beve.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Helper: Get type name for debugging
func getTypeName(v interface{}) string {
	if v == nil {
		return "nil"
	}
	switch t := v.(type) {
	case bool:
		return "bool"
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case uint:
		return "uint"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case string:
		return "string"
	case []interface{}:
		return "[]interface{}"
	case map[string]interface{}:
		return "map[string]interface{}"
	case []byte:
		return "[]byte"
	default:
		_ = t // use t to avoid warning
		return "UNKNOWN TYPE"
	}
}

// Helper: Convert JS value to Go interface
func jsToGo(val js.Value) interface{} {
	switch val.Type() {
	case js.TypeNull, js.TypeUndefined:
		return nil
	case js.TypeBoolean:
		return val.Bool()
	case js.TypeNumber:
		// JavaScript numbers are float64, but check if it's an integer
		f := val.Float()
		if f == float64(int64(f)) {
			// It's a whole number, return as int64
			return int64(f)
		}
		// It's a decimal, return as float64
		return f
	case js.TypeString:
		return val.String()
	case js.TypeObject:
		if val.Get("constructor").Get("name").String() == "Array" {
			length := val.Get("length").Int()
			arr := make([]interface{}, length)
			for i := 0; i < length; i++ {
				arr[i] = jsToGo(val.Index(i))
			}
			return arr
		}
		// Regular object
		obj := make(map[string]interface{})
		keys := js.Global().Get("Object").Call("keys", val)
		keysLen := keys.Get("length").Int()
		for i := 0; i < keysLen; i++ {
			key := keys.Index(i).String()
			obj[key] = jsToGo(val.Get(key))
		}
		return obj
	default:
		return nil
	}
}

// Helper: Convert Go interface to JS value
func goToJs(val interface{}) js.Value {
	if val == nil {
		return js.Null()
	}

	switch v := val.(type) {
	case bool:
		return js.ValueOf(v)

	// Signed integers
	case int:
		return js.ValueOf(v)
	case int8:
		return js.ValueOf(v)
	case int16:
		return js.ValueOf(v)
	case int32:
		return js.ValueOf(v)
	case int64:
		return js.ValueOf(v)

	// Unsigned integers
	case uint:
		return js.ValueOf(v)
	case uint8:
		return js.ValueOf(v)
	case uint16:
		return js.ValueOf(v)
	case uint32:
		return js.ValueOf(v)
	case uint64:
		return js.ValueOf(v)

	// Floats
	case float32:
		return js.ValueOf(v)
	case float64:
		return js.ValueOf(v)

	// String
	case string:
		return js.ValueOf(v)

	// Slice of interfaces (array)
	case []interface{}:
		arr := js.Global().Get("Array").New(len(v))
		for i, item := range v {
			arr.SetIndex(i, goToJs(item))
		}
		return arr

	// Map (object)
	case map[string]interface{}:
		obj := js.Global().Get("Object").New()
		for key, item := range v {
			obj.Set(key, goToJs(item))
		}
		return obj

	// Byte array (Uint8Array)
	case []byte:
		dst := js.Global().Get("Uint8Array").New(len(v))
		js.CopyBytesToJS(dst, v)
		return dst

	default:
		// Fallback: log unsupported type and return null
		println("[BEVE WASM] Warning: Unsupported type in goToJs:", getTypeName(val))
		return js.Null()
	}
}

func main() {
	// Register BEVE functions in JavaScript global scope
	js.Global().Set("beveWasm", map[string]interface{}{
		"marshal":   js.FuncOf(marshal),
		"unmarshal": js.FuncOf(unmarshal),
		"version":   js.FuncOf(version),
		"benchmark": js.FuncOf(benchmark),
	})

	println("🎉 BEVE-Go WASM loaded successfully!")
	println("Available functions: beveWasm.{marshal, unmarshal, version, benchmark}")

	// Keep the program running
	select {}
}
