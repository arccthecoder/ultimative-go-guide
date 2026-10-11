package main

import "fmt"

func main() {
	number := 15

	var ptr1 *int // nil-указатель

	fmt.Println("number", number)
	fmt.Println("ptr1", ptr1)

	if ptr1 != nil {
		fmt.Println("четко на -", ptr1)
	} else {
		fmt.Println("там нил брат")
	}

	name := "Игорь"
	fmt.Println(name)
	ptr := &name

	changeName(ptr)

	fmt.Println(name, *ptr)

	pointer := &number // & - указатель

	foo(pointer)

	fmt.Println(number)
	fmt.Println("разыменовываем поинтер", *pointer)
}

func foo(n *int) { // * означает что ожидаем указатель
	fmt.Println(n)
	fmt.Println(*n)

	*n = 101
}

func changeName(n *string) {
	*n = "Ivan"
}

func boo(n int) {
	n = 1011
}
