package main

import "fmt"

// Node struct
type Node struct {
	data int
	next *Node
}

// List struct
type List struct {
	length int
	start  *Node
}

func (list *List) addNode(node *Node) {
	if list.length == 0 {
		list.start = node
	} else {
		currentNode := list.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = node
	}
	list.length++
}

func (list *List) display() {
	currentNode := list.start

	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("")
}

func isPalindrome(left **Node, right *Node) bool {
	if right == nil {
		return true
	}

	isPal := isPalindrome(left, right.next)
	if isPal == false {
		return false
	}

	isAPal := ((*left).data == right.data)
	*left = (*left).next

	return isAPal
}

func main() {
	list := &List{}
	list.addNode(&Node{data: 1})
	list.addNode(&Node{data: 2})
	list.addNode(&Node{data: 4})
	list.addNode(&Node{data: 2})
	list.addNode(&Node{data: 1})

	list.display()

	result := isPalindrome(&list.start, list.start)
	fmt.Println("result", result)
}
