/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。
*/
package main

/*Node is */
type Node struct {
	value int
	next  *Node
}

func deleteNode(head *Node, val int) *Node {
	if head == nil {
		return nil
	}
	if head.value == val {
		head = head.next
		return head
	}
	tmp := head
	for tmp.next != nil {
		if tmp.next.value == val {
			tmp.next = tmp.next.next
			return head
		}
		tmp = tmp.next
	}
	return head
}
func main() {

}
