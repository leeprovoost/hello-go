package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	var y int

	x, y = swap(x, y)

	fmt.Println("x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y))
}

func swap(a int, b int) (int, int) {
	return b, a
}
