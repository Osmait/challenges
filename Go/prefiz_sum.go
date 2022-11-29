package main

import "fmt"

type PrefixSum struct {
	Prefix_sum []int
}

func (p *PrefixSum) GetSum(start, end int) int {

	if start == 0 {
		return p.Prefix_sum[end]
	}
	return p.Prefix_sum[end] - p.Prefix_sum[start-1]
}
func ConstructorPrefix(array []int) *PrefixSum {
	return &PrefixSum{
		Prefix_sum: array,
	}
}
func main() {
	ar := []int{1, 2, 3, 4}
	p := ConstructorPrefix(ar)

	fmt.Println(p.GetSum(0, 2))
}
