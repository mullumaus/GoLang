package main

import (
	"fmt"
	"testing"
)

func TestInsertFirst(t *testing.T) {
	node := Node{data: 10, next: nil}
	l := LinkedList{head: &node}
	l.printf()
	l.insertFirst(88)
	if (l.head.data != 88){
		t.Errorf("insertFirst fail")
	}
	fmt.Println("TestInsertFirst end ----------------")
}
func TestMidPoint(t *testing.T) {
	link := createLink(10)
	link.printf()

	midNode := link.midpoint()
	fmt.Println("mid node =", midNode.data)
	fmt.Println("TestMidPoint end ----------------")
}

func TestFromLast(t *testing.T){
	link := createLink(10)
	link.printf()
	node := link.fromLast(3)
	fmt.Println("Last 3th node",node.data)
	fmt.Println("TestFromLast end ----------------")
}
