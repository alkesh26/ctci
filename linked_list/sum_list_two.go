package main

import (
	"fmt"
)

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

func sumLists(num1, num2 *Node, carry *int) (node *Node) {
	if num1 == nil {
		return
	}

	result := new(Node)

	result.next = sumLists(num1.next, num2.next, carry)
	sum := num1.data + num2.data + *carry
	*carry = sum / 10
	sum = sum % 10
	result.data = sum

	return result
}

func (list *List) addCarryAndDisplaySum(carry int) {
	if carry > 0 {
		p := &Node{data: 1}
		p.next = list.start
		list.start = p
	}
	list.display()
}

func prependZeros(num *List, diff int) {
	for ; diff > 0; diff-- {
		node := &Node{data: 0}
		node.next = num.start
		num.start = node
	}
}

func addLists(num1, num2 *List) (node *Node) {
	diff := num1.length - num2.length
	if diff < 0 {
		diff *= -1
	}

	if num1.length > num2.length {
		prependZeros(num2, diff)
	} else {
		prependZeros(num1, diff)
	}

	carry := 0
	return sumLists(num1.start, num2.start, &carry)
}

func calculateSum(num1, num2 *List) {
	carry := 0
	result := new(Node)

	if num1.length != num2.length {
		result = addLists(num1, num2)
	} else {
		result = sumLists(num1.start, num2.start, &carry)
	}

	final := &List{}
	final.start = result
	final.addCarryAndDisplaySum(carry)
}

func main() {
	num1 := &List{}
	num1.addNode(&Node{data: 2})
	num1.addNode(&Node{data: 1})
	num1.addNode(&Node{data: 6})

	num2 := &List{}
	num2.addNode(&Node{data: 9})
	num2.addNode(&Node{data: 6})
	num2.addNode(&Node{data: 5})
	num2.addNode(&Node{data: 5})

	num1.display()
	num2.display()

	fmt.Println("Sum is:: ")
	calculateSum(num1, num2)
}
