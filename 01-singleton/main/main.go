package main

import (
	"client_one"
	"client_two"
	"fmt"
	"singleton"
)

func main() {
	client_one.IncremetAge()
	client_two.IncremetAge()

	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
