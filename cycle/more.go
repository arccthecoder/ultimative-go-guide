package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func main() {

	for i := 1; i <= 50; i += rand.Intn(20) {
		fmt.Println("Новая итерация! i:", i)
		time.Sleep(500 * time.Millisecond)
	}

}
