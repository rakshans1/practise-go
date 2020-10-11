package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	started := time.Now()
	foods := []string{"pizza", "pasta", "cake"}
	var wg sync.WaitGroup
	for _, f := range foods {
		wg.Add(1)
		go func(f string) {
			cook(f)
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Printf("done in %v\n", time.Since(started))
}

func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s...\n", food)
}
