package main

import "fmt"

//balance binary tree
type binNode struct {
	data  int
	left  *binNode
	right *binNode
}

//insert a new node with data at the appropriate location in the tree
func (root *binNode) insert(inData int) {
	newNode := &binNode{data: inData, left: nil, right: nil}
	node := root
	for {
		if node.data > inData {
			if node.left == nil {
				node.left = newNode
				return
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = newNode
				return
			}
			node = node.right
		}
	}
}

/*
implement the 'contains' method for the node.
Contains should accept a 'data' argument and return the node
in the tree with the same value. If the value isn't in the tree
return NULL
*/
func (root *binNode) contain(inData int) *binNode {
	if root == nil {
		return nil
	}
	node := root
	for node != nil {
		if node.data == inData {
			return node
		}
		if node.data > inData {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

/*
Given a node, validate the binary search tree
ensuring that every node's left hand child is less than parent node's
value and that every node's right hand child is greater than the parent
*/
func validate(root *binNode) bool {
	if root == nil {
		return false
	}

	list := []*binNode{root}
	index := 0
	for index != len(list) {
		node := list[index]
		if node.left != nil {
			if node.left.data > node.data {
				return false
			}
			list = append(list, node.left)
		}
		if node.right != nil {
			if node.right.data < node.data {
				return false
			}
			list = append(list, node.left)
		}
		index++
	}
	return true
}

func validateMinMax(node *binNode, min *binNode, max *binNode) bool {
	if node == nil {
		return false
	}
	if max != nil && node.data > max.data {
		return false
	}
	if min != nil && node.data < min.data {
		return false
	}
	if node.left != nil && validateMinMax(node.left, min, node) {
		return true
	}
	if node.right != nil && validateMinMax(node.right, node, max) {
		return false
	}
	return true
}

func traverseBinTree(root *binNode) []*binNode {
	if root == nil {
		return nil
	}
	list := []*binNode{root}
	index := 0
	for index != len(list) {
		node := list[index]
		if node.left != nil {
			list = append(list, node.left)
		}
		if node.right != nil {
			list = append(list, node.right)
		}
		index++
	}
	return list
}

func printf(node *binNode) {
	if node == nil {
		return
	}
	list := traverseBinTree(node)
	for _, node := range list {
		fmt.Print(node.data, ",")
	}
	fmt.Println()
}
func generateBalanceBinaryTree() *binNode {
	root := &binNode{data: 10, left: nil, right: nil}
	root.insert(5)
	root.insert(12)
	root.insert(4)
	root.insert(30)
	root.insert(1)
	return root
}
