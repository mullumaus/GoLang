package main

import "fmt"

/*
write a function that accepts an integer N
and returns a N*N spiral matrix
*/
type intslice []int

func matrix(n int) {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	counter := 1
	startRow := 0
	endRow := n - 1
	startColumn := 0
	endColumn := n - 1
	for startRow <= endRow && startColumn <= endColumn {
		for i := startColumn; i <= endColumn; i++ {
			result[startRow][i] = counter
			counter++
		}
		startRow++
		for i := startRow; i <= endRow; i++ {
			result[i][endColumn] = counter
			counter++
		}
		endColumn--
		for i := endColumn; i >= startColumn; i-- {
			result[endRow][i] = counter
			counter++
		}
		endRow--
		for i := endRow; i >= startRow; i-- {
			result[i][startColumn] = counter
			counter++
		}
		startColumn++
	}
	fmt.Println(result)
}

func main() {
	matrix(3)
}
