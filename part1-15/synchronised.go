package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

type numberStruct struct {
	source string
	number int
}

func addNumbers(source string, wg *sync.WaitGroup, numberChannel chan numberStruct, max int, startNumber int) {
	defer wg.Done()
    for i := startNumber; i < max; i += 2 {
		func() {
			mu.Lock()
			defer mu.Unlock()
			numberChannel <- numberStruct{source, i}
			time.Sleep(1 * time.Millisecond)
		} ()
    }
}

func main() {
	numberChannel := make(chan numberStruct)
	done := make(chan bool)
	var wg sync.WaitGroup
	
	wg.Add(2)

	max := 100
	go addNumbers("[even]", &wg, numberChannel, max, 0)
	go addNumbers("[odd]", &wg, numberChannel, max, 1)

	numberCount := 0
	go func () {
		for number := range numberChannel {
			numberCount++
			fmt.Printf("%s\t number: %d\n", number.source, number.number)
		}
		done <- true
	}()

	wg.Wait()
	close(numberChannel)
	<-done

	fmt.Println("Number count: ", numberCount)
}