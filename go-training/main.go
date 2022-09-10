package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, index int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	fmt.Println(index)
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

// func main() {
// 	c := SafeCounter{v: make(map[string]int)}
// 	for i := 0; i < 1000; i++ {
// 		go c.Inc("somekey", i)
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(c.Value("somekey"))
// }

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	Hatched       bool
	CreateReptile ReptileCreator
}

func (egg *ReptileEgg) Hatch() Reptile {
	if egg.Hatched == false {
		egg.Hatched = true
		return egg.CreateReptile()
	}
	return nil
}

type FireDragon struct {
}

func (f *FireDragon) Lay() ReptileEgg {
	return ReptileEgg{
		CreateReptile: func() Reptile {
			return &FireDragon{}
		}}
}

func main() {
	var testMap map[string]int
	fmt.Print(testMap == nil)
	fmt.Print(testMap)
}
