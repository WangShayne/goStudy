package main

import (
	"fmt"
)

type Phone struct {
	Name string
	Call bool
}

func (a Phone) Message() {
	fmt.Println("I can message")
}

type Nokia struct {
	Model string
	Phone
}

func (a Nokia) Zao(name string) {
	a.Name = name
	fmt.Println("I can za hetao")
}

type Apple struct {
	Model string
	Phone
}

func (a Apple) Photo() {
	fmt.Println("I can photo")
}

func main() {
	a := Nokia{
		Model: "5230",
		Phone: Phone{
			Name: "nokia",
			Call: true,
		},
	}

	b := Apple{
		Model: "iphone",
		Phone: Phone{
			Name: "X",
			Call: false,
		},
	}
	fmt.Printf("%v", a)
	a.Message()
	a.Zao("sting")
	fmt.Printf("%v",a)
	fmt.Printf("%v", b)
	b.Message()
	b.Photo()
}
