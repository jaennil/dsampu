package main

type node struct {
	value int
	left *node
	right *node
}

func (node node) hasRight() bool {
	return node.right != nil
}

func (node node) hasLeft() bool {
	return node.left != nil
}
