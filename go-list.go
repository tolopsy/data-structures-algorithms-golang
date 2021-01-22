package main

import (
	"fmt"
	"container/list"
)


func main() {
	var intList list.List
	intList.PushBack(12)
	intList.PushBack(17)
	intList.PushBack("Hi")
	intList.PushBack(5)

	//fmt.Println(intList)
	//fmt.Println(intList.Front())

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}