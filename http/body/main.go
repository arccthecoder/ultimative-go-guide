package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mtx = sync.Mutex{}
var money = 1000 // usd
var bank = 0     // usd

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body) // читаем тело запроса
	if err != nil {
		fmt.Println("fail to read HTTP body", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody) // переводим байты в строку

	paymentAmount, err := strconv.Atoi(httpRequestBodyString) // переводим строку в число
	if err != nil {
		fmt.Println("fail to convert HTTP body to integer", err)
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		time.Sleep(3 * time.Second)
		money -= paymentAmount
		fmt.Println("Оплата прошла успешно:", money)
	} else {
		fmt.Println("Не хватает денег на проведение оплаты")
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP request body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to integer", err)
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		// взять с кошелька
		money -= saveAmount

		// положить в копилку
		bank += saveAmount

		fmt.Println("Новое значение money:", money)
		fmt.Println("Новое значение bank:", bank)
	} else {
		fmt.Println("Не хватает денег на пополнение копилки")
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error:", err)
	}

}
