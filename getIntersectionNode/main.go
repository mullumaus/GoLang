//输入两个链表，找出它们的第一个公共节点
package main

//ListNode is
type ListNode struct {
	val  int
	next *ListNode
}

func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != nil || b != nil {
		if a == nil {
			a = headB
		}
		if b == nil {
			b = headA
		}
		if a == b {
			return a
		}
		a = a.next
		b = b.next
	}
	return nil
}

func getIntersectionNodeMethod2(headA *ListNode, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		if a != nil {
			a = a.next
		} else {
			a = headB
		}
		if b != nil {
			b = b.next
		} else {
			b = headA
		}
	}
	return a
}
func main() {

}
