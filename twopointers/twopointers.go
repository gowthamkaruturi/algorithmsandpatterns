package twopointers

import "slices"

func IsPalindrome(s string) bool {
	isPallindrome := true
	chars := []rune(s)
	if len(chars) <= 0 {
		return isPallindrome
	}
	start, end := 0, len(chars)-1

	for start < end {
		if string(chars[start]) != string(chars[end]) {
			isPallindrome = false
		}
		start = start + 1
		end = end - 1
	}
	return isPallindrome
}

func SumOfThreeNumbers(nums []int, targetSum int) []int {
	slices.Sort(nums)
	start, end, sum := 0, 0, 0
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		start = i + 1
		end = len(nums) - 1
		for start < end {
			sum = nums[i] + nums[start] + nums[end]
			if sum == targetSum {
				result = append(result, nums[i], nums[start], nums[end])
				start++
				end--
			} else if sum < targetSum {
				start++
			} else {
				end--
			}
		}

	}
	return result
}
