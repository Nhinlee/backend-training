package basic

import (
	"fmt"
	"strconv"
)

type TypeConversions struct{}

// 3 things need to aware when converting types in Go:
// - Data loss
// - Overflow
// - Type assertion & conversion error

// Function demonstrates data loss when converting types in Go.
func (t *TypeConversions) DataLoss() {
	// Convert float64 to int.
	var y float64 = 3.14
	x := int(y)
	fmt.Printf("data loss: y = %f => x = %d\n", y, x)
}

// Function demonstrates overflow when converting types in Go.
func (t *TypeConversions) Overflow() {
	// Convert int to uint8.
	var y int32 = 1000
	x := uint8(y)
	fmt.Printf("overflow: y = %d => x = %d\n", y, x)
}

// Function demonstrates conversion error when converting types in Go.
func (t *TypeConversions) ConversionError() {
	// Convert string to int.
	var y string = "3.14"
	x, ok := strconv.Atoi(y)
	if ok != nil {
		fmt.Printf("conversion error: y = %s => x = %d\n", y, x)
	}
}

// Function demonstrates how to type assert in Go.
func (t *TypeConversions) TypeAssertion() {
	// Create a map with string keys and integer values.
	m := map[string]interface{}{
		"John": 25,
	}

	// Get the value of the key "John" from the map.
	v, ok := m["John"].(int)
	if ok {
		fmt.Printf("Type assertion: %d\n", v)
	}
}

// Summary:
// - Go requires explicit type conversions.
// - Be aware of data loss, overflow, and type assertion & conversion error.
// - Use type assertion to check the type of an interface value.
