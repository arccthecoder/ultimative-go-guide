package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var mtx = sync.Mutex{}
var money = 1000
var paymentHistory = make([]Payment, 0)

type HttpResponse struct {
	Money          int       `json:"mmmoney"`
	PaymentHistory []Payment `json:"phistory"`
}

type Payment struct {
	// Описание покупки
	Description string `json:"description"`

	// Сумма покупки
	USD int `json:"usd"`

	// ФИО человека, совершающего покупку
	FullName string `json:"FullName"`

	// Адрес счета
	Address string `json:"address"`

	// Время проведения оплаты
	Time time.Time
}

// Выводит основные сведения о платеже в консоль.
func (p Payment) Println() {
	fmt.Println("Descripton:", p.Description)
	fmt.Println("Usd:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
}

// Проверяет, что платёж содержит сумму, имя покупателя и адрес.
func (p Payment) Validate() bool {
	if p.USD == 0 || p.Address == "" || p.FullName == "" {
		return false
	}
	return true
}

// Принимает платёж, обновляет баланс и возвращает баланс
// вместе с историей платежей в формате JSON.
func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	payment.Println()

	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	}

	paymentHistory = append(paymentHistory, payment)

	// Делаем JSON ответ клиенту
	HttpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.Marshal(HttpResponse)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}

	mtx.Unlock()
}

// Запускает HTTP-сервер и регистрирует обработчик платежей.
func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}

}
