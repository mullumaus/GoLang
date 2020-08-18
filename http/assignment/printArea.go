package main

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{sideLength: 10}
	ta := triangle{base: 10, height: 40}
	printArea(sq)
	printArea(ta)
}

func (ta triangle) getArea() float64 {
	return 0.5 * ta.base * ta.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	println(s.getArea())
}
