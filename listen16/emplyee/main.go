package main

import (
	"fmt"
)

type Emplyee interface {
	CaleSalary() float32
}

type Programer struct {
	name  string
	base  float32
	extra float32
}

func NewProgramer(name string, base, extra float32) Programer {
	return Programer{
		name:  name,
		base:  base,
		extra: extra,
	}
}

func (p Programer) CaleSalary() float32 {
	return p.base
}

type Sala struct {
	name  string
	base  float32
	extra float32
}

func NewSala(name string, base, extra float32) Sala {
	return Sala{
		name:  name,
		base:  base,
		extra: extra,
	}
}

func (s Sala) CaleSalary() float32 {
	return s.base + s.extra*2
}

func CaleAll(e []Emplyee) float32 {
	var cont float32
	for _, v := range e {
		cont += v.CaleSalary()
	}
	return cont
}

func main() {
	p1 := NewProgramer("程序员1", 5000.0, 0)
	p2 := NewProgramer("程序员2", 5000.0, 0)
	p3 := NewProgramer("程序员3", 5000.0, 0)

	s1 := NewSala("销售1", 2000.0, 5)
	s2 := NewSala("销售2", 2000.0, 5)
	s3 := NewSala("销售3", 2000.0, 5)

	// emplyeelist := []Emplyee
	var emplyeelist []Emplyee
	emplyeelist = append(emplyeelist, p1)
	emplyeelist = append(emplyeelist, p2)
	emplyeelist = append(emplyeelist, p3)
	emplyeelist = append(emplyeelist, s1)
	emplyeelist = append(emplyeelist, s2)
	emplyeelist = append(emplyeelist, s3)
	cont := CaleAll(emplyeelist)
	fmt.Printf("当月合计发工资:%f", cont)
}
