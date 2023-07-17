package solutions

// Problem: https://leetcode.com/problems/contains-duplicate/

func ContainsDuplicate(nums []int) bool {
	hashmap := map[int]int{}

	for _, i := range nums {
		if _, isExist := hashmap[i]; isExist {
			return true
		}
		hashmap[i]++
	}
	return false
}
