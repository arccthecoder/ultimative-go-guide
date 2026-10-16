package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()

	// зависит от вашего желания способа работы. подробнее ниже
	for {
		fmt.Println("Я шахтер номер:", n, "Начал добывать уголь!")

		// сразу завершают рабочий день после отмены контекста
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер номер:", n, "Мой рабочий день закончен")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Я шахтер номер:", n, "Добыл уголь!")
		}

		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер номер:", n, "Мой рабочий день закончен")
			return
		case transferPoint <- power:
			fmt.Println("Я шахтер номер:", n, "Передал уголь:", power)
		}

		// доделывают работу и завершают после отмены контекста
		// select {
		// case <-ctx.Done():
		// 	fmt.Println("Я шахтер номер:", n, "Мой рабочий день закончен")
		// 	return
		// default:
		// 	fmt.Println("Я шахтер номер:", n, "Начал добывать уголь!")
		// 	time.Sleep(1 * time.Second)
		// 	fmt.Println("Я шахтер номер:", n, "Добыл уголь!")

		// 	transferPoint <- power
		// 	fmt.Println("Я шахтер номер:", n, "Передал уголь:", power)
		// }
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
