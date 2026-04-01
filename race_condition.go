package main

import (
	"net/http"
	"sync"
)

// Global mutable state without proper synchronization
var (
	cache     = make(map[string]string)
	cacheLock sync.Mutex // defined but inconsistently used
)

// Race condition — reads map without lock
func getFromCache(key string) (string, bool) {
	val, ok := cache[key]
	return val, ok
}

// Only write path is locked, read path above is not
func setCache(key, value string) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	cache[key] = value
}

// HTTP handler with race condition on shared state
var requestCount int

func countingHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++ // data race in concurrent HTTP server
	w.Write([]byte("ok"))
}

// Mutex copied by value — will not work correctly
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c SafeCounter) Inc(key string) { // receiver is value, not pointer — copies mutex
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

// WaitGroup misuse — Add inside goroutine
func waitGroupMisuse() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(n int) {
			wg.Add(1) // Add called inside goroutine — race with Wait
			defer wg.Done()
			_ = n * 2
		}(i)
	}
	wg.Wait()
}
