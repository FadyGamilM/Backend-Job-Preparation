package main

import (
	"log"

	"github.com/FadyGamilM/dsa/linkedlist"
)

func main() {
	linkedList := &linkedlist.LinkedList{}

	nodeA1uilder := &linkedlist.NodeBuilder{}
	nodeA1uilder.Val("fady")

	node2Builder := &linkedlist.NodeBuilder{}
	node2Builder.Val("ahmed")

	node3Builder := &linkedlist.NodeBuilder{}
	node3Builder.Val("marwan")

	node4Builder := &linkedlist.NodeBuilder{}
	node4Builder.Val("sa3ed")

	node1 := nodeA1uilder.Build()
	node2 := node2Builder.Build()
	node3 := node3Builder.Build()
	node4 := node4Builder.Build()

	log.Printf("node 1 : %v", node1)
	log.Printf("node 2 : %v", node2)
	log.Printf("node 3 : %v", node3)
	log.Printf("node 4 : %v", node4)

	linkedList.PreAppend(node2)
	log.Println("after preappending node 2 into the list =>")
	log.Printf("node 2 : %v", node2)

	linkedList.PreAppend(node1)
	log.Println("after preappending node 1 into the list =>")
	log.Printf("node 1 : %v", node1)

	linkedList.PreAppend(node3)
	log.Println("after preappending node 3 into the list =>")
	log.Printf("node 3 : %v", node3)

	linkedList.PreAppend(node4)
	log.Println("after preappending node 4 into the list =>")
	log.Printf("node 4 : %v", node4)

	linkedList.Explore()

	linkedList.DeleteByValue("sa3ed")

	log.Println("after deletion...")

	linkedList.Explore()
}
