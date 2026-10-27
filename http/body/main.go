package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx = sync.Mutex{}
var money = 1000 // usd
var bank = 0     // usd

func payHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body) // читаем тело запроса
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody) // переводим байты в строку

	paymentAmount, err := strconv.Atoi(httpRequestBodyString) // переводим строку в число
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert HTTP body to integer" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money -= paymentAmount

		msg := "Оплата прошла успешно:" + strconv.Itoa(money)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
	} else {
		fmt.Println("Не хватает денег на проведение оплаты")
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "fail to read HTTP request body:" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
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

		msg := "Новое значение money:" + strconv.Itoa(money) + "\nНовое значение bank:" + strconv.Itoa(bank)
		fmt.Println(msg)
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
