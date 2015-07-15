package main

import (
	"fmt"
)

func main() {
	fmt.Println(formatMessage("Lee"))
}

func formatMessage(name string) string {
	return "Hello " + name + "!"
}
