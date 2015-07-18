package main

import "fmt"

type Passport struct {
	number    int
	firstName string
	lastName  string
}

func main() {
	p := Passport{1, "Lee", "Provoost"}
	fmt.Println(p.number)
	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
}
