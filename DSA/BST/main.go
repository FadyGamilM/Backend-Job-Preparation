package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	root := newNode(5)

	/*
			5
		  /   \
		nil    nil
	*/
	root.left = newNode(3)
	/*
			5
		  /   \
		3    nil
	*/
	root.left.left = newNode(1)
	/*
					5
				  /   \
				3    nil
			  /   \
		     1    nil
	*/
	root.left.right = newNode(4)
	/*
					5
				  /   \
				3    nil
			  /   \
		     1    4
	*/
	root.right = newNode(6)
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

	bst.add(9)
	sb = strings.Builder{}
	bst.InOrderTraverse(&sb)
	fmt.Println(sb.String())

	bst.add(0)
	sb = strings.Builder{}
	bst.InOrderTraverse(&sb)
	fmt.Println(sb.String())

	_, ok := bst.search(20)
	if !ok {
		log.Println("not found")
	} else {
		log.Println("found")
	}
}
