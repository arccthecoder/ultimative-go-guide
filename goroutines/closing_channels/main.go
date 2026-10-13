// РАСКОММЕНТИТЬ НУЖНУЮ ЧАСТЬ КОДА ДО РАЗДЕЛИТЕЛЯ "-----.."
// 1-я часть - работа с закрытым каналом
// 2-я часть - рабочий пример на анонимной горутине и советы

package main

func main() {
	// // создаем новый открытый канал
	// ch := make(chan int)

	// // закрываем канал
	// close(ch)

	// // читаю значение из закрытого канала
	// v1, ok1 := <-ch
	// v2, ok2 := <-ch
	// v3, ok3 := <-ch

	// fmt.Println("v:", v1, v2, v3)     // выводит дефолт значение для типа
	// fmt.Println("ok:", ok1, ok2, ok3) // проверка на наличие информации в канале

	// // записывание значения в закрытый канал
	// // ch <- 10 --- сработает паника
	// -------------------------------

	// transferPoint := make(chan int)

	// miner (с прошлого примера)
	// go func() {
	// 	iterations := 3 + rand.Intn(4)
	// 	fmt.Println("iterations:", iterations)

	// 	for i := 1; i <= iterations; i++ {
	// 		transferPoint <- 10
	// 		time.Sleep(300 * time.Millisecond)
	// 	}

	// 	close(transferPoint)
	// }()

	// coal := 0
	// for {
	// 	v, ok := <-transferPoint
	// 	if !ok {
	// 		fmt.Println("Весь уголь добыт!")
	// 		break
	// 	}

	// 	coal += v
	// 	fmt.Println("coal:", coal)
	// }

	// for v := range transferPoint { // удобная запись - не словим дедлок
	// 	coal += v
	// 	fmt.Println("coal:", coal)
	// }

	// fmt.Println("Суммарно добыто:", coal, "угля")
}
