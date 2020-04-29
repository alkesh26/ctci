package main

import "fmt"

// Node struct for a linked list single node
type Node struct {
	data int
	next *Node
}

// List struct for a maintaining a list of nodes
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
	return
}

func (list *List) display() {
	currentNode := list.start

	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (list *List) removeDuplicates() {
	hash := make(map[int]int)
	currentNode := list.start
	newList := &List{}

	for currentNode != nil {
		if hash[currentNode.data] == 0 {
			newList.addNode(&Node{data: currentNode.data})
			hash[currentNode.data]++
		}
		currentNode = currentNode.next
	}

	newList.display()
}

func main() {
	list := &List{}
	list.addNode(&Node{data: 1})
	list.addNode(&Node{data: 2})
	list.addNode(&Node{data: 19})
	list.addNode(&Node{data: 4})
	list.addNode(&Node{data: 6})
	list.addNode(&Node{data: 5})
	list.addNode(&Node{data: 6})
	list.addNode(&Node{data: 19})
	list.addNode(&Node{data: 8})

	list.display()
	list.removeDuplicates()
}
