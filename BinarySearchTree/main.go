package main

import "fmt"

type Node struct {
	key int;
	leftNode *Node;
	rightNode *Node;
}


func (n *Node) Insert(newKey int) {
	// takes a key as a param that is not already in the searchTree
	if (!n.Search(newKey)) {
		return
	}
	if (n.key < newKey) {
		// move right
		if (n.rightNode == nil) {
			n.rightNode = &Node{key: newKey};
		} else {
			n.rightNode.Insert(newKey);
		}
	} else if n.key > newKey{
		// move left
		if (n.leftNode == nil) {
			n.leftNode = &Node{key : newKey}
		} else {
			n.leftNode.Insert(newKey);
		}
	}

}


func (n *Node) Search(key int) bool {
	if (n == nil || n.key == key) {
		return false
	}
	if (n.key < key) {
		return n.rightNode.Search(key)
	} else if n.key > key {
		return n.leftNode.Search(key);
	}
	return true;
}

func main() {
	tree := &Node{key: 100};
	tree.Insert(200);
	tree.Insert(300);
	fmt.Println(tree.leftNode, tree.rightNode, tree);
}