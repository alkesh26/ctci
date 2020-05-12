package main

import "fmt"

// Node struct
type Node struct {
	data int
	next *Node
}

// List type
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

func findIntersection(list1, list2 *List) {
	if list1.length == 0 || list2.length == 0 {
		fmt.Println("One of the list is empty so it has no node")
		return
	}

	var shorterListNode *Node
	var longerListNode *Node
	diff := list1.length - list2.length
	if diff < 0 {
		diff *= -1
	}

	if list1.length > list2.length {
		shorterListNode = list2.start
		longerListNode = list1.start
	} else {
		shorterListNode = list1.start
		longerListNode = list2.start
	}

	for ; diff > 0; diff-- {
		longerListNode = longerListNode.next
	}

	for shorterListNode != longerListNode {
		shorterListNode = shorterListNode.next
		longerListNode = longerListNode.next
	}

	fmt.Println(shorterListNode.data)
}

func main() {
	list1 := &List{}
	intersectingNode := &Node{data: 57}

	list1.addNode(&Node{data: 105})
	list1.addNode(&Node{data: 6})
	list1.addNode(&Node{data: 12})
	list1.addNode(&Node{data: 90})
	list1.addNode(intersectingNode)
	list1.addNode(&Node{data: 45})
	list1.addNode(&Node{data: 34})
	list1.addNode(&Node{data: 112})
	list1.addNode(&Node{data: 879})

	list1.display()

	list2 := &List{}
	list2.addNode(&Node{data: 11})
	list2.addNode(&Node{data: 21})
	list2.addNode(intersectingNode)
	list2.length = 7

	list2.display()

	findIntersection(list1, list2)
}
