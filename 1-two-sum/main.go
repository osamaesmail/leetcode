package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum2([]int{2,7,11,15}, 9))
	fmt.Println(twoSum2([]int{3,2,4}, 6))
	fmt.Println(twoSum2([]int{3,3}, 6))
	fmt.Println(twoSum2([]int{3, 2, 3}, 6))
}

// brute force solution
func twoSum1(nums []int, target int) []int {
	arrLrn := len(nums)
	for i := 0; i < arrLrn; i++ {
		for k := i + 1; k < arrLrn; k++ {
			if nums[i] + nums[k] == target {
				return []int{i, k}
			}
		}
	}

	return nil
}

// map solution
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, num := range nums {
		complement := target - num
		if k, exist := m[complement]; exist {
			return []int{k, i}
		}
		m[nums[i]] = i
	}

	return nil
}