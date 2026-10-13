package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- i
			i++

			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		i := 1
		for {
			strCh <- "hi" + strconv.Itoa(i)
			i++

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select { //записывает информацию сразу с двух каналов без ожидания
		case number := <-intCh:
			fmt.Println("intCh:", number)
		case strin := <-strCh:
			fmt.Println("strCh:", strin)
		default:
			fmt.Println("Никакой канал не готов")
		}
	}
}
