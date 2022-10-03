package main

import "fmt"


const ArrayBufferSize int = 64;

type HeapStructure struct {
	heap [] *int
}



func (h *HeapStructure) checkHeapAlloc () bool{
	return len(h.heap) == cap(h.heap)
}

func (h *HeapStructure) reassignHeap () {
	var oldData []*int = h.heap[:];
	var length = len(oldData);
	h.heap = make([]*int, length, length + ArrayBufferSize)
	h.heap = append(oldData, h.heap...)
}

func InitHeap () *HeapStructure{
	newArray := make([]*int, 0, ArrayBufferSize)
	newHeap := &HeapStructure{heap: newArray};
	return newHeap
}

func main () {
	fmt.Println("hello")
}