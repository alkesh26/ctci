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

func (list *List) reverse() {
	if list.length == 0 {
		return
	}

	var p *Node
	q := list.start

	for q != nil {
		temp := q.next
		q.next = p
		p = q
		q = temp

		if temp != nil {
			temp = temp.next
		}
	}

	list.start = p
}

func sumLists(num1, num2 *List) {
	carry := 0

	if num1.length == 0 {
		num2.display()
		return
	}

	if num2.length == 0 {
		num1.display()
		return
	}

	num1Iterator := num1.start
	num2Iterator := num2.start
	sumList := &List{}

	for num1Iterator != nil || num2Iterator != nil {
		sum := 0
		if num1Iterator != nil {
			sum += num1Iterator.data
		}
		if num2Iterator != nil {
			sum += num2Iterator.data
		}
		sum += carry

		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		dataNode := sum % 10
		sumList.addNode(&Node{data: dataNode})

		if num1Iterator != nil {
			num1Iterator = num1Iterator.next
		}

		if num2Iterator != nil {
			num2Iterator = num2Iterator.next
		}
	}

	if carry > 0 {
		sumList.addNode(&Node{data: carry})
	}

	sumList.reverse()
	sumList.display()
}

func main() {
	num1 := &List{}
	num1.addNode(&Node{data: 4})
	num1.addNode(&Node{data: 1})
	num1.addNode(&Node{data: 6})

	num2 := &List{}
	num2.addNode(&Node{data: 5})
	num2.addNode(&Node{data: 9})
	num2.addNode(&Node{data: 2})

	num1.display()
	num2.display()

	fmt.Println("Sum is:: ")
	sumLists(num1, num2)
}
