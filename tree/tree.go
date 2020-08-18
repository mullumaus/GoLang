package main

import "fmt"

//Node is
type Node struct {
	data     int
	children []*Node
}

func (node *Node) add(n int) *Node {
	newNode := Node{data: n, children: nil}
	node.children = append(node.children, &newNode)
	return &newNode
}

/*
Given some data, look at each child of the current node and remove any node with data===data
*/
func (node *Node) remove(n int) {
	for idx, child := range node.children {
		if child.data == n {
			node.children = append(node.children[:idx], node.children[idx+1:]...)
		}
	}
}

func (node *Node) printf() {
	fmt.Print("Parent:", node.data, ", children: ")
	fmt.Println(node.children)
}

/*
Tree is tree traversal is interating through a tree
*/
type Tree struct {
	root *Node
}

/* breadth first traverse: recursive */
func traverseBFRecursive(node *Node) {
	if node != nil {
		//fmt.Println(node.data)
		for _, child := range node.children {
			fmt.Print(child.data, ", ")
		}
	}
	for _, child := range node.children {
		traverseBFRecursive(child)
	}
}

/* breadth first traverse: non-recursive
use a queue to store the node */
func traverseBFNonRecursive(node *Node) {
	qu := []*Node{node}
	head := 0
	for head != len(qu) {
		for _, child := range qu[head].children {
			qu = append(qu, child)
		}
		head++
	}
	for _, n := range qu {
		fmt.Println(n.data)
	}
}

func traserveDeepthRecursive(node *Node) {
	if node != nil {
		fmt.Print(node.data, ", ")
	}
	for _, child := range node.children {
		traserveDeepthRecursive(child)
	}
}

/* deepth traverse */
func traserveDeepthNonRecursive(node *Node) {
	arr := []*Node{node}
	top := 1
	for top > 0 {
		parent := arr[top-1]
		fmt.Print(parent.data, ", ")
		top--
		for i := len(parent.children) - 1; i >= 0; i-- {
			if top >= len(arr) {
				arr = append(arr, parent.children[i])
			} else {
				arr[top] = parent.children[i]
			}
			top++
		}
	}

}

func level(root *Node) {
	counter := []int{}
	arr := []*Node{root, nil}
	pointer := 0
	for {
		nodeCount := 0
		for arr[pointer] != nil {
			if len(arr[pointer].children) > 0 {
				arr = append(arr, arr[pointer].children...)
			}
			nodeCount++
			pointer++
		}
		counter = append(counter, nodeCount)
		if pointer == len(arr)-1 {
			break
		}
		arr = append(arr, nil)
		pointer++
	}
	fmt.Println(counter)
}

func generateTree() *Node {
	root := Node{data: 10}
	node1 := root.add(11)
	node2 := root.add(12)
	root.add(13)
	node1.add(111)
	node1.add(112)
	node2.add(121)
	node2.add(122)
	return &root
}
