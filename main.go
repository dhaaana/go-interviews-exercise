package main

import (
	"go-interviews-exercise/solutions"
	"go-interviews-exercise/utils"
)

func main() {
	utils.PrintResult(solutions.ContainsDuplicate([]int{1, 2, 3, 1}))
	utils.PrintResult(solutions.IsAnagram("race", "ecar"))
	utils.PrintResult(solutions.TwoSum([]int{2, 9, 11, 7, 15}, 9))
	utils.PrintResult(solutions.GroupAnagram([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
