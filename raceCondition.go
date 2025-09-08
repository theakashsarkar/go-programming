package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	time.Sleep(1 * time.Second) // wait for goroutines to finish
	fmt.Println("Final counter:", counter)
}
