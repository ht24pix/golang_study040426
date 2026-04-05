package main

//"crypto/sha256"
import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

// insert data
func (l *LinkedList) Insert(val int) {
	newNode := &Node{data: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// dispaly linked list
func (l *LinkedList) DisplayLinkedLst() {
	current := l.head
	for current != nil {
		fmt.Printf("%d - > ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// delete node
func (l *LinkedList) DeleteNode(val int) {

	if l.head == nil {
		return
	}
	//case 1: if data is a head
	if l.head.data == val {
		l.head = l.head.next
		return
	}
	current := l.head
	//case 2: if data is not head, tranverse to find the node
	for current.next != nil && current.next.data != val {
		current = current.next
	}
	//If the value was found, skip the target node
	if current.next != nil {
		current.next = current.next.next
	}

}

// search data in linked list
func (l *LinkedList) SearchVal(val int) {
	current := l.head
	for current != nil {
		if current.data == val {
			fmt.Printf("found: %d\n", current.data)
			return
		}
		current = current.next
	}
	fmt.Println("Not found")
}

func main() {

	link := LinkedList{}
	link.Insert(3)
	link.Insert(4)
	link.Insert(60)
	link.Insert(70)

	fmt.Println("Linked list example")
	link.DisplayLinkedLst()
	link.SearchVal(60)

	fmt.Println("Linked list (after delete)")
	link.DeleteNode(60)
	link.DisplayLinkedLst()
}
