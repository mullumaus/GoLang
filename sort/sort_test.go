package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := generateArray(20)
	bubbleSort(arr)
	fmt.Println(arr)
}

func TestSelectionSort(t *testing.T) {
	arr := generateArray(20)
	selectionSort(arr)
	fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {

}
