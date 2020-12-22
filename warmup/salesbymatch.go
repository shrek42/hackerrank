package main

import (
	"fmt"
)

func sockMerchant(n int, ar []int) int {
	socksPerColor := make(map[int]int)
	for _, v := range ar {
		socksPerColor[v]++
	}
	var pairs int
	for k, v := range socksPerColor {
		fmt.Printf("key: %v, value: %d\n", k, v)
		if v % 2 == 0 {
			pairs += v / 2
		} else {
			if v != 0 && (v - 1) % 2 == 0 {
				pairs += (v - 1) / 2
			}
		}
	}
	return pairs
}


func main() {
	n := 9
	ar := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(n, ar))
}
