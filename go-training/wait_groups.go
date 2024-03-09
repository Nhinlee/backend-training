package main

import (
	"sync/atomic"
)

type CustomWaitGroup struct {
	counter uint64
	waitCh  chan bool // Channel for signaling
}

func (wg *CustomWaitGroup) Add(delta int) {
	atomic.AddUint64(&wg.counter, uint64(delta))
}

func (wg *CustomWaitGroup) Done() {
	atomic.AddUint64(&wg.counter, ^uint64(0)) // Decrement by 1
	if atomic.LoadUint64(&wg.counter) == 0 {
		close(wg.waitCh) // Signal completion
	}
}

func (wg *CustomWaitGroup) Wait() {
	<-wg.waitCh // Wait for signal
}

// func main() {
// 	// Custom wait group
// 	wg1 := CustomWaitGroup{waitCh: make(chan bool)}
// 	wg1.Add(2)
// 	go func() {
// 		defer wg1.Done()

// 		// Do something
// 		fmt.Println("First goroutine")
// 	}()
// 	go func() {
// 		defer wg1.Done()

// 		// Do something
// 		fmt.Println("Second goroutine")
// 	}()

// 	wg1.Wait()

// 	// Wait group from sync package
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()

// 		// Do something
// 		fmt.Println("First goroutine")
// 	}()
// 	go func() {
// 		defer wg.Done()

// 		// Do something
// 		fmt.Println("Second goroutine")
// 	}()

// 	wg.Wait()
// }
