package main

import (
	"fmt"
)

type MaxHeap struct {
	elements []int
}


// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.elements = append(h.elements, key)
	h.maxHeapifyUp(len(h.elements) - 1)
}


func (h *MaxHeap) Extract() int {

	if len(h.elements) == 0 {
		fmt.Println("Cannot extract from an empty heap")
		return -1
	}

	extracted := h.elements[0]

	lastIndex := len(h.elements) - 1

	// removes the last element and make it the first element
	h.elements[0] = h.elements[lastIndex]
	h.elements = h.elements[:lastIndex]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp heapifies bottom to top with higher keys above
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.elements[index] > h.elements[parent(index)] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown heapifies top to bottom with higher keys above
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.elements) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {

		if l == lastIndex { // if left child is the last child
			childToCompare = l
		} else if h.elements[l] > h.elements[r] { // if left child is larger
			childToCompare = l
		} else { // if right child is larger
			childToCompare = r
		}


		// compare element value of current index to the larger child and swap if lower
		if h.elements[index] < h.elements[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
		
	}
}



// get the parent index
func parent(i int) int {
	return (i - 1 ) / 2
}

// get the left child index
func left(i int) int {
	return (i * 2) + 1
}

// get the right child index
func right(i int) int {
	return (i * 2 ) + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}



func main() {
	myHeap := &MaxHeap{}
	loadToHeap := []int{3, 5, 8, 2, 20, 4, 12}

	for _, v := range loadToHeap {
		myHeap.Insert(v)
		fmt.Println(myHeap)
	}

	fmt.Println("Done inerting")

	for i := 0; i < 5; i++ {
		myHeap.Extract()
		fmt.Println(myHeap)
	}
}