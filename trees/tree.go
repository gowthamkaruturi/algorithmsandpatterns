package trees

import (
	"fmt"

	"github.com/algorithmsandpatterns/models"
)

type Tree struct {
	Root *models.Node
}

func (t *Tree) PrintTree() {
	fmt.Print("Tree : ")
	printPreOrder(t.Root)
	fmt.Println()
}

func printPreOrder(n *models.Node) {
	if n == nil {
		return
	}
	fmt.Print(n.Value, " ")
	printPreOrder(n.Left)
	printPreOrder(n.Right)
}
func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.Root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *models.Node {
	curr := &models.Node{
		Value: arr[start],
		Left:  nil,
		Right: nil,
	}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.Left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.Right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}
