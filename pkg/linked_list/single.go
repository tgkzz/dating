package linked_list

import "errors"

type Node struct {
	Val  interface{}
	Next *Node
}

type SingleLinkedList struct {
	Head *Node
}

var (
	ImpossibleInsertion = errors.New("cannot insert val into list: ")
	ImpossibleDeletion  = errors.New("cannot delete after position: ")
	NoData              = errors.New("could not find data")
)

// InsertEnd insert data at the of SingleLinkedList
func (list *SingleLinkedList) InsertEnd(val interface{}) {
	insertNode := &Node{
		Val:  val,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = insertNode
		return
	}

	curr := list.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = insertNode
}

// InsertFront insert data at the start of SingleLinkedList
func (list *SingleLinkedList) InsertFront(val interface{}) {
	if list.Head == nil {
		insertNode := &Node{Val: val, Next: nil}
		list.Head = insertNode
		return
	}

	insertNode := &Node{Val: val, Next: list.Head}
	list.Head = insertNode
}

// InsertAfterPosition insert data after special position
func (list *SingleLinkedList) InsertAfterPosition(pos int, val interface{}) error {
	if pos < 0 {
		return errors.Join(ImpossibleInsertion, errors.New("position cannot be negative"))
	}

	if list.Head == nil {
		if pos == 0 {
			node := &Node{Val: val, Next: nil}
			list.Head = node
			return nil
		}

		return errors.Join(ImpossibleInsertion, errors.New("position is out of bound of empty list"))
	}

	counter := 0

	curr := list.Head
	for curr != nil {
		if counter == pos {
			curr.Next = &Node{Val: val, Next: curr.Next}
			return nil
		}

		curr = curr.Next
		counter++
	}

	return errors.Join(ImpossibleInsertion, errors.New("size of linked list is smaller than given position"))
}

// GetSize returns size of list
func (list *SingleLinkedList) GetSize() int {
	if list.Head == nil {
		return 0
	}

	counter := 0
	curr := list.Head
	for curr != nil {
		curr = curr.Next
		counter++
	}

	return counter
}

// GetPosition return the position of data in list
func (list *SingleLinkedList) GetPosition(data interface{}) (int, error) {
	curr := list.Head
	pos := 0

	for curr.Next != nil {
		if data == curr.Val {
			return pos, nil
		}

		curr = curr.Next
		pos++
	}

	return 0, NoData
}

// DeleteFront deletes first element in linked list
func (list *SingleLinkedList) DeleteFront() {
	if list.Head != nil {
		list.Head = list.Head.Next
	}

	return
}

// DeleteBack deletes last element in linked list
func (list *SingleLinkedList) DeleteBack() {
	if list.Head == nil {
		return
	}

	if list.Head.Next == nil {
		list.Head = nil
		return
	}

	curr := list.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}

	curr.Next = nil
}

// DeleteAfterPos deletes element after special position
func (list *SingleLinkedList) DeleteAfterPos(pos int) error {
	if pos < 0 {
		return errors.Join(ImpossibleDeletion, errors.New("position cannot be negative"))
	}

	if list.Head == nil {
		return errors.Join(ImpossibleDeletion, errors.New("list is empty"))
	}

	counter := 0
	curr := list.Head
	if curr.Next != nil {
		if counter == pos {
			curr.Next = curr.Next.Next
			return nil
		}
		curr = curr.Next
		counter++
	}

	return errors.Join(ImpossibleDeletion, errors.New("position is out of bound"))
}
