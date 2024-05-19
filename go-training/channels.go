package main

import (
	"fmt"
	"time"
)

// func main() {
	// ---------------------
	// //Create an un-buffered channel
	// baton := make(chan int)

	// // Start run
	// go Runner(baton)

	// // Start the
	// baton <- 1

	// // Give the runners time to race
	// time.Sleep(500 * time.Millisecond)

	// ---------------------
	// BufferedTest()

	// ---------------------
	//
	//	UnBuffered_WaitForTest()

	// Block mechanism ---------------------
	// Block1()
	// Block2()
	// Block3()

	// Load balancing ---------------------
	// LoadBalancingEx()

	// Timeout and cancel ---------------------
	// TimeoutAndCancelEx()

	// Pipeline ---------------------
	// PipelineEx()
// }

func Block1() {
	// Sending to an unbuffered channel with no receiver.
	// This will block the current goroutine.
	ch := make(chan int, 10)

	ch <- 1

	fmt.Println("Finished")
}

func Block2() {
	// Receiving from an empty channel
	// This will block the current goroutine.
	ch := make(chan int)

	<-ch

	fmt.Println("Finished")
}

func Block3() {
	// Sending to a full buffered channel
	// This will block the current goroutine.
	ch := make(chan int, 5)

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Sent", i)
	}

	fmt.Println("Finished")
}

func worker(id int, tasks <-chan int, results chan<- int) {
	// Process tasks from the input channel
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		t := 10
		if id == 1 {
			t = 1000
		}

		time.Sleep(time.Duration(t) * time.Millisecond) // Simulate processing time
		results <- task * 2                             // Send the result to the output channel
		fmt.Printf("Worker %d finished task %d\n", id, task)
	}
}

func LoadBalancingEx() {
	// Create channels for tasks and results
	tasks := make(chan int, 20)
	results := make(chan int)

	// Start multiple workers to process tasks
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	// Send some tasks to the input channel
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	defer close(tasks)

	// Collect the results from the output channel
	for i := 1; i <= 10; i++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}

func TimeoutAndCancelEx() {
	// Create a channel to receive the result
	result := make(chan int)

	// Start a goroutine to perform a long-running operation
	go func() {
		time.Sleep(3 * time.Second) // Simulate a long-running operation
		result <- 42
	}()

	// Wait for the result, but timeout if it takes too long
	timeout := time.After(2 * time.Second)
	select {
	case res := <-result:
		fmt.Println("Result:", res)
	case <-timeout:
		fmt.Println("Timed out!")
	}
}

func PipelineEx() {
	// Create channels to connect the different stages of the pipeline
	nums := make(chan int)
	squares := make(chan int)

	// Start a goroutine to generate a sequence of numbers
	go func() {
		for i := 1; i <= 10; i++ {
			nums <- i
			time.Sleep(time.Second)
		}
		close(nums)
	}()

	// Start a goroutine to square each number in the sequence
	go func() {
		for num := range nums {
			squares <- num * num
		}
		close(squares)
	}()
	// Read the squared numbers from the channel and print them
	for square := range squares {
		fmt.Println(square)
	}
}

func BufferedTest() {
	ch := make(chan string, 9)

	go func() {
		for d := range ch {
			fmt.Println(d)
		}
	}()

	ch <- "hello"
	ch <- "world 1"
	ch <- "world 2"
	ch <- "world 3"
	ch <- "world 4"
	ch <- "world 5"
	ch <- "world 6"
	ch <- "world 7"
	ch <- "world 8"
	ch <- "world 9"
	ch <- "world 10"

	fmt.Println("finished")
}

func UnBuffered_WaitForTest() {
	ch := make(chan string)

	go func() {
		fmt.Printf("employee : waiting for signal\n")
		// time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		p := <-ch

		// Employee performs work here.
		fmt.Println("employee : received signal :", p)

		// Employee is done and free to go.
		fmt.Println("employee : sent signal")
	}()

	// time.Sleep(100 * time.Millisecond)
	fmt.Printf("manager : send paper\n")

	ch <- "paper"

	// time.Sleep(1 * time.Millisecond)

	fmt.Println("manager : i'm done")
}
