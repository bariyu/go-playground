package linkedlist

import "testing"

func ensureEmptyList(llist *LinkedList, t *testing.T) {
	if llist.Head() != nil {
		t.Errorf("empty linked list should have nil head")
	}

	if llist.Tail() != nil {
		t.Errorf("empty linked list should have nil tail")
	}

	if llist.Len() != 0 {
		t.Errorf("empty linked list should have len: 0")
	}
}

func TestEmpty(t *testing.T) {
	llist := New()

	if llist == nil {
		t.Errorf("cannot initialize new linked list")
	}

	ensureEmptyList(llist, t)
}
func TestSingleNode(t *testing.T) {
	llist := New()

	ensureEmptyList(llist, t)

	llist.PushFront(42)

	head := llist.Head()
	tail := llist.Tail()

	if head == nil {
		t.Errorf("linked list head should not be nil")
	}

	if tail == nil {
		t.Errorf("linked list tail should not be nil")
	}

	if head != tail {
		t.Errorf("linked list tail and head should be the same element")
	}

	if head.Value != 42 || tail.Value != 42 {
		t.Errorf("linked list head and tail values should be 42")
	}

	if llist.Len() != 1 {
		t.Errorf("length of the linked list should be 1")
	}

	llist.Clear()

	ensureEmptyList(llist, t)

	llist.PushBack("42")

	head = llist.Head()
	tail = llist.Tail()

	if head == nil {
		t.Errorf("linked list head should not be nil")
	}

	if tail == nil {
		t.Errorf("linked list tail should not be nil")
	}

	if head != tail {
		t.Errorf("linked list tail and head should be the same element")
	}

	if head.Value != "42" || tail.Value != "42" {
		t.Errorf("linked list head and tail values should be 42")
	}

	if llist.Len() != 1 {
		t.Errorf("length of the linked list should be 1")
	}
}

func TestMultiNode(t *testing.T) {
	llist := New()

	ensureEmptyList(llist, t)

	llist.PushFront(1)
	llist.PushFront(1)

	head := llist.Head()
	tail := llist.Tail()

	if head == tail {
		t.Errorf("head and tail should be different")
	}

	if head.Value != tail.Value {
		t.Errorf("head and tail should have same value")
	}

	if llist.Len() != 2 {
		t.Errorf("length of linked list should be 2")
	}

	llist.PushFront(0)
	llist.PushBack(2)

	head = llist.Head()
	tail = llist.Tail()

	if head.Value != 0 {
		t.Errorf("first element should be 0")
	}

	if tail.Value != 2 {
		t.Errorf("last element should be 2")
	}

	if head.prev != tail {
		t.Errorf("heads prev should point to tail")
	}

	if tail.next != head {
		t.Errorf("tails next should point to head")
	}

	llist.printList()
}
