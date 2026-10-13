package main

import "fmt"

func main() {
	number := 999

	// i := 1 -- объявление переменной
	// i <= 5 -- условие, при котором выполняется цикл
	// i++ -- изменение от одной итерации к другой
	for i := 1; i <= 5; i++ {
		score := 5

		score = score + 3
		fmt.Println(number)
		fmt.Println("hello world", i)
		number++
	}
}
