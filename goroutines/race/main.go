package main

import (
	"fmt"
	"sync"
)

// var number int = 0
// var number atomic.Int64
var slice []int

var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		// number.Add(1) // атомарная операция
		mtx.Lock() // "очередь" горутин --> гонки не будет
		slice = append(slice, i)
		mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	mtx.Lock()
	fmt.Println("slice len:", len(slice))
	mtx.Unlock()
}
