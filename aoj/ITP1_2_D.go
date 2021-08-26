package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var W, H, x, y, r int
	fmt.Fscan(stdin, &W, &H, &x, &y, &r)

	if x+r <= W && y+r <= H && x-r >= 0 && y-r >= 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
