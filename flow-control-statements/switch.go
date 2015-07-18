package main

import "fmt"

func main() {
	os := "linux"

	switch os {
	case "linux":
		fmt.Println("LINUX")
	case "windows":
		fmt.Println("WINDOWS")
	default:
		fmt.Println("OSX")
	}
}
