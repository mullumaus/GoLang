package main

import "fmt"

func main() {
	node := Node{data: 10, next: nil}
	l := LinkedList{head: &node}
	l.printf()
	l.insertFirst(88)
	l.insertFirst(18)
	l.insertFirst(9)
	l.printf()
	fmt.Println("size=", l.size())
	first := l.getFirst()
	fmt.Println("First=", (*first).data)
	last := l.getLast()
	fmt.Println("Last =", (*last).data)
	fmt.Println("Index=2", l.getAt(2).data)
	l.insertAt(2, 38)
	l.printf()
	l.removeAt(3)
	l.printf()
	l.clear()
	fmt.Println("clear link, size=", l.size())
}
