package ds

import (
	"errors"
	"fmt"
)

type Queue struct {
	list LinkedList
}

func (q *Queue) Push(value string) { // add tail node
	tmp := &Node{data: value}
	if q.list.IsEmpty() {
		q.list.Head = tmp
		q.list.Tail = tmp
	} else {
		q.list.Tail.Next = tmp
		q.list.Tail = tmp
	}
	q.list.Size++
}
func (q *Queue) Pop() (string, error) { // remove the head
	if q.list.IsEmpty() {
		return "", errors.New("False.")
	}
	tmp := q.list.Head
	q.list.Head = q.list.Head.Next
	q.list.Size--
	if q.list.Head == nil {
		q.list.Tail = nil
	}
	return tmp.data, nil
}

func (q *Queue) PrintQueue() { //print the queue
	curr := q.list.Head
	for curr != nil {
		fmt.Print(curr.data, "-> ")
		curr = curr.Next
	}
	fmt.Println()
}
