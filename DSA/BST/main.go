package main

import (
	"fmt"
	"log"
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
	bst.InOrderTraverse(&sb)
	return sb.String()
}

func NewBST(r *node) *bst {
	return &bst{root: r, length: 1}
}

// fidn the min node in whatever the tree/subtree given and used in the remove method
func (bst bst) FindMin(root *node) *node {
	curr := root
	for curr != nil && curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (bst *bst) Remove(root *node, val int) *node {
	return bst.remove_logic(root, val)
}

func (bst *bst) remove_logic(root *node, val int) *node {

	if root == nil {
		return nil
	}

	if val < root.val {
		root.left = bst.remove_logic(root.left, val)
	} else if val > root.val {
		root.right = bst.remove_logic(root.right, val)
	} else {

		// thats means the root is the value we need to remove, lets process it based on its case

		// when we reach the leafs when its a node that has no left or right, so just delete it and return nil to its parent
		if root.left == nil && root.right == nil {
			return nil
		} else if root.left == nil && root.right != nil {
			return root.right
		} else if root.left != nil && root.right == nil {
			return root.left
		} else { // when it has 2 childs [baad]
			minNodeFromRightSubTree := bst.FindMin(root.right)
			root.val = minNodeFromRightSubTree.val
			root.right = bst.remove_logic(root.right, minNodeFromRightSubTree.val)
		}
	}
	return root
}

// the wrapper
func (bst bst) InOrderTraverse(sb *strings.Builder) {
	bst.inOrderTraverse_Logic(sb, bst.root)
}

// the actual algorithm (inner work)
func (bst bst) inOrderTraverse_Logic(sb *strings.Builder, currRoot *node) {
	// Base Case => if null return
	if currRoot == nil {
		return
	}

	// work on the left
	bst.inOrderTraverse_Logic(sb, currRoot.left)
	// then work on the current node itself (the work on means do what you need to do with the node, here we need to print (traverse) the node so we add it to the builder)
	sb.WriteString(strconv.Itoa(currRoot.val) + " => ")
	// then work on the right
	bst.inOrderTraverse_Logic(sb, currRoot.right)
}

func (bst *bst) Add(v int) {
	bst.root = bst.add_logic(bst.root, v)
	bst.length++
}

func (bst *bst) add_logic(root *node, v int) *node {

	// if the tree is empty, so the root is nil
	// -> create a new node and return it to be the root
	if root == nil {
		n := NewNode(v)
		return n
	}

	if v < root.val {
		root.left = bst.add_logic(root.left, v)
	} else if v >= root.val {
		root.right = bst.add_logic(root.right, v)
	}

	// to return the root of the sub-tree you are working it in the currenct recursive func
	return root
}

func (bst bst) Search(v int) (*node, bool) {
	return bst.search_logic(bst.root, v)
}

func (bst bst) search_logic(root *node, v int) (*node, bool) {
	if root == nil {
		return nil, false
	}

	// base case
	if root.val == v {
		return root, true
	} else if v < root.val {
		// work on the left sub tree
		return bst.search_logic(root.left, v)
	} else if v > root.val {
		// work on the right sub tree
		return bst.search_logic(root.right, v)
	}

	return nil, false
}

func main() {
	root := NewNode(5)

	/*
			5
		  /   \
		nil    nil
	*/
	root.left = NewNode(3)
	/*
			5
		  /   \
		3    nil
	*/
	root.left.left = NewNode(1)
	/*
					5
				  /   \
				3    nil
			  /   \
		     1    nil
	*/
	root.left.right = NewNode(4)
	/*
					5
				  /   \
				3    nil
			  /   \
		     1    4
	*/
	root.right = NewNode(6)
	/*
					5
				  /   \
				3      6
			  /   \
		     1    4
	*/
	bst := NewBST(root)
	bst.length = 5

	sb := strings.Builder{}
	bst.InOrderTraverse(&sb)
	fmt.Println(sb.String())

	bst.Add(9)
	sb = strings.Builder{}
	bst.InOrderTraverse(&sb)
	fmt.Println(sb.String())

	bst.Add(0)
	sb = strings.Builder{}
	bst.InOrderTraverse(&sb)
	fmt.Println(sb.String())

	_, ok := bst.Search(20)
	if !ok {
		log.Println("not found")
	} else {
		log.Println("found")
	}

	bst.root = bst.Remove(bst.root, 3)

	sb = strings.Builder{}
	bst.InOrderTraverse(&sb)
	fmt.Println(sb.String())

}
