package ds

type LinkedList struct {
    Head *Node
    Tail *Node
    Size int
}

func (l *LinkedList) Insert(value string)// insert at the tail
func (l *LinkedList) InsertAt(position int, value string) error //inserts at a position, returns true if success or false if not, like if pos doesn't exist
func (l *LinkedList) Remove(value string) error // remove first occurrence of the value
func (l *LinkedList) RemoveAll(value string) error //remove all occurrences of a value
func (l *LinkedList) RemoveAt(pos int) error // remove at a position, if index exists
func (l *LinkedList) IsEmpty() bool // checks if the linked list is empty
func (l *LinkedList) GetSize() int // get size of ll
func (l *LinkedList) Reverse() //reverse the list
func (l *LinkedList) PrintList() //print the list 