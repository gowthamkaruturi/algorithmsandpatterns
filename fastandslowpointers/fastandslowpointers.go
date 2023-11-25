package fastandslowpointers

import (
	"github.com/algorithmsandpatterns/models"
	"github.com/algorithmsandpatterns/utils"
)

func maxProfitForShare(prices []int) int {
	maxProfit := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxProfit = utils.Max(maxProfit, profit)
		} else {
			l = r
		}
		r++
	}
	return maxProfit
}

func FindCycleInLinkedList(head *models.LinkedListNode) bool {
	slow, fast := head, head
	detetecedCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			detetecedCycle = true
		}
	}
	return detetecedCycle
}
func power(digit, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result = digit * result

	}
	return result
}

func sumOfSquareOfNumbers(n int) int {
	totalSum := 0
	for n > 0 {
		digit := n % 10
		n = n / 10
		totalSum += power(digit, 2)
	}
	return totalSum

}

func isHappy(n int) bool {
	slowpointer := n

	fastPointer := sumOfSquareOfNumbers(n)
	for slowpointer != fastPointer && fastPointer != 1 {
		slowpointer = sumOfSquareOfNumbers(slowpointer)
		fastPointer = sumOfSquareOfNumbers(sumOfSquareOfNumbers(fastPointer))

	}
	return fastPointer == 1

}
