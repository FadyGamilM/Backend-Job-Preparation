package linkedlist

import "fmt"

type node struct {
	val  string
	next *node
}

func (n node) GetValue() string {
	return n.val
}

func (n node) GetNext() *node {
	return n.next
}

type NodeBuilder struct {
	val  string
	next *node
}

func (nb *NodeBuilder) Val(v string) {
	nb.val = v
}

func (nb *NodeBuilder) Next(n *node) {
	nb.next = n
}

func (nb *NodeBuilder) Build() *node {
	return &node{
		val:  nb.val,
		next: nb.next,
	}
}

type LinkedList struct {
	// hold the address of the Head Node only and you can go trough the entire list
	head   *node
	length int
}

func (l *LinkedList) PreAppend(newNode *node) {
	// put the curr head into a temp node
	temp := l.head

	// replace the curr head by the newNode
	l.head = newNode

	// set the next of the new updated head to be the temp which is the old head
	l.head.next = temp

	// update the length of the linkedlist
	l.length += 1
}

func (l LinkedList) Explore() {
	currNode := node{}
	currNode.val = l.head.val
	currNode.next = l.head.next
	for i := 0; i < l.length; i++ {
		fmt.Println("[node] : ", currNode, " [value] = ", currNode.val, " [next node] : ", currNode.next)
		if currNode.next == nil {
			return
		}
		currNode = *currNode.next
	}
}

func (l *LinkedList) DeleteByValue(valToDelete string) {
	currNode := l.head
	prevNode := &node{}
	nextNode := &node{}

	if currNode.val == valToDelete {
		l.head = currNode.next
		l.length--
		return
	}

	// start from the node number 2 (after the head node)
	//  curr  -> nil
	//	[1]   -> [2] -> [3] -> [4]

	// the nextNode should be updated to the currNode
	//  the currNode should be updated to the nextNode
	// the prevNode should be updated to the currNode
	//  prev  -> curr  -> next -> nil
	//	[1]   -> [2]   -> [3]  -> [4]
	prevNode = currNode
	currNode = currNode.next
	nextNode = currNode.next

	for i := 1; i < l.length; i++ {
		if currNode.val == valToDelete {
			prevNode.next = nextNode
			l.length--
			return
		}
		if currNode.next != nil {
			prevNode = currNode
			currNode = currNode.next
			nextNode = currNode.next
		} else {
			return
		}
	}
}
