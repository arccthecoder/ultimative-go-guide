package main

import "fmt"

func main() {
	score := 10

	// score < 6 == false
	// score > 16 == false
	if score < 6 || score > 16 {
		fmt.Println("плохо")
	} else {
		fmt.Println("четко")
	}

	// score > 8 == true
	// score < 15 == true
	if score > 8 && score < 15 {
		fmt.Println("яблочко")
	} else {
		fmt.Println("не попал в яблочко")
	}

	star := false

	if !star {
		fmt.Println("поставь стар")
	}
}
