package main

import (
	"fmt"
	"time"
)
	
func addNumbers(from string, numbers *[]int, max int, startNumber int) {
    for i := startNumber; i < max; i += 2 {
		*numbers = append(*numbers, i)
		fmt.Printf("%s: %d\n", from, i)
    }
}

func main() {
	var numbers []int;

	max := 100
	go addNumbers("Even", &numbers, max, 0)
	go addNumbers("Odd", &numbers, max, 1)

	time.Sleep(time.Second)
	fmt.Printf("Numbers: %d\n", numbers)
}