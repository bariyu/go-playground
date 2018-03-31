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

// Moves given node to the front of the list
func (ll *LinkedList) MoveFront(node *Node) {
	if node == nil || ll.len == 0 || ll.head == node {
		return
	}

	oldHead := ll.head
	if node == ll.tail {
		ll.tail = node.prev
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil

	ll.head = node
	ll.tail.next = node
	node.prev = ll.tail
	node.next = oldHead
}

// Moves given node to the back of the list
func (ll *LinkedList) MoveBack(node *Node) {
	if node == nil || ll.len == 0 || ll.tail == node {
		return
	}

	oldTail := ll.tail
	if node == ll.head {
		ll.head = node.next
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil

	ll.tail = node
	ll.head.prev = node
	node.prev = oldTail
	node.next = ll.head
}

// Adds new element to the end of the linked list and returns new node
func (ll *LinkedList) PushBack(v interface{}) *Node {
	if ll == nil {
		return nil
	}

	return ll.push(v, false)
}

// Removes given node from the linked list
func (ll *LinkedList) RemoveNode(node *Node) {
	if node == nil {
		return
	}

	next := node.next
	prev := node.prev

	// if this is the last elem in the list
	if ll.len == 1 {
		ll.head = nil
		ll.tail = nil
	} else {
		next.prev = prev
		prev.next = next
		if node == ll.head {
			ll.head = next
		}
		if node == ll.tail {
			ll.tail = prev
		}
	}

	ll.len -= 1
}

// Removes the head of the linked list
func (ll *LinkedList) RemoveHead() {
	if ll == nil || ll.Len() == 0 {
		return
	}
	ll.RemoveNode(ll.Head())
}

// Removes the tail of the linked list
func (ll *LinkedList) RemoveTail() {
	if ll == nil || ll.Len() == 0 {
		return
	}
	ll.RemoveNode(ll.Tail())
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
