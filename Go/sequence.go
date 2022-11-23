package main

import "fmt"

/*
Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing sequence by removing no more than one element from the array.

Note: sequence a0, a1, ..., an is considered to be a strictly increasing if a0 < a1 < ... < an. Sequence containing only one element is also considered to be strictly increasing.
*/

func SolutionSequence(sequence []int) bool {
	screwups, tail := 0, 0
	for i := 1; i < len(sequence); i++ {
		if sequence[tail] < sequence[i] {
			tail = i
		} else {
			if tail == 0 || sequence[i] > sequence[tail-1] {
				tail = i
			}
			screwups++
		}
	}
	return screwups == 1
}

func main() {
	lista := []int{1, 3, 2}

	fmt.Print(SolutionSequence(lista))
}
