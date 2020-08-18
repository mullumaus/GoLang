package main

import "fmt"

/*
Implement two queues as arguments and combines the contents of each into a new third
queue. THe third  queue should contain the "alterating" content of the twp queues
The function should handle queues of different lengths without inserting
undefined into the new one.
Do not access the array inside of any queue, only use 'add', 'remove'
and 'peek' function
Implement a peek function in queue class, Peek should return the last
element (the next one to be returned) from the queue without removing it
*/

//Queue is struct
type Queue struct {
	data []int
}

func (q *Queue) add(n int) {
	q.data = append(q.data, n)
}

func (q *Queue) remove() int {
	d := q.data[0]
	q.data = q.data[1:len(q.data)]

	return d
}

func (q *Queue) peek() bool {
	if len(q.data) > 0 {
		return true
	}
	return false
}

func (q Queue) length() int {
	return len(q.data)
}

func (q Queue) printf() {
	fmt.Println(q.data)
}

func weave(q1 Queue, q2 Queue) Queue {
	q := Queue{}
	for q1.peek() || q2.peek() {
		if q1.peek() {
			q.add(q1.remove())
		}
		if q2.peek() {
			q.add(q2.remove())
		}
	}
	return q
}

func main() {
	q1 := Queue{}
	q2 := Queue{}
	for i := 0; i < 10; i++ {
		q1.add(i)
	}
	q1.printf()

	for i := 10; i < 33; i++ {
		q2.add(i)
	}
	q2.printf()

	q3 := weave(q1, q2)
	q3.printf()
}
