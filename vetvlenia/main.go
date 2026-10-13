package main // объявление главного пакета

import "fmt" // импорт библиотеки fmt для принта итд..

func main() { // объявление главной части кода
	score := 16 // укороченное обозначение переменной с авто-типом

	// if score > 10 {
	// 	if score > 15 {
	// 		fmt.Println("Ты мега красава")
	// 	} else {
	// 		fmt.Println("Ты прост красава")
	// 	}
	// } else {
	// 	fmt.Println("фу нуб")
	// }

	// 	score := 17

	// 	if score > 15 {
	// 		fmt.Println("Ты мега красваа")
	// 	} else if score > 10 {				// укороч вложенность
	// 		fmt.Println("Ты красава")
	// 	} else {
	// 		fmt.Println("иди учи уроки ебанатище")
	// 	}

	// score := 21

	if score == 12 {
		fmt.Println("Дюжина!")
	} else if score == 21 {
		fmt.Println("Очко!")
	} else if score == 50 {
		fmt.Println("Полтос")
	} else {
		fmt.Println("ты никуда не попал нубяра")
	}

}
