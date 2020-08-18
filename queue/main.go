package main

import "fmt"

//Queue is a queue struct
type Queue struct {
	data []int
}

func createQueue(n int) Queue {
	q := Queue{}
	return q
}

func (q *Queue) add(n int) {
	q.data = append(q.data, n)
}

func (q *Queue) remove() int {
	rval := q.data[0]
	q.data = q.data[1:len(q.data)]
	return rval
}

func main() {
	q := createQueue(10)
	for i := 0; i < 10; i++ {
		q.add(i)
	}
	q.remove()
	fmt.Println(q)
}
