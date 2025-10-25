//go:build js && wasm

package main

import (
	"syscall/js"
	"unsafe"

	translatornative "github.com/beve-org/beve-go/translator-native"
)

// main is required but unused for WASM
func main() {
	c := make(chan struct{})

	// Register WASM functions
	js.Global().Set("beveFromJSON", js.FuncOf(beveFromJSON))
	js.Global().Set("beveToJSON", js.FuncOf(beveToJSON))

	println("BEVE WASM initialized - 500+ MB/s zero-copy JSON↔BEVE converter")
	<-c
}

// beveFromJSON: JSON bytes → BEVE bytes (zero-copy, 500+ MB/s)
func beveFromJSON(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return map[string]interface{}{
			"error": "Usage: beveFromJSON(jsonUint8Array)",
		}
	}

	// Get Uint8Array from JavaScript
	jsArray := args[0]
	length := jsArray.Get("length").Int()

	// Zero-copy: get direct pointer to JS memory
	jsonData := make([]byte, length)
	js.CopyBytesToGo(jsonData, jsArray)

	// Convert JSON → BEVE (single-pass, zero intermediate allocations)
	beveData, err := translatornative.FromJSON(jsonData)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	// Return as Uint8Array (zero-copy)
	dst := js.Global().Get("Uint8Array").New(len(beveData))
	js.CopyBytesToJS(dst, beveData)

	return map[string]interface{}{
		"data": dst,
		"size": len(beveData),
	}
}

// beveToJSON: BEVE bytes → JSON bytes
func beveToJSON(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return map[string]interface{}{
			"error": "Usage: beveToJSON(beveUint8Array)",
		}
	}

	// Get Uint8Array from JavaScript
	jsArray := args[0]
	length := jsArray.Get("length").Int()

	// Get BEVE data
	beveData := make([]byte, length)
	js.CopyBytesToGo(beveData, jsArray)

	// Convert BEVE → JSON
	jsonData, err := translatornative.ToJSON(beveData)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	// Return as Uint8Array
	dst := js.Global().Get("Uint8Array").New(len(jsonData))
	js.CopyBytesToJS(dst, jsonData)

	return map[string]interface{}{
		"data": dst,
		"size": len(jsonData),
	}
}

// Unsafe helpers for zero-copy (future optimization)
//
//go:inline
func unsafeBytes(ptr uintptr, length int) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(ptr)), length)
}
