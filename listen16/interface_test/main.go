package main

import (
	"fmt"
)

type Animal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {
}

func (d Dog) Talk() {
	fmt.Println("wang wang wang")
}

func (d Dog) Eat() {
	fmt.Println("chi dongxi")
}

func (d Dog) Name() string {
	fmt.Println("I am dog,ma name is wangcai")
	return "wangcai"
}

type Pig struct {
}

func (p Pig) Talk() {
	fmt.Println("heng heng heng")
}

func (p Pig) Eat() {
	fmt.Println("吃草")
}

func (p Pig) Name() string {
	fmt.Println("我是猪 ")
	return "猪八戒"
}

func main() {
	var d Dog
	var a Animal
	a = d
	a.Eat()
	a.Talk()
	a.Name()
	var pig Pig

	a = pig
	a.Eat()
	a.Talk()
	a.Name()
}
