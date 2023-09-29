package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

type bst struct {
	root   *node
	length int
}

func (bst bst) Print() string {
	sb := strings.Builder{}
	bst.InOrderTraverseWrapper(&sb)
	return sb.String()
}

func NewBST(r *node) *bst {
	return &bst{root: r, length: 1}
}

// the wrapper
func (bst bst) InOrderTraverseWrapper(sb *strings.Builder) {
	bst.InOrderTraverse(sb, bst.root)
}

// the actual algorithm (inner work)
func (bst bst) InOrderTraverse(sb *strings.Builder, currRoot *node) {
	// Base Case => if null return
	if currRoot == nil {
		return
	}

	// work on the left
	bst.InOrderTraverse(sb, currRoot.left)
	// then work on the current node itself (the work on means do what you need to do with the node, here we need to print (traverse) the node so we add it to the builder)
	sb.WriteString(strconv.Itoa(currRoot.val) + " => ")
	// then work on the right
	bst.InOrderTraverse(sb, currRoot.right)
}

func main() {
	root := NewNode(5)

	root.left = NewNode(3)
	root.left.left = NewNode(1)
	root.left.right = NewNode(4)

	root.right = NewNode(6)

	bst := NewBST(root)
	bst.length = 5

	sb := strings.Builder{}
	bst.InOrderTraverseWrapper(&sb)
	fmt.Println(sb.String())
}
