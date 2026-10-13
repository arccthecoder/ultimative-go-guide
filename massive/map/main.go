package main

import "fmt"

func main() {
	weather := map[int]int{
		11: +3,
		12: +6,
		13: +9,
		14: -4,
		15: +1,
	}

	for k, _ := range weather {
		weather[k] += 1
	}

	weather[20] = -10

	/*c, ok := weather[30]
	if ok {
		fmt.Println("okay")
	} else {
		fmt.Println("bad bro")
	}

	fmt.Println(c, ok)*/

	criminal := map[string]bool{
		"Vasya": true,
		"Petya": false,
		"Anton": false,
		"Vova":  false,
		"Alex":  true,
	}

	c, ok := criminal["Mekit"]

	if !ok {
		fmt.Println("No")
		return
	}

	fmt.Println("Найден!")
	if c {
		fmt.Println("Судим")
	} else {
		fmt.Println("Не судим")
	}

}
