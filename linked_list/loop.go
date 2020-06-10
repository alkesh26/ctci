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

func (list *List) hasLoop() {
	slowNode := list.start
	fastNode := slowNode

	for fastNode != nil && fastNode.next != nil {
		slowNode = slowNode.next
		fastNode = fastNode.next.next

		if slowNode == fastNode {
			break
		}
	}

	if fastNode == nil || fastNode.next == nil {
		return
	}

	slowNode = list.start
	for slowNode != fastNode {
		slowNode = slowNode.next
		fastNode = fastNode.next
	}

	fmt.Println("common node", fastNode.data)
}

func main() {
	list := &List{}
	list.addNode(&Node{data: 1})
	list.addNode(&Node{data: 2})
	list.addNode(&Node{data: 3})
	commonNode := &Node{data: 4}
	list.addNode(commonNode)
	list.addNode(&Node{data: 5})
	list.addNode(&Node{data: 6})
	loopedNode := &Node{data: 7}
	list.addNode(loopedNode)
	loopedNode.next = commonNode

	list.hasLoop()
}
