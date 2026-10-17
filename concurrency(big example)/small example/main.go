package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	ID    int
	Error error
}

func doTask() error {
	if rand.Intn(4) == 3 {
		return errors.New("Ошибка выполнения")
	}

	return nil
}

func ProcessTask(wg *sync.WaitGroup, transferPoint chan<- Result, n int) {
	fmt.Println("Запускаю", n, "задач")

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			timer := rand.Intn(12)
			fmt.Println("Обработка задачи", id)
			time.Sleep(time.Duration(timer) * time.Second)
			err := doTask()
			transferPoint <- Result{
				ID:    id,
				Error: err,
			}
		}(i)
	}
}

func main() {
	results := make(chan Result, 5)
	wg := &sync.WaitGroup{}
	wgk := &sync.WaitGroup{}
	number := rand.Intn(10) + 1

	fmt.Println("=== CONCURRENT TASK RUNNER ===")

	ProcessTask(wg, results, number)

	wgk.Add(1)
	go func() {
		defer wgk.Done()

		for {
			select {
			case result, ok := <-results:
				if !ok {
					fmt.Println("Все результаты получены")
					return
				} else if result.Error != nil {
					fmt.Println("Ошибка в задаче", result.ID, result.Error)
				} else {
					fmt.Println("Задача номер", result.ID, "выполнена")
				}
			case <-time.After(5 * time.Second):
				fmt.Println("Тайм-аут. Продолжаю ожидать результат")
				continue
			}
		}
	}()

	wgk.Add(1)
	go func() {
		defer wgk.Done()

		wg.Wait()
		close(results)
	}()

	wgk.Wait()
}
