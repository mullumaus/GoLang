package main

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

func main() {

}
