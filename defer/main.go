package main

import "fmt"

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func d(a int) int {
	return a + 1
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func main() {
	fmt.Println(c())
}
