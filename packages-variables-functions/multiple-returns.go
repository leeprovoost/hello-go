package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("x: 5, y: 10")
	x, y := swap(5, 10)
	fmt.Println("x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y))
}

func swap(x int, y int) (int, int) {
	return y, x
}
