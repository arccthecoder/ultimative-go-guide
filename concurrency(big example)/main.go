package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"time"
)

func main() {
	var coal int
	var mails []string

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

	coalTransferPoint := miner.MinerPool(minerContext, 10)
	mailTransferPoint := postman.PostmanPool(postmanContext, 10)

	isCoalClosed := false
	isMailClosed := false

	for !isCoalClosed || !isMailClosed {
		select {
		case c, ok := <-coalTransferPoint:
			if !ok {
				isCoalClosed = true
				continue
			}

			coal += c

		case m, ok := <-mailTransferPoint:
			if !ok {
				isMailClosed = true
				continue
			}

			mails = append(mails, m)
		}
	}

	fmt.Println("Суммарно добытый уголь:", coal)
	fmt.Println("Суммарное количество полученных писем:", len(mails))

}
