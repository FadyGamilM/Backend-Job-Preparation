package main

type node struct {
	val   int
	left  *node
	right *node
}

func newNode(v int) *node {
	return &node{
		val:   v,
		left:  nil,
		right: nil,
	}
}
