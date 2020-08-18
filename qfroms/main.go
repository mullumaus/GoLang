package main

import "fmt"

/*
Implement a Queue data structure using two stacks.
Do not create an array inside of the queue class
Queue should implement the methods 'add' 'remove' and 'peek'
*/

//Stack is FILO
type Stack struct {
	data []int
}

func (s *Stack) push(n int) {
	s.data = append(s.data, n)
}

func (s *Stack) pop() int {
	d := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]

	return d
}

func (s Stack) peek() bool {
	if len(s.data) > 0 {
		return true
	}
	return false
}

//Queue is
type Queue struct {
	sIn  Stack
	sOut Stack
}

func (q *Queue) add(n int) {
	q.sIn.data = append(q.sIn.data, n)
}

func (q *Queue) remove() int {
	if !q.sOut.peek() {
		for q.sIn.peek() {
			q.sOut.push(q.sIn.pop())
		}
	}
	return q.sOut.pop()
}

func (q Queue) peek() bool {
	return q.sIn.peek() || q.sOut.peek()
}

func (q Queue) printf() {
	fmt.Println(q)
}

func main() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		q.add(i)
	}
	for q.peek() {
		q.remove()
		q.printf()
	}
}
