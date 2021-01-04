package main

import "fmt"

func solution(a [][]int) int {
	total := 0
	for ix, x := range a {
		if ix / 2 != 0 {
			for i, y := range x[1:len(x)-1] {
				total += y[i] + y[i+1] + y[i+2]
			}
		} else {
			for _, y := range x {
				total += y
			}
		}
//		if ix + 3 > len(a) {
//			continue
//		}
//		for iy, y := range x {
//			if iy + 3 > len(x) {
//				continue
//			}
//			total += y
//			fmt.Print(y)
//		}
//		fmt.Println("")
	}
	return total
}

func main() {
	arr := [][]int{
		{1,1,1,0,0,0},
		{0,1,0,0,0,0},
		{1,1,1,0,0,0},
		{0,0,2,4,4,0},
		{0,0,0,2,0,0},
		{0,0,1,2,4,0},
	}
	fmt.Println(solution(arr))
}
