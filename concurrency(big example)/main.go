package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64 // чтобы избавиться от гонки
	var mails []string
	mtx := sync.Mutex{} // чтобы избавиться от гонки

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("--->>> Рабочий день шахтеров окончен")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("--->>> Рабочий день почтальонов окончен")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 10)       // свое кол-во указать
	mailTransferPoint := postman.PostmanPool(postmanContext, 10) // свое кол-во указать

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("Суммарно добытый уголь:", coal.Load())

	mtx.Lock()
	fmt.Println("Суммарное количество полученных писем:", len(mails))
	mtx.Unlock()

}
