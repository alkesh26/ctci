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
	fmt.Printf("\n")
}

func (list *List) kthFromLast(k int) {
	if k > list.length || k == 0 {
		fmt.Println("K should be greater than zero and less than length of list")
	} else {
		currentNode := list.start
		previousNode := list.start

		for k > 0 {
			currentNode = currentNode.next
			k--
		}

		for currentNode != nil {
			previousNode = previousNode.next
			currentNode = currentNode.next
		}

		fmt.Println("Kth element is", previousNode.data)
	}
}

func main() {
	list := &List{}
	list.addNode(&Node{data: 1})
	list.addNode(&Node{data: 12})
	list.addNode(&Node{data: 9})
	list.addNode(&Node{data: 3})
	list.addNode(&Node{data: 7})
	list.addNode(&Node{data: 10})
	list.addNode(&Node{data: 5})

	list.display()

	k := 2
	list.kthFromLast(k)
}
