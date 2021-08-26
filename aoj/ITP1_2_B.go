package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var c int
	fmt.Scan(&c)

	numbers := []int{a, b, c}
	sort.Ints(numbers) // sort.Sort(sort.IntSlice(numbers))
	fmt.Printf("%d %d %d\n", numbers[0], numbers[1], numbers[2])
}
