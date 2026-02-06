//Nina Shober
//COSC 3750
//31 January 2026
//
/*
	Don't forget to run your go mod init command in your terminal
	Review the assignment instructions for running your code
	All the code you need to write should be put in the /ds/ package files
	Uncomment the import statement for your module name
	you can uncomment the tests in main as you go to test
	The code in main is not an extensive test, you should add more and test your code as needed
*/
package main

import (
	"fmt"
	"linkedlist-nshober/ds"
)

func main() {
	fmt.Println("Only here so the import doesn't leave an error")

	linkedlist := &ds.LinkedList{}
	err := linkedlist.RemoveAt(3)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err2 := linkedlist.Remove("first")
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	err3 := linkedlist.RemoveAll("first")
	if err3 != nil {
		fmt.Println("Error:", err3)
	}
	linkedlist.Insert("first")
	linkedlist.Insert("first")
	linkedlist.Insert("second")
	linkedlist.Insert("third")
	linkedlist.Insert("fourth")
	linkedlist.Insert("fifth")
	linkedlist.PrintList()
	linkedlist.InsertAt(5, "first")
	err4 := linkedlist.InsertAt(-5, "first")
	if err4 != nil {
		fmt.Println("Error:", err4)
	}
	err5 := linkedlist.InsertAt(10, "first")
	if err5 != nil {
		fmt.Println("Error:", err5)
	}
	err6 := linkedlist.RemoveAt(20)
	if err6 != nil {
		fmt.Println("Error:", err6)
	}
	err7 := linkedlist.Remove("sixth")
	if err7 != nil {
		fmt.Println("Error:", err7)
	}
	err8 := linkedlist.RemoveAll("sixth")
	if err8 != nil {
		fmt.Println("Error:", err8)
	}
	linkedlist.PrintList()
	linkedlist.RemoveAt(3)
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")
	linkedlist.Remove("first")
	linkedlist.PrintList()
	linkedlist.RemoveAll("first")
	linkedlist.PrintList()
	fmt.Println("-------------")
	linkedlist.Reverse()
	linkedlist.PrintList()
	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")

	stack := &ds.Stack{}
	stack.Pop()
	stack.Push("first")
	stack.PrintStack()
	stack.Push("second")
	stack.PrintStack()
	stack.Push("third")
	stack.PrintStack()
	data, _ := stack.Pop()
	println("Popped from stack:", data)

	queue := &ds.Queue{}
	data, err9 := queue.Pop()
	if err9 != nil {
		fmt.Println("Error:", err9)
	}
	queue.Push("first")
	queue.PrintQueue()
	queue.Push("second")
	queue.PrintQueue()
	queue.Push("third")
	queue.PrintQueue()
	data, _ = queue.Pop()
	println("Popped from queue:", data)
}
