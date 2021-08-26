package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	var h int
	var m int
	var s int
	h = (a / 60 / 60)
	m = (a - (h * 3600)) / 60
	s = (a - (h * 3600) - (m * 60))

	fmt.Print(h)
	fmt.Print(":")
	fmt.Print(m)
	fmt.Print(":")
	fmt.Println(s)
}
