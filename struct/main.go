package main

import "fmt"

type User struct {
	// не пустое
	Name string // ""

	// 0<age<150
	Age int // 0

	// не пустой
	PhoneNumber string // ""

	// закрыт ли?
	isClose bool // false

	// 0<rat<10
	Rating float64 // 0.0 -- 10.0
}

func newUser( // создание структуры юзера через метод и валидация
	name string,
	age int,
	phonenumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		return User{}
	}

	if age <= 0 || age >= 150 {
		return User{}
	}

	if phonenumber == "" {
		return User{}
	}

	if rating < 0.0 || rating > 10.0 {
		return User{}
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phonenumber,
		isClose:     isClose,
		Rating:      rating,
	}
}

func (u *User) changeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) changeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) changePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	}
}

func (u *User) closeAccount() {
	u.isClose = true
}

func (u *User) openAccount() {
	u.isClose = false
}

func main() {
	user := newUser(
		"Danchik",
		18,
		"79234604284",
		false,
		4.5,
	)

	fmt.Println("user:", user)

	fmt.Println("rating до:", user.Rating)

	user.ratingUp(1.5)

	fmt.Println("рейтинг после", user.Rating)
}

func (u *User) ratingUp(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
		fmt.Println("я добавил рейтинг")
	} else {
		fmt.Println("я не прошел валидацию")
	}
}

func (u *User) ratingDown(rating float64) {
	if u.Rating-rating >= 0.0 {
		u.Rating -= rating
		fmt.Println("я отнял рейтинг")
	} else {
		fmt.Println("я не прошел валидацию")
	}
}

func (u User) goodbye() {
	fmt.Println("Всем пока!")
	fmt.Println("Меня звали", u.Name)
	fmt.Println("Мой рейтинг был:", u.Rating)
	fmt.Println("")
}

/*func main() {
	user := User{
		Name:   "Васек",
		Rating: 6.0,
	}
	fmt.Println("user до:", user)
	fmt.Println("")

	// ptr := &user

	user.ratingUp(2.0)

	fmt.Println("user после:", user)
}
*/
