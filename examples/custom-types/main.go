package main

import (
	"encoding/binary"
	"fmt"
	"log"

	beve "github.com/beve-org/beve-go"
)

// Point represents a 2D point with custom binary encoding
type Point struct {
	X, Y float64
}

// MarshalBinary implements encoding.BinaryMarshaler
func (p Point) MarshalBinary() ([]byte, error) {
	buf := make([]byte, 16)
	binary.LittleEndian.PutUint64(buf[0:8], uint64(p.X))
	binary.LittleEndian.PutUint64(buf[8:16], uint64(p.Y))
	return buf, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (p *Point) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return fmt.Errorf("invalid data length: %d", len(data))
	}
	p.X = float64(binary.LittleEndian.Uint64(data[0:8]))
	p.Y = float64(binary.LittleEndian.Uint64(data[8:16]))
	return nil
}

// Shape with custom type using BinaryMarshaler
type Shape struct {
	Name   string
	Center Point
	Points []Point
}

func main() {
	shape := Shape{
		Name:   "Triangle",
		Center: Point{X: 10.5, Y: 20.5},
		Points: []Point{
			{X: 0, Y: 0},
			{X: 10, Y: 0},
			{X: 5, Y: 10},
		},
	}

	// Marshal
	data, err := beve.Marshal(shape)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Encoded shape with custom type in %d bytes\n", len(data))

	// Unmarshal
	var decoded Shape
	if err := beve.Unmarshal(data, &decoded); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decoded: %s\n", decoded.Name)
	fmt.Printf("  Center: (%.1f, %.1f)\n", decoded.Center.X, decoded.Center.Y)
	fmt.Printf("  Points: %d vertices\n", len(decoded.Points))
	for i, p := range decoded.Points {
		fmt.Printf("    [%d] (%.1f, %.1f)\n", i, p.X, p.Y)
	}
}
