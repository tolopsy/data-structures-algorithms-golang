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


func (l *LinkedList) addToEnd(value int) {
	lastNode := l.getLastNode()
	newNode := &Node{data: value}
	if lastNode != nil {
		lastNode.next = newNode
	}

}


func (l LinkedList) PrintList() {
	for picker := l.head; picker != nil; picker = picker.next {
		fmt.Printf("%d -> ", picker.data)
	}

	fmt.Println("end")
}


func main() {
	newList := LinkedList{}

	value1 := 7
	newList.prepend(value1)
	newList.PrintList() // 7
	fmt.Printf("The last value is %v \n", newList.getLastNode())

	value2 := 5
	newList.prepend(value2)
	newList.PrintList()  // 5 -> 7

	value3 := 3
	newList.prepend(value3)
	newList.PrintList() // 3 -> 5 -> 7

	value4 := 2
	newList.addToEnd(value4)
	newList.PrintList() // 3 -> 5 -> 7 -> 2

	fmt.Printf("The last node is %v \n", newList.getLastNode())
}