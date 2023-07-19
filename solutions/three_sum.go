package solutions

import (
	"sort"
)

// Problem: https://leetcode.com/problems/3sum/submissions/
// 15. 3Sum

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 3 {
		var s int
		for _, num := range nums {
			s += num
		}
		if s == 0 {
			return [][]int{nums}
		}
		return [][]int{}
	}

	sort.Ints(nums)

	var prev int

	for i, num := range nums {
		if i == 0 {
			prev = num
		} else {
			if prev == num {
				continue
			}
		}

		restOfNums := nums[i+1:]

		l := 0
		r := len(restOfNums) - 1
		prevL := 0
		prevR := len(restOfNums) - 1

		for l < r {
			if l != prevL && restOfNums[l] == restOfNums[prevL] {
				prevL = l
				l++
				continue
			}
			if r != prevR && restOfNums[r] == restOfNums[prevR] {
				prevR = r
				r--
				continue
			}
			prevL = l
			prevR = r
			if restOfNums[l]+restOfNums[r]+num == 0 {
				result = append(result, []int{nums[i], restOfNums[l], restOfNums[r]})
				r--
				l++
				continue
			}
			if restOfNums[l]+restOfNums[r]+num > 0 {
				r--
				continue
			}
			if restOfNums[l]+restOfNums[r]+num < 0 {
				l++
				continue
			}
		}

		prev = num
	}

	return result
}
