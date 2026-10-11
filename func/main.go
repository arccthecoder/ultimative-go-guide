package main

import "fmt"

var number int = 5 // глобальная переменная

func main() {
	fmt.Println("до вызова")
	hello()
	aboba()
	fmt.Println("")

	square(5)
	square(11)
	square(123)

	ofic("ванек")

	sum(5, 6, "здаолва заебал")

	fmt.Println(suma(9, 632756483756834))

	fmt.Println("после вызова")
}

func hello() { // тест функция
	fmt.Println("hello 1")
	fmt.Println("hello 2")
}

func suma(a int, b int) int { // возвращает значение переменной
	s := a + b
	return s
}

func aboba() { //тест сумм переменных
	a := 0
	b := 10
	c := 3

	fmt.Println("Сумма:", a+b+c)
}

func square(x int) { // принимает значение типа int
	fmt.Println("приняли x:", x)
	fmt.Println("x в квадрате", x*x)
}

func sum(a int, b int, greetings string) { // принимает разные типы
	s := a + b

	fmt.Println(greetings)
	fmt.Println("сумма -", s)
	fmt.Println("")
}

/* допустим официант в ресторане
задача
для каждого нового гостя нужно:
1 накрыть ему стол
2 поприветствовать по имени
3 принять заказ
4 принести блюдо */

func ofic(name string) { // использование полученной переменной
	fmt.Println("Накрываю стол")
	fmt.Println("Привет,", name, "!")
	fmt.Println("я принял заказ")
	fmt.Println("я принес блюдо")
}
