package main

import "fmt"

// Node type struct
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

	fmt.Println("\n")
}

func (list *List) displayMiddle() {
	if list.length == 0 {
		fmt.Println("no nodes")
		return
	} else if list.length <= 2 {
		fmt.Println("middle node", list.start.data)
		return
	}

	currentNode := list.start
	doubleNode := list.start
	for doubleNode.next != nil && doubleNode.next.next != nil {
		currentNode = currentNode.next
		doubleNode = doubleNode.next.next
	}

	fmt.Println("middle node", currentNode.data)
}

func main() {
	list := &List{}
	list.addNode(&Node{data: 1})
	list.addNode(&Node{data: 2})
	list.addNode(&Node{data: 3})
	list.addNode(&Node{data: 4})
	list.addNode(&Node{data: 5})
	list.addNode(&Node{data: 6})
	list.addNode(&Node{data: 7})
	list.addNode(&Node{data: 8})
	list.addNode(&Node{data: 9})
	list.addNode(&Node{data: 10})
	list.addNode(&Node{data: 11})
	list.addNode(&Node{data: 12})
	list.addNode(&Node{data: 13})

	list.display()
	list.displayMiddle()
}
