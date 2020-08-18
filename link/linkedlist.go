package main

import "fmt"

//Node is
type Node struct {
	data int
	next *Node
}

//LinkedList is a link
type LinkedList struct {
	head *Node
}

//create a new node from argument 'data' and assigns the resulting node
//to the head of the link
func (l *LinkedList) insertFirst(d int) {
	node := Node{data: d, next: nil}
	if l.head != nil {
		node.next = l.head
	}
	l.head = &node
}

func (l *LinkedList) size() int {
	node := l.head
	count := 0
	for node != nil {
		node = node.next
		count++
	}
	return count
}

func (l *LinkedList) getFirst() *Node {
	return l.head
}

func (l *LinkedList) getLast() *Node {
	node := l.head
	for node != nil {
		if node.next == nil {
			return node
		}
		node = node.next
	}
	return nil
}

func (l *LinkedList) clear() {
	l.head = nil
}

func (l *LinkedList) removeFirst() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *LinkedList) removeLast() {
	if l.head == nil {
		return
	}
	if l.head.next == nil {
		l.head = nil
		return
	}
	preNode := l.head
	node := preNode.next
	for node.next != nil {
		preNode = node
		node = node.next
	}
	preNode.next = nil
}

func (l *LinkedList) insertLast(data int) {
	lastNode := l.getLast()
	newNode := Node{data: data, next: nil}
	if lastNode == nil {
		l.head = &newNode
		return
	}
	lastNode.next = &newNode

}

func (l *LinkedList) getAt(index int) *Node {
	count := 0
	node := l.head
	for node != nil {
		if count == index {
			return node
		}
		count++
		node = node.next
	}
	return nil
}

func (l *LinkedList) removeAt(index int) {
	if index == 0 {
		l.head = l.head.next
		return
	}
	preNode := l.getAt(index - 1)
	if preNode != nil && preNode.next != nil {
		preNode.next = preNode.next.next
	}

}

func (l *LinkedList) insertAt(index int, data int) {
	newNode := Node{data: data, next: nil}
	if l.head == nil {
		l.head = &newNode
		return
	}
	if index == 0 {
		newNode.next = l.head.next
		l.head = &newNode
		return
	}
	preNode := l.getAt(index - 1)
	if preNode == nil {
		l.insertLast(data)
		return
	}
	newNode.next = preNode.next
	preNode.next = &newNode

}

func (l *LinkedList) printf() {
	node := l.head
	for node != nil {
		fmt.Print(node.data, ", ")
		node = node.next
	}
	fmt.Println()
}

/*
return the middle node of a linked list
if the list an even number of elements, return
the node at the end of the first half of the list.
Do not use a counter variable,
do not retrieve the size of the list, and only iterate through the list
one time
*/

func (l LinkedList) midpoint() *Node {
	slow := l.head
	fast := l.head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

/*
check if a link is circular
*/

func (l LinkedList) circular() bool {
	slow := l.head
	fast := l.head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

/*
Given a linked list and integer n, return the element n node from
the last node in the list. Do not all the size method of the linked list
Assume that n will always be less than length of the list
*/
func (l LinkedList) fromLast(n int) *Node {
	fast := l.head
	slow := l.head
	for i := 0; i < n; i++ {
		fast = fast.next
	}
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	return slow

}

func createLink(n int) LinkedList {
	link := LinkedList{}
	if n <= 0 {
		return link
	}
	node := &Node{data: 0, next: nil}
	link.head = node
	for i := 1; i < n; i++ {
		tmpNode := &Node{data: i, next: nil}
		node.next = tmpNode
		node = tmpNode
	}
	return link
}
