package main

import (
	"fmt"
	"net/http"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Новый платеж обработан!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно совершил оплату!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Оплата отменена!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно отменил оплату!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello, world!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я успешно обработал HTTP запрос!")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Запускаю HTTP сервер")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	}

	fmt.Println("Программа завершила свое выполнение")
}
