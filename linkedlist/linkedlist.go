package linkedlist

import (
	"fmt"
)

// Package linkedlist implements a doubly linked list.
// API is inspired from container/list

// A Node in the linkedlist.
type Node struct {
	Value      interface{}
	next, prev *Node
}

// Next returns the next node or nil.
func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

// Prev returns the previous node or nil.
func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

// LinkedList struct holds head, tail and len.
type LinkedList struct {
	head, tail *Node
	len        int
}

// Returns a new initialized linked list.
func New() *LinkedList {
	return &LinkedList{head: nil, tail: nil, len: 0}
}

// Returns the length of linked list.
func (ll *LinkedList) Len() int {
	return ll.len
}

// Returns head node of the linked list
func (ll *LinkedList) Head() *Node {
	if ll == nil {
		return nil
	}
	return ll.head
}

// Returns tail node of the linked list
func (ll *LinkedList) Tail() *Node {
	if ll == nil {
		return nil
	}
	return ll.tail
}

// Clears the linked list
func (ll *LinkedList) Clear() {
	if ll != nil {
		ll.head = nil
		ll.tail = nil
		ll.len = 0
	}
}

// Add new element to the list and return new node.
func (ll *LinkedList) push(v interface{}, front bool) *Node {
	if ll == nil {
		return nil
	}

	newNode := &Node{Value: v}
	newNode.next = newNode
	newNode.prev = newNode

	if ll.len == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		newNode.prev = ll.tail
		ll.head.prev = newNode
		ll.tail.next = newNode
	}

	if front {
		ll.head = newNode
	} else {
		ll.tail = newNode
	}

	ll.len += 1

	return newNode
}

// Adds new element to front of the linked list and returns new node
func (ll *LinkedList) PushFront(v interface{}) *Node {
	if ll == nil {
		return nil
	}

	return ll.push(v, true)
}

// Adds new element to the end of the linked list and returns new node
func (ll *LinkedList) PushBack(v interface{}) *Node {
	if ll == nil {
		return nil
	}

	return ll.push(v, false)
}

// Prints the list used for debugging
func (ll *LinkedList) printList() {
	fmt.Println("---Start---")

	fmt.Printf("len: %d\n", ll.Len())

	runner := ll.Head()

	for i := 0; i < ll.Len(); i++ {
		fmt.Printf("elem index: %d, value: %v\n", i, runner.Value)
		runner = runner.Next()
	}

	fmt.Print("---End---")
}
