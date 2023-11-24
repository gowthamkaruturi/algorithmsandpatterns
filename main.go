package main

import (
	"fmt"

	"github.com/algorithmsandpatterns/trees"
	"github.com/algorithmsandpatterns/twopointers"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := trees.LevelOrderBinaryTree(arr)
	t.PrintTree()

	fmt.Println(twopointers.IsPalindrome("raceacar"))

}
