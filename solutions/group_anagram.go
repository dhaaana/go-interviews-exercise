package solutions

import (
	"fmt"
	"strings"
)

// Problem: https://leetcode.com/problems/group-anagrams/
// 49. Group Anagrams

func GroupAnagram(strs []string) [][]string {
	hashmap := map[string][]string{}
	result := [][]string{}

	for _, str := range strs {
		char := make([]int, 26)
		for _, s := range str {
			char[int(s-'a')]++
		}

		hashmap[intsToString(char)] = append(hashmap[intsToString(char)], str)
	}

	for _, value := range hashmap {
		result = append(result, value)
	}

	return result
}

func intsToString(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = fmt.Sprint(num)
	}
	return strings.Join(strs, ",")
}
