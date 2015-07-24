package main

import "fmt"

type Passport struct {
	number    int
	firstname string
	lastname  string
}

var m map[string]Passport

func main() {
	m := make(map[string]Passport)
	m["123456789"] = Passport{123456789, "lee", "provoost"}
	m["987654321"] = Passport{987654321, "Hannah", "Mills"}
	fmt.Println(m["123456789"])
	fmt.Println(m)

	delete(m, "123456789")
	fmt.Println(m)

	elem, ok := m["987654321"]
	fmt.Println(ok)
	fmt.Println(elem)

	elem, ok = m["123456789"]
	fmt.Println(ok)
	fmt.Println(elem)
}
