package models

type Node struct {
	Value       int
	Left, Right *Node
}

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}
