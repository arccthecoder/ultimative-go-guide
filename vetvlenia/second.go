package main

import "fmt"

//func main() {
	score := 10

	// score < 6 == false
	// score > 16 == false
	if score < 6 || score > 16 {
		fmt.Println("петушара")
	} else {
		fmt.Println("ты нге петушара")
	}

	// score > 8 == true
	// score < 15 == true
	if score > 8 && score < 15 {
		fmt.Println("яблочко")
	} else {
		fmt.Println("нихуя не яблочко")
	}

	subscribed := false

	if !subscribed {
		fmt.Println("подпишись дурак")
	}
}
