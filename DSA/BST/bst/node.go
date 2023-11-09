package bst

type node struct {
	val   int
	left  *node
	right *node
}

func NewNode(v int) *node {
	return &node{
		val:   v,
		left:  nil,
		right: nil,
	}
}
