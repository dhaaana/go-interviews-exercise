package solutions

// Problem: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// 167. Two Sum II - Input Array Is Sorted

func TwoSumSorted(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
		if numbers[l]+numbers[r] > target {
			r--
		}
		if numbers[l]+numbers[r] < target {
			l++
		}
	}
}
