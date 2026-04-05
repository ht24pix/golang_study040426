package main

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}
type DoubleLinkedList struct {
	head *Node
}

func (dlst *DoubleLinkedList) Insert(val int) {
	newNode := &Node{data: val}
	if dlst.head == nil {
		dlst.head = newNode
		return
	}
	current := dlst.head
	for current.next != nil {
		current = current.next
		current.next.previous = current.next
	}
	current.next = newNode
}

// dispaly linked list
func (l *DoubleLinkedList) DisplayLinkedLst() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	fmt.Println("Double Link List")
	link := DoubleLinkedList{}
	link.Insert(3)
	link.Insert(4)
	link.Insert(60)
	link.Insert(70)
	link.DisplayLinkedLst()
}
