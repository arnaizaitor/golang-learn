package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, Goldfinger!")
	fmt.Println("")

	var l int = len(os.Args) - 1
	var s string

	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}

	fmt.Printf("The %v input arguments are: %v", l, s)
}
