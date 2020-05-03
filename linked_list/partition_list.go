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

func (list *List) partitionList(x int) {
	var prevStart, prevLast, nextStart, nextLast *Node

	currentNode := list.start
	for currentNode != nil {
		if currentNode.data < x {
			if prevStart == nil {
				prevStart = currentNode
			} else {
				prevLast.next = currentNode
			}
			prevLast = currentNode
		} else {
			if nextStart == nil {
				nextStart = currentNode
			} else {
				nextLast.next = currentNode
			}
			nextLast = currentNode
		}
		currentNode = currentNode.next
	}

	if nextStart != nil {
		nextLast.next = nil
	}

	if prevStart != nil {
		prevLast.next = nil
		prevLast.next = nextStart
		list.start = prevStart
	} else {
		list.start = nextStart
	}

	list.display()
}

func main() {
	list := &List{}
	list.addNode(&Node{data: 104})
	list.addNode(&Node{data: 4})
	list.addNode(&Node{data: 63})
	list.addNode(&Node{data: 89})
	list.addNode(&Node{data: 34})
	list.addNode(&Node{data: 65})
	list.addNode(&Node{data: 129})

	list.display()

	x := 70
	list.partitionList(x)
}
