package ds

import (
	"fmt"
)

type Stack struct {
	list LinkedList
}

func (s *Stack) Push(value string) { // add new head node
	tmp := &Node{data: value}
	if s.list.IsEmpty() {
		s.list.Head = tmp
		s.list.Tail = tmp
	} else {
		tmp.Next = s.list.Head
		s.list.Head = tmp
	}
	s.list.Size++
}
func (s *Stack) Pop() (string, bool) { // remove the head
	if s.list.IsEmpty() {
		return "", false
	}
	tmp := s.list.Head
	s.list.Head = s.list.Head.Next
	s.list.Size--
	return tmp.data, true
}

func (s *Stack) PrintStack() { //print the stack
	curr := s.list.Head
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.Next
	}
	fmt.Println()
}
