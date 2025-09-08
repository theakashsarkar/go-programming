package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBack(data int)
{
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nill {
		current = current.next
	}
}

func (list LinkedList) display() {
	current := list.head
	for current != nill {
		fmt.Printf("%d --> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList()
	list.insertAtBack(10)
	list.insertAtBack(20)
	list.insertAtBack(30)
	list.display()

}
