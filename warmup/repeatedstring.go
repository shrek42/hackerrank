package main

import "fmt"

func solution(s string, n int) int {
	result := ""
	for len(result) < n {
		result += s
	}
	if len(result) < n && len(result) != n {
		dif := n - len(result)
		result += s[:dif]
	}
	counter := 0
	for _, x := range result {
		if string(x) == "a" {
			counter++
		}
	}
	return counter
}

func main() {
	s := "abcac"
	n := 10
	fmt.Println(solution(s, n))
}
