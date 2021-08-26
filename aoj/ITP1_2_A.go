package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)

	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else if a == b {
		fmt.Println("a == b")
	}

}
