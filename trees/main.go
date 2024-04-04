package main

import (
	"fmt"
)

func main() {
	tree := NewTree()
	fmt.Println(tree)
	tree.insert(10)
	fmt.Println(tree)
	tree.insert(10)
	tree.insert(11)
	tree.insert(0)
	tree.insert(4)
	tree.insert(15)
	fmt.Println(tree)
}
