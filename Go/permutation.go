package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	n := len(nums)
	var backTrack func(int)
	backTrack = func(first int) {
		if first == n {
			temp := make([]int, n)
			copy(temp, nums)
			res = append(res, temp)
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backTrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	backTrack(0)
	return res
}

func main() {
	p := []int{1, 2, 3}
	fmt.Println(permute(p))
}
