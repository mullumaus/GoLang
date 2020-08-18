/*
#输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
*/

package main

/*ListNode is */
type ListNode struct {
	val  int
	next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	n1 := l1
	n2 := l2
	var head *ListNode
	if n1.val > n2.val {
		head = n2
		n2 = n2.next
	} else {
		head = n1
		n1 = n1.next
	}
	node := head
	var tmp *ListNode
	for n1 != nil && n2 != nil {
		if n1.val < n2.val {
			tmp = n1.next
			node.next = n1
			n1 = tmp
		} else {
			tmp = n2.next
			node.next = n2
			n2 = tmp
		}
		if node.next != nil {
			node = node.next
		}
	}
	if n1 != nil {
		node.next = n1
	}
	if n2 != nil {
		node.next = n2
	}
	return head
}

func main() {

}
