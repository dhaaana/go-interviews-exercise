package solutions

import (
	"strings"
	"unicode"
)

// Problem: https://leetcode.com/problems/valid-palindrome/
// 125. Valid Palindrome

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	// Convert to lowercase
	s = strings.ToLower(s)

	// Remove non-alphanumeric characters
	s = removeNonAlphanumeric(s)

	reversed := ""

	for _, v := range s {
		reversed = string(v) + reversed
	}

	return s == string(reversed)
}

// Helper function to remove non-alphanumeric characters
func removeNonAlphanumeric(s string) string {
	var sb strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

func PointerIsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)

	// Left and right pointer
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
