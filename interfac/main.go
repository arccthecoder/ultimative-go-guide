package main

import "fmt"

type Auto interface { // позволяет принимать все функции, которые имеют метод StepOnGas 
	StepOnGas()
}

type BMW struct{}

func (b BMW) StepOnGas() {
	fmt.Println("Я БМВ ебейшая")
}

type Zhiga struct{}

func (z Zhiga) StepOnGas() {
	fmt.Println("Я жига полутычка")
}

func ride(auto Auto) {
	fmt.Println("я жму на газ типа")
	auto.StepOnGas()
}

func main() {
	bmw := BMW{}
	bmw.StepOnGas()

	zhiga := Zhiga{}
	zhiga.StepOnGas()
}
