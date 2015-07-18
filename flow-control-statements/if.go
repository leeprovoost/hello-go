package main

import "fmt"

func main() {
	sum := 0

	for sum < 10 {
		if sum%2 == 0 {
			fmt.Println(sum)
		} else {
			fmt.Println("meh")
		}
		sum += 1
	}
}
