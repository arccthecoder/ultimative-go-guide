package main

import (
	"fmt"
	"time"
)

func foo() {
	for {
		fmt.Println("hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go foo()

	go func() {
		for {
			fmt.Println("anon")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1000 * time.Millisecond)
}
