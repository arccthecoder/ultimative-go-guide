package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) error {
	if user.Ballance-usd < 0 {
		return errors.New("Недостаточно средств!")

	}

	return nil
}

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	if c.Armor-10 <= 0 {
		return 0, errors.New("Мы не газуем чтоб тачка не развалилась")
	}

	kmch := rand.Intn(150)

	c.Armor -= 10

	return kmch, nil
}

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника", p)
		}
	}()

	slice := []int{1, 2, 3}
	fmt.Println(slice[4])

	car := Car{
		Armor: 25,
	}

	for {
		pp.Println("Car до:", car)

		kmch, err := car.Gas()
		if err != nil {
			fmt.Println("Ошибка!", err.Error())
			break
		}
		fmt.Println("Получившийся разгон:", kmch)
		pp.Println("Car после:", car)
		fmt.Println("")
	}

}
