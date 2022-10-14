package main

import "fmt"

// global Alphabet size
const AlphabetSize = 26;

type Node struct {
	children [AlphabetSize]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func CreateTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (h *Node) Insert(str string) bool {
	return true
}

func (h *Node) Search(str string) {

}

func (h *Node) Delete(str string) bool {
	return true
}


func main () {
	test := CreateTrie();
	fmt.Println("hello", test);
}