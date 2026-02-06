package ds

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Insert(value string) { // insert at the tail
	tmp := &Node{data: value}
	if l.IsEmpty() {
		l.Head = tmp
		l.Tail = tmp
	} else {
		l.Tail.Next = tmp
		l.Tail = tmp
	}
	l.Size++
}
func (l *LinkedList) InsertAt(position int, value string) error { //inserts at a position, returns true if success or false if not, like if pos doesn't exist
	if position < 1 || position > l.Size+1 {
		return errors.New("False.")
	}
	tmp := &Node{data: value}
	if position == 1 {
		tmp.Next = l.Head
		l.Head = tmp
		if l.Size == 0 {
			l.Tail = tmp
		}
		l.Size++
		return nil
	}
	curr := l.Head
	for i := 1; i < position-1; i++ {
		curr = curr.Next
	}
	tmp.Next = curr.Next
	curr.Next = tmp
	if tmp.Next == nil {
		l.Tail = tmp
	}
	l.Size++
	return nil
}
func (l *LinkedList) Remove(value string) error { // remove first occurrence of the value
	curr := l.Head
	var prev *Node
	for curr != nil {
		if curr.data == value {
			if prev == nil {
				l.Head = curr.Next
			} else {
				prev.Next = curr.Next
			}
			if curr.Next == nil {
				l.Tail = prev
			}
			l.Size--
			return nil
		}
		prev = curr
		curr = curr.Next
	}
	return errors.New("False.")
}
func (l *LinkedList) RemoveAll(value string) error { //remove all occurrences of a value
	curr := l.Head
	var prev *Node
	removed := false
	for curr != nil {
		if curr.data == value {
			removed = true
			l.Size--
			if prev == nil {
				l.Head = curr.Next
				curr = l.Head
			} else {
				prev.Next = curr.Next
				curr = prev.Next
			}
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	if !removed {
		return errors.New("False.")
	}
	l.Tail = prev
	if l.Size == 0 {
		l.Tail = nil
	}
	return nil
}
func (l *LinkedList) RemoveAt(pos int) error { // remove at a position, if index exists
	if pos < 1 || pos > l.Size {
		return errors.New("False.")
	}
	if pos == 1 {
		l.Head = l.Head.Next
		l.Size--
		if l.Size == 0 {
			l.Tail = nil
		}
		return nil
	}
	tmp := l.Head
	for i := 1; i < pos; i++ {
		tmp = tmp.Next
	}
	if tmp.Next.Next == nil {
		l.Tail = tmp
	}
	tmp.Next = tmp.Next.Next
	l.Size--
	return nil
}
func (l *LinkedList) IsEmpty() bool { // checks if the linked list is empty
	return l.Size == 0
}
func (l *LinkedList) GetSize() int { // get size of ll
	return l.Size
}
func (l *LinkedList) Reverse() { //reverse the list
	var prev *Node
	curr := l.Head
	l.Tail = l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}
func (l *LinkedList) PrintList() { //print the list
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.data, "-> ")
		curr = curr.Next
	}
	fmt.Println()
}
