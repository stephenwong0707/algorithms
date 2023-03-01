package linkedlist

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head ListNode
	node := &head
	for list1 != nil || list2 != nil {
		if list1 == nil || (list2 != nil && list1.Val > list2.Val) {
			node.Next = &ListNode{
				Val: list2.Val,
			}
			node = node.Next
			list2 = list2.Next
		} else {
			node.Next = &ListNode{
				Val: list1.Val,
			}
			node = node.Next
			list1 = list1.Next
		}
	}
	return head.Next
}
