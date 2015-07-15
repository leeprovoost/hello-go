package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 5
	y := 9 // shorthand

	x, y = swap(x, y)

	fmt.Println("x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y))
}

func swap(a int, b int) (int, int) {
	return b, a
}
