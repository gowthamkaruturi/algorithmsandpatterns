package twopointers

import (
	"fmt"
	"slices"

	"github.com/algorithmsandpatterns/models"
	"github.com/algorithmsandpatterns/utils"
)

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
func ContainerWithMostWater(a []int) int {
	left, right, maxWater := 0, len(a)-1, 0

	for left < right {
		width := right - left
		maxWater = utils.Max(maxWater, utils.Min(a[left], a[right])*width)
		if a[left] < a[right] {
			left++
		} else {
			right--
		}

	}

	return maxWater
}

// division solutuion does not work if we have zeros in the array
func ProductOfArrayOtherItSelf(nums []int) []int {
	result := make([]int, 0)
	finalProduct := 1
	for i := 0; i < len(nums); i++ {
		finalProduct = finalProduct * nums[i]
	}
	for i := 0; i < len(nums); i++ {
		result = append(result, finalProduct/nums[i])
	}
	return result
}

// this solution works for all integers
func ProductOfArrayOtherThanItself(nums []int) []int {
	//intialize this to store the resulet
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 1
	}
	leftResult, rightResult, left, right := 1, 1, 0, n-1
	for left < n && right >= 0 {
		result[left] = result[left] * leftResult
		result[right] = result[right] * rightResult
		fmt.Println(result)
		leftResult = leftResult * nums[left]
		rightResult = rightResult * nums[right]
		left++
		right--
	}
	return result
}

func RemoveNNodeFroLinkedList(head *models.LinkedListNode, n int) *models.LinkedListNode {
	slow := head
	fast := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func SortColors(colors []int) []int {
	red, white, blue := 0, 0, len(colors)-1
	for white <= blue {
		if colors[white] == 0 {
			//check if colors are in valid state
			if colors[red] != 0 {
				colors[red], colors[white] = colors[white], colors[red]

			}
			white++
			red++
		} else if colors[white] == 1 { // if white is right
			white++
		} else {
			if colors[blue] != 2 {
				colors[white], colors[blue] = colors[blue], colors[white]
			}
			blue--
		}
		white++
	}
	return colors
}
