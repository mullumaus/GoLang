package main

import "fmt"

/*
Given an array and chunk size, divide the array into many subarray,
where each subarray is of length size
for example:
chunk([1,2,3,4],2) -->[[1,2],[3,4]]
chunk([1,2,3,4,5],2) -->[[1,2],[3,4],[5]]
*/

type chunkdata []int

func chunk(input []int, size int) []chunkdata {
	result := []chunkdata{}
	last := len(input) / size * size
	for i := 0; i < last; i += size {
		subarray := input[i : i+size]
		result = append(result, subarray)
	}
	if last < len(input) {
		result = append(result, input[last:len(input)])
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	size := 5
	fmt.Println(chunk(input, size))
}
