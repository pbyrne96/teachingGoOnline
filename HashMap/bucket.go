package main

import "fmt"

type bucket struct {
	head *bucketNode

}

type bucketValue struct {
	value int
}

type bucketNode struct {
	key string
	value *bucketValue
	next *bucketNode
}



func (b *bucket) insert(k string, v int) bool{
	// improve this to double linked list
	if (!b.search(k)) {
		newNode := &bucketNode{key:k, value: &bucketValue{value: v}}
		newNode.next = b.head
		b.head = newNode
		return true
	} else {
		fmt.Println("already exists")
		return false
	}

}

func (b *bucket) search(k string) bool{
	if (b == nil || b.head == nil) {
		return false
	}

	if (b.head.key == k) {
		return true
	}

	current := b.head
	for current != nil {
		if (current.key == k) {
			return true
		}
		current = current.next
	}

	return false
}

func (b *bucket) delete(k string) bool{
	// delete a node in bucket by severing the connection
	if (b.head == nil) {
		return false
	}

	if (b.head.key == k) {
		b.head = b.head.next
		return true
	}

	previousNode := b.head
	for previousNode.next != nil {
		if(previousNode.next.key == k) {
			previousNode.next = previousNode.next.next
			return true
		}
		previousNode = previousNode.next
	}
	return false
}