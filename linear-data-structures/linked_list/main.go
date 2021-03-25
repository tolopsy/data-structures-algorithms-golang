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


func (l LinkedList) PrintList() {
	picker := l.head
	fmt.Print("Start ")
	for picker != nil {
		fmt.Printf("-> %d ", picker.data)
		picker = picker.next
	}
}

func main() {
	newList := LinkedList{}
	node1 := 7
	node2 := 5
	node3 := 9
	node4 := 2
	newList.prepend(node1)
	newList.prepend(node2)
	newList.prepend(node3)
	newList.prepend(node4)
	newList.PrintList()
}