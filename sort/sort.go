package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) {
	arrSize := len(arr)
	for i := 0; i < arrSize; i++ {
		for j := 0; j < arrSize-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}

func selectionSort(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		minIndex := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			tmp := arr[minIndex]
			arr[minIndex] = arr[i]
			arr[i] = tmp
		}
	}
}

//join two sort arrays together
func merge(left []int, right []int) []int {
	result := []int{}
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
		} else {
			result = append(result, right[r])
		}
		l++
		r++
	}
	if l < len(left) {
		result = append(result, left[l:]...)
	}
	if r < len(right) {
		result = append(result, right[r:]...)
	}
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[0:mid]
	right := arr[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func generateArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		arr[i] = r.Intn(100)
	}
	fmt.Println(arr)
	return arr
}
