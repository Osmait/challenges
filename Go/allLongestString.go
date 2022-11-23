package main

import "fmt"

func solution(inputArray []string) []string {

	n := []int{}

	j := []string{}

	for _, v := range inputArray {
		n = append(n, len(v))

	}

	max := n[0]

	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]

		}
	}

	for _, v := range inputArray {
		if len(v) == max {
			j = append(j, v)
		}
	}

	return j
}

func main() {
	a := []string{"abacaba",
		"abacab",
		"abac",
		"xxxxxx"}
	fmt.Println(solution(a))
}
