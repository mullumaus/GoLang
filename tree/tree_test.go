package main

import (
	"fmt"
	"testing"
)

func TestNodeAdd(t *testing.T) {
	root := generateTree()
	root.printf()
	fmt.Println("end TestNodeAdd-------------")
}

func TestTraverseBFRecursive(t *testing.T) {
	root := generateTree()
	traverseBFRecursive(root)
	fmt.Println("end TestTraverseBFRecursive-------------")
}

func TestTraverseBFNonRecursive(t *testing.T) {
	root := generateTree()
	traverseBFNonRecursive(root)
	fmt.Println("end TestTraverseBFNonRecursive-------------")
}

func TestTraserveDeepthRecursive(t *testing.T) {
	root := generateTree()
	traserveDeepthRecursive(root)
	fmt.Println("end TestTraserveDeepthRecursive-------------")
}

func TestTraserveDeepthNonRecursive(t *testing.T) {
	root := generateTree()
	traserveDeepthNonRecursive(root)
	fmt.Println("end TesttraserveDeepthNonRecursive--------------------------")
}

func TestLevel(t *testing.T) {
	root := generateTree()
	level(root)
	fmt.Println("end TestLevel--------------------------")
}

func TestBinaryTreeInsert(t *testing.T) {
	root := generateBalanceBinaryTree()
	printf(root)
}

func TestBinaryTreeContains(t *testing.T) {
	root := generateBalanceBinaryTree()
	node := root.contain(30)
	if node == nil || node.data != 30 {
		t.Error("Error: do not find the node")
	}
}

func TestValidateBinaryTree(t *testing.T) {
	root := generateBalanceBinaryTree()
	result := validate(root)
	if !result {
		t.Error("The tree is not binary tree")
	}
}
