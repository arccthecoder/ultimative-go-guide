package main

import "fmt"

func main() {
	fmt.Println("я main и я начался")
	defer func() { // анонимная функция
		fmt.Println("я main и законч")
	}()

	database()

	//func hello() {
	defer func() { // анонимная функция
		fmt.Println("я хелло и законч")
	}()
	fmt.Println("я функция hello1")
}

func database() {
	// создать подключение
	defer func() {
		// закрыть подключение
	}()

	boolean := true

	if boolean {
		//
		return
	} else {

	}

	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			//
			// закрыть подключение
			return
		} else {
			//
			// закрыть подключение
			return
		}
	}

	// закрыть подключение
}
