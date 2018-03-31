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

	llist.MoveFront(head)
	if head != llist.Head() {
		t.Errorf("42 should be the head of the linked list after moving it to front")
	}
	llist.printList()

	if llist.Len() != 1 {
		t.Errorf("length of the linked list should be 1")
	}
}

func TestMultiNode(t *testing.T) {
	llist := New()

	ensureEmptyList(llist, t)

	// 1 <-> 1
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

	// 0 <-> 1 <-> 1 <-> 2
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

	// 2 <-> 0 <-> 1 <-> 1
	llist.MoveFront(tail)
	if llist.Head() != tail {
		t.Errorf("old tail should be new head")
	}
	if llist.Tail().Value != 1 {
		t.Errorf("new tail should be 1")
	}

	node2 := llist.Head()
	// 0 <-> 1 <-> 1 <-> 2
	llist.MoveBack(node2)
	if llist.Head().Value != 0 {
		t.Errorf("0 should be new head")
	}
	if llist.Tail() != node2 {
		t.Errorf("new tail should be 2")
	}

	llist.printList()
}

func TestMoveFront(t *testing.T) {
	llist := New()

	ensureEmptyList(llist, t)

	// 1 <-> 2
	node1 := llist.PushFront(1)
	node2 := llist.PushBack(2)

	llist.printList()

	// 1 <-> 2
	llist.MoveFront(node1)
	if llist.Head() != node1 {
		t.Errorf("head should 1 after moving it to head")
	}
	if llist.Tail() != node2 {
		t.Errorf("tail should 2 after moving 1 to head")
	}

	llist.printList()

	// 2 <-> 1
	llist.MoveFront(node2)
	if llist.Head() != node2 {
		t.Errorf("head should 2 after moving tail to head")
	}
	if llist.Tail() != node1 {
		t.Errorf("tail should 1 after moving tail to head")
	}
}

func TestMoveBack(t *testing.T) {
	llist := New()

	ensureEmptyList(llist, t)

	// 1 <-> 2
	node1 := llist.PushFront(1)
	node2 := llist.PushBack(2)

	llist.printList()

	// 1 <-> 2
	llist.MoveBack(node2)
	if llist.Head() != node1 {
		t.Errorf("head should 1 after moving 2 to tail")
	}
	if llist.Tail() != node2 {
		t.Errorf("tail should 2 after moving 2 to tail")
	}

	llist.printList()

	// 2 <-> 1
	llist.MoveBack(node1)
	if llist.Head() != node2 {
		t.Errorf("head should 2 after moving 1 to tail")
	}
	if llist.Tail() != node1 {
		t.Errorf("tail should 1 after moving 1 to tail")
	}
}

func TestRemoveNode(t *testing.T) {
	llist := New()

	ensureEmptyList(llist, t)

	llist.PushFront(1)
	node := llist.Head()

	llist.RemoveNode(node)

	ensureEmptyList(llist, t)

	// 0 <-> 1 <-> 2
	node1 := llist.PushBack(1)
	node2 := llist.PushBack(2)
	node0 := llist.PushFront(0)

	// 0 <-> 2
	llist.RemoveNode(node1)

	if llist.Head() != node0 {
		t.Errorf("head of the linked list should be the node with value 0")
	}

	if llist.Tail() != node2 {
		t.Errorf("tail of the linked list should be the node with value 2")
	}

	if llist.Len() != 2 {
		t.Errorf("length of the linked list should be 2")
	}

	// 0 <-> 2 <-> 3
	node3 := llist.PushBack(3)

	// 2 <-> 3
	llist.RemoveNode(node0)

	if llist.Head() != node2 {
		t.Errorf("head of the linked list should be the node with value 2")
	}

	if llist.Tail() != node3 {
		t.Errorf("tail of the linked list should be the node with value 3")
	}

	// -1 <-> 2 <-> 3 <-> 5
	nodem1 := llist.PushFront(-1)
	node5 := llist.PushBack(5)

	// -1 <-> 2 <-> 3
	llist.RemoveNode(node5)

	if llist.Head() != nodem1 {
		t.Errorf("tail of the linked list should be the node with value -1")
	}

	if llist.Tail() != node3 {
		t.Errorf("tail of the linked list should be the node with value 3")
	}

	if llist.Len() != 3 {
		t.Errorf("length of the linked list should be 3")
	}

	// 2 <-> 3
	llist.RemoveHead()

	if llist.Head() != node2 {
		t.Errorf("head of the linked list should be the node with value 2")
	}

	// 2
	llist.RemoveTail()

	if llist.Tail() != node2 {
		t.Errorf("tail of the linked list should be the node with value 2")
	}

	llist.RemoveHead()

	ensureEmptyList(llist, t)
}
