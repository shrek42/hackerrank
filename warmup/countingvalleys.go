package main

import (
	"fmt"
)

func contval(x int, y string) int {
	valleys := 0 
	prev := 0
	count := 0
	for _, v := range y {
		if v == 'U' {
			count++
		}
		if v == 'D' {
			count--
		}
		if prev == 0 && count == -1 {
			valleys++
		}
		prev = count
	}
	return valleys
}


func main() {
	x := 8
	y := "UDDDUDUU"
	fmt.Println(contval(x, y))
}
