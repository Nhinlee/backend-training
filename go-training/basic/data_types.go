// Package basic provides examples of different data types in Go.
package basic

import (
	"fmt"
	"sync"
)

// DataTypes represents various data types in Go.
type DataTypes struct {
	// Primitive data types

	// boolVar represents a boolean value.
	boolVar bool

	// Numeric data types
	int8Var       int8       // int8Var represents an 8-bit signed integer.
	int16Var      int16      // int16Var represents a 16-bit signed integer.
	int32Var      int32      // int32Var represents a 32-bit signed integer.
	int64Var      int64      // int64Var represents a 64-bit signed integer.
	uint8Var      uint8      // uint8Var represents an 8-bit unsigned integer.
	uint16Var     uint16     // uint16Var represents a 16-bit unsigned integer.
	uint32Var     uint32     // uint32Var represents a 32-bit unsigned integer.
	uint64Var     uint64     // uint64Var represents a 64-bit unsigned integer.
	float32Var    float32    // float32Var represents a 32-bit floating-point number.
	float64Var    float64    // float64Var represents a 64-bit floating-point number.
	complex64Var  complex64  // complex64Var represents a complex number with float32 real and imaginary parts.
	complex128Var complex128 // complex128Var represents a complex number with float64 real and imaginary parts.

	// stringVar represents a string value.
	stringVar string

	// byteVar represents a byte value.
	byteVar byte

	// --------------------------------------------
	// Composite data types

	// arrayVar represents an array of 5 integers.
	arrayVar [5]int

	// sliceVar represents a slice of integers.
	sliceVar []int

	// mapVar represents a map with string keys and integer values.
	mapVar map[string]int

	// structVar represents a struct with name and age fields.
	structVar struct {
		name string
		age  int
	}

	// interfaceVar represents an interface value.
	interfaceVar interface{}

	// funcVar represents a function value.
	funcVar func()
}

// Pass map into function
type Person struct {
	name string
	age  int
}

func (d *DataTypes) PassMapIntoFunction(m map[string]Person) {
	// Get the value of the key "John" from the map.
	john := m["John"]
	// Try change the name of the person.
	john.name = "Jane"
	// Update the value of the key "John" in the map.
	m["John"] = john
}

// TestPassMapIntoFunction demonstrates how to use a map in Go.
func (d *DataTypes) TestPassMapIntoFunction() {
	// Create a map with string keys and integer values.
	m := map[string]Person{
		"John": {"John", 25},
	}

	// Pass the map into a function.
	d.PassMapIntoFunction(m)

	// Get the value of the key "John" from the map.
	john := m["John"]
	// Print the value of the key "John".
	fmt.Printf("Name: %s, Age: %d\n", john.name, john.age)
}

func (d *DataTypes) PassSliceIntoFunction(s []int) {
	// Try to change the first element of the slice.
	s[0] = 100
}

// TestPassSliceIntoFunction demonstrates how to use a slice in Go.
func (d *DataTypes) TestPassSliceIntoFunction() {
	// Create a slice of integers.
	s := []int{1, 2, 3}

	// Pass the slice into a function.
	d.PassSliceIntoFunction(s[0:1])

	// Print the slice after passing it into the function.
	fmt.Println(s)
}

// Test 2 functions consume the same channel
func (d *DataTypes) TestChannelType() {
	// Create a channel of integers.
	ch := make(chan int)

	// Add wait group to wait for all goroutines to finish.
	wg := sync.WaitGroup{}

	// Create a goroutine to send data to the channel.
	for i := 0; i < 10; i++ {
		go func(i int) {
			// Send data to the channel.
			ch <- i
		}(i)
		wg.Add(1)
	}

	// Create a goroutine to receive data from the channel.
	go func() {
		// Receive data from the channel.
		for i := range ch {
			fmt.Printf("Consumer 1: %d\n", i)
			wg.Done()
		}
	}()

	// Create a goroutine to receive data from the channel.
	go func() {
		// Receive data from the channel.
		for i := range ch {
			fmt.Printf("Consumer 2: %d\n", i)
			wg.Done()
		}
	}()

	wg.Wait()
	close(ch)
}
