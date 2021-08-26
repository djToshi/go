package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)

	fmt.Print(a * b)
	fmt.Print(" ")
	fmt.Println((a + b) * 2)
}
