package main

import "fmt"

type Passport struct {
	number    int
	firstname string
	lastname  string
}

func (p *Passport) getFullName() string {
	return p.firstname + " " + p.lastname
}

func main() {
	p := &Passport{123456789, "Lee", "Provoost"}
	fmt.Println(p.getFullName())
}
