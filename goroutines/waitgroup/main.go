package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(text string, wg *sync.WaitGroup) {
	defer wg.Done() // завершение горутины

	for i := 1; i <= 3; i++ {
		fmt.Println("я почтальон, я отнес газету", text, "в", i, "раз")
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	wg := &sync.WaitGroup{} //указатель!

	wg.Add(1)
	go postman("Новости", wg)

	wg.Add(1)
	go postman("Игровой журнал", wg)

	wg.Add(1)
	go postman("авто-новости", wg)

	wg.Wait() // ждем выполнение всех горутин

	fmt.Println("done")

}
