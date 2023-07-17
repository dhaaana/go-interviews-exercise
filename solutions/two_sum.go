package solutions

import (
	"go-interviews-exercise/utils"
)

// Problem: https://leetcode.com/problems/two-sum/
// 1. Two Sum

func SlowTwoSum(nums []int, target int) []int {
	result := []int{}

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}

	return result
}

func TwoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	result := []int{}

	for i, num := range nums {
		hashmap[target-num] = i
		utils.PrintResult(hashmap)
	}

	for i, num := range nums {
		if value, isExist := hashmap[num]; isExist {
			if value == i {
				continue
			}
			result = append(result, value)
			result = append(result, i)
			break
		}
	}

	return result
}

// Basically the same with above but straight to the point
func EfficientTwoSum(nums []int, target int) []int {
	hashmap := map[int]int{}

	for i, num := range nums {
		if value, isExist := hashmap[target-num]; isExist {
			return []int{value, i}
		}
		hashmap[num] = i
	}

	return []int{}
}
