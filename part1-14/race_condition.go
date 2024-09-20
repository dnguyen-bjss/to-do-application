package main

import (
	"fmt"
	"sync"
)
	
func addNumbers(source string, wg *sync.WaitGroup, numbers *[]int, max int, startNumber int) {
	defer wg.Done()
    for i := startNumber; i < max; i += 2 {
		*numbers = append(*numbers, i)
		fmt.Printf("%s\t number: %d - numbers: %d\n", source, i, numbers)
    }
}

func main() {
	var numbers []int;
	var wg sync.WaitGroup
	
	wg.Add(2)

	max := 100
	go addNumbers("[even]", &wg, &numbers, max, 0)
	go addNumbers("[odd]", &wg, &numbers, max, 1)

	wg.Wait()

	fmt.Println("Length of numbers:", len(numbers))
}