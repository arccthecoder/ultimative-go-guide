// техническая точка зрения горутин

package main

import (
	"fmt"
	"time"
)

// функция, описывающая поход рабочего в шахту
// Принимает:
// 1. Номер похода в шахту
// Возвращает:
// 1. Добытый уголь
func mine(trasferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")

	trasferPoint <- 10 // отдаем значение в канал
	fmt.Println("Поход номер", n, "уголь передал")
}

func main() {
	coal := 0

	transferPoint := make(chan int, 3) // инициализация канала
	// в буферизированном канале горутины не ждут приема значений и завершают работу
	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3) // запуск горутин

	time.Sleep(2 * time.Second)
	coal += <-transferPoint

	time.Sleep(2 * time.Second)
	coal += <-transferPoint

	time.Sleep(2 * time.Second)
	coal += <-transferPoint // принимаем значение с канала

	fmt.Println("Добыли", coal, "угля")
	fmt.Println("Прошло времени:", time.Since(initTime))

	// время для видимой работы горутин
}
