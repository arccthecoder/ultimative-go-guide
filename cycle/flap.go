package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	score := 0
	fmt.Println("гет реди")
	fmt.Println("")

	for { // бесконечный цикл (без условий)

		fmt.Println("вы подлетаете к трубе")
		fmt.Println("п тт")
		fmt.Println("")

		fmt.Println("вы пролетаете через трубу")
		fmt.Println("т п т")
		fmt.Println("")

		if rand.Intn(4) == 1 {
			fmt.Println("вы разбились к хуям(")
			fmt.Println("")
			break
		}

		fmt.Println("вы пролетели через трубу")
		fmt.Println("тт п")
		fmt.Println("")

		score++

		fmt.Println("Счет -", score)
		fmt.Println("")

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("game over")
	fmt.Println("ваш итоговый счет:", score)
}
