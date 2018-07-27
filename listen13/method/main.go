package main

import (
	"fmt"
)

type People struct {
	Name    string
	Country string
}

func (p People) Print() {
	fmt.Printf("name=%s country=%s\n", p.Name, p.Country)
}

func (p People) Set(name string, country string) {
	p.Name = name
	p.Country = country
}

func (p *People) SetV2(name string, country string) {
	p.Name = name
	p.Country = country
}

func main() {
	var p1 People = People{
		Name:    "people1",
		Country: "US",
	}
	p1.Print()
	p1.Set("people2", "UK")
	p1.Print()

	fmt.Println()
	p1.Print()
	p1.SetV2("people2", "UK")
	p1.Print()
	fmt.Printf("%v", p1)
}
