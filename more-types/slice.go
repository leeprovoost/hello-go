package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []string{"Lee", "Provoost"}
	fmt.Println(a[0], a[1])

	b := make([]int, 5)
	fmt.Println(b)
	b = append(b, 1)
	fmt.Println(b)
	b[0] = 3
	b[1] = 10
	b[2] = 4
	fmt.Println(b)

	for i, v := range b {
		fmt.Println("i: " + strconv.Itoa(i) + ", v: " + strconv.Itoa(v))
	}

	for _, v := range b {
		fmt.Println("v: " + strconv.Itoa(v))
	}
}
