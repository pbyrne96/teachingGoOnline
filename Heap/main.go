package main

import "fmt"


type MaxHeap struct {
	nodes  []int;
}

func (h *MaxHeap) Insert (key int) {
	h.nodes = append(h.nodes, key)
	h.maxHeapifyUp(len(h.nodes) -1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.nodes[parent(index)] < h.nodes[index] {
		h.swop(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.nodes) -1;
	l,r := left(index), right(index);
	childToCompare := 0;

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.nodes[l] > h.nodes[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.nodes[index] < h.nodes[childToCompare] {
			h.swop(index, childToCompare);
			l, r = left(childToCompare), right(childToCompare)
		} else {
			return
		}
	}
}

func parent (i int) int {
	return (i-1)/2
}

func left(i int) int {
	return 2*i +1
}

func right (i int) int {
	return 2*1 +2
}

func (h *MaxHeap) swop(i1,i2 int) {
	h.nodes[i1], h.nodes[i2] = h.nodes[i2], h.nodes[i1];
}

func (h* MaxHeap) Extract () int{
	extracted := h.nodes[0];
	length := len(h.nodes)-1;

	if length == 0 {
		return -1;
	}

	h.nodes[0] = h.nodes[length];
	h.nodes = h.nodes[:length]
	h.maxHeapifyDown(0)
	// h.maxHeapifyUp(length)
	return extracted
}

// complexity in heap depends on how high the heap is at the time or the height of the tree
// could also be considered o(log n) to extract data to insert is complicated as you need ro orginize the heap

func main () {
		m := &MaxHeap{};
		buildHeap := []int{10, 20, 30, 5, 7, 11, 13, 15, 17}

		for _, v := range buildHeap {
			m.Insert(v);
		}
		fmt.Println(m)

		for i:=0; i < 5; i++ {
			m.Extract();
			fmt.Println(m)
			i++;
		}
		fmt.Println(m)
}