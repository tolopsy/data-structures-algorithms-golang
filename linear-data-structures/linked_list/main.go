package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) prepend(value int) {
	formerHead := l.head
	l.head = &Node{data: value}
	l.head.next = formerHead // formerHead is moved to the next cell
}


func (l *LinkedList) getLastNode() *Node {
	var lastNode *Node
	for picker := l.head; picker != nil; picker = picker.next {
		if picker.next == nil {
			lastNode = picker
			break
		}
	}
	return lastNode
}


func (l *LinkedList) append(value int) {
	lastNode := l.getLastNode()
	newNode := &Node{data: value}

	if lastNode != nil {
		lastNode.next = newNode
	}

}


func PrintList(l LinkedList) {
	for picker := l.head; picker != nil; picker = picker.next {
		fmt.Printf("%d -> ", picker.data)
	}
	fmt.Println("end")
}


func (l LinkedList) getNodeWithValue(value int) *Node {
	var nodeToFind *Node

	for picker := l.head; picker != nil; picker = picker.next {
		if picker.data == value {
			nodeToFind = picker
			break
		}
	}	

	return nodeToFind
}


func (l LinkedList) insertAfter(value, insertValue int) {
	var nodeToInsert *Node
	for picker := l.head; picker != nil; picker = picker.next {

		if picker.data == value {

			nodeToInsert = &Node{data: insertValue}
			nextNode := picker.next
			picker.next = nodeToInsert
			nodeToInsert.next = nextNode
			break

		}
	}	

	if nodeToInsert == nil {
		fmt.Printf("value %v is not in LinkedList\n", value)
	}
}


func (l *LinkedList) insertBefore(value, insertValue int) {
	if l.head.data == value {
		l.prepend(insertValue)
		return
	}

	for picker := l.head; picker.next != nil; picker = picker.next {
		if picker.next.data == value {
			nodeToInsert := &Node{data: insertValue}
			nextNode := picker.next
			picker.next = nodeToInsert
			nodeToInsert.next = nextNode
			break
		}
	}
}


func main() {
	newList := LinkedList{}

	value1 := 7
	newList.prepend(value1)
	PrintList(newList) // 7
	fmt.Printf("The last value is %v \n", newList.getLastNode())

	value2 := 5
	newList.prepend(value2)
	PrintList(newList)  // 5 -> 7

	value3 := 3
	newList.prepend(value3)
	PrintList(newList) // 3 -> 5 -> 7

	value4 := 2
	newList.append(value4)
	PrintList(newList) // 3 -> 5 -> 7 -> 2

	newList.insertAfter(3, 1)
	PrintList(newList)
	

	newList.insertBefore(5, 9)
	PrintList(newList) // 3 -> 5 -> 7 -> 9 -> 2

	fmt.Printf("The last node is %v \n", newList.getLastNode())
}