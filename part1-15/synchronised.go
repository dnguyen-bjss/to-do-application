package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
	
func addNumbers(wg *sync.WaitGroup, numberChannel chan int, max int, startNumber int) {
	defer wg.Done()
	mu.Lock()
    for i := startNumber; i < max; i += 2 {
		numberChannel <- i
		time.Sleep(1 * time.Millisecond)
    }
	mu.Unlock()
}

func main() {
	numberChannel := make(chan int)
	var wg sync.WaitGroup
	
	wg.Add(2)

	max := 100
	go addNumbers(&wg, numberChannel, max, 0)
	go addNumbers(&wg, numberChannel, max, 1)

	go func() {
		for number := range numberChannel {
			fmt.Printf("Received: %d\n", number)
		}
	}()

	wg.Wait()
	close(numberChannel)
}