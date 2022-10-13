package main

import "fmt"

type Node struct {
	key int;
	leftNode *Node;
	rightNode *Node;
}


func (n *Node) Insert(newKey int) {
	if (n.Search(newKey)) {
		fmt.Println("already inserted");
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

	if (n == nil) {
		return false
	}
	if (n.key == key) {
		return true;
	}
	if (n.key < key) {
		return n.rightNode.Search(key)
	} else if n.key > key {
		return n.leftNode.Search(key);
	}
	return false;
}


func (n *Node) SearchAndInsert(key int) {
	// searches the binary Tree and returns the node where it should be the root of

}

func main() {
	tree := &Node{key: 100};
	tree.Insert(10);
	tree.Insert(20);
	tree.Insert(90);
	tree.Insert(101);
	tree.Insert(100);
	tree.Insert(203);
	tree.Insert(44);
	tree.Insert(7);
	tree.Insert(52);
	tree.Insert(88);
	tree.Insert(276);
	tree.Insert(200);
	tree.Insert(300);
	tree.Insert(300);
	fmt.Println(tree.leftNode, tree.rightNode, tree);
	fmt.Println(tree.Search(76));
	fmt.Println(tree.Search(203));
}