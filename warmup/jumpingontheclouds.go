package main

import (
	"fmt"
)

func contval(clouds []int) int {
	jumps := 0
	c := 0
	for _, _ = range clouds {
		if c+2 > len(clouds) {
			break
		}
		if clouds[c+2] == 0 {
			jumps++
			c += 2
			continue
		} else {
			c++
			jumps++
		}
	}
	return jumps
}


func main() {
	// x := 8
	y := []int{0, 0, 1, 0, 0, 1, 0}
	fmt.Println(contval(y))
}
