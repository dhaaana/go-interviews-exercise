package solutions

import "sort"

// Problem: https://leetcode.com/problems/valid-anagram/
// 242. Valid Anagram

func IsAnagram(s string, t string) bool {
	chars := make([]int, 26)

	for _, v := range s {
		i := int(v - 'a')
		chars[i]++
	}

	for _, v := range t {
		i := int(v - 'a')
		chars[i]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

func SlowIsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := []rune(s)
	tArr := []rune(t)

	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i] < sArr[j]
	})

	sort.Slice(tArr, func(i, j int) bool {
		return tArr[i] < tArr[j]
	})

	return string(sArr) == string(tArr)
}
