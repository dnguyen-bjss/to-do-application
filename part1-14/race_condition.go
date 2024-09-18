package main

import (
	"fmt"
	"sync"
	"time"
)
	
func addNumbers(wg *sync.WaitGroup, numbers *[]int, max int, startNumber int) {
	defer wg.Done()
    for i := startNumber; i < max; i += 2 {
		*numbers = append(*numbers, i)
		fmt.Printf("Added %d\n", i)
		time.Sleep(1 * time.Millisecond)
    }
}

func main() {
	var numbers []int;
	var wg sync.WaitGroup
	
	wg.Add(2)

	max := 100
	go addNumbers(&wg, &numbers, max, 0)
	go addNumbers(&wg,&numbers, max, 1)

	wg.Wait()
}