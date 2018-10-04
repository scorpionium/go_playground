package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	s := []int{2, 7, 11, 15}
	t := 9
	res := twoSum(s, t)
	fmt.Printf("[%d, %d]", res[0], res[1])
}
