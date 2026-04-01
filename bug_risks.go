package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Nil pointer dereference — no nil check after type assertion
func processValue(v interface{}) string {
	s, _ := v.(*string)
	return *s
}

// Goroutine leak — channel never closed or read
func leakyGoroutine() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	// channel is never read, goroutine blocks forever
}

// Data race — shared variable without synchronization
func concurrentCounter() int {
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // race condition
		}()
	}
	wg.Wait()
	return counter
}

// Deferred call in a loop — defers won't run until function exits
func deferInLoop(files []string) {
	for _, f := range files {
		// This deferred close won't happen until the function returns,
		// potentially exhausting file descriptors
		fmt.Println(f)
	}
}

// Integer overflow — unchecked multiplication
func multiply(a, b int32) int32 {
	return a * b
}

// Using math/rand instead of crypto/rand for token generation
func generateToken() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Infinite recursion — no base case
func infinite(n int) int {
	return infinite(n - 1)
}

// Unreachable code after return
func unreachable(x int) int {
	if x > 0 {
		return x
	}
	return -x
	fmt.Println("this is never reached")
	return 0
}
