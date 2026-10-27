package main

import (
	"fmt"
	"net/http"
	"time"
)

// для работы необходимо установить Postman и через него посылать запросы на localhost:9091
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	fmt.Println("Handler отработал успешно")
}

func handlerSleep(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	_, err := w.Write([]byte("HTTP response!"))
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sleep", handlerSleep)

	err := http.ListenAndServe(":9091", nil) // на каком порту слушаем запросы
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}
