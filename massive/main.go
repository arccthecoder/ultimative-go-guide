package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	arr := [5]int{5, 66, 7, 100, 1} // статический массив

	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			arr[i] *= 2
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i, "-", arr[i])
	}

	if arr[0]%2 == 0 {
		arr[0] *= 2
	}

	userArray := []User{
		User{
			Name:    "Danil",
			Rating:  5.5,
			Premium: true,
		},
		User{
			Name:    "Sergey",
			Rating:  3,
			Premium: false,
		},
		User{
			Name:    "Leha",
			Rating:  7.5,
			Premium: true,
		},
	}

	for i := 0; i < len(userArray); i++ {
		if userArray[i].Premium {
			userArray[i].Rating += 1
		}
	}

	for i := 0; i < len(userArray); i++ {
		pp.Println(userArray[i])
	}

	for index, user := range userArray {
		pp.Println(index, user)
	}

	for i, v := range userArray {
		if v.Premium {
			userArray[i].Rating += 1.0
		}
	}

	userArray = append(
		userArray,
		User{
			Name:    "Виталя",
			Rating:  4.0,
			Premium: true,
		})

	for _, v := range userArray {
		pp.Println(v)
	}

}
