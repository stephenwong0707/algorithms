package linkedlist

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    result := &ListNode{}
    node := result
    slow := head
    for slow != nil {
        fast := slow.Next
        if fast != nil && slow.Val == fast.Val {
            for fast != nil && slow.Val == fast.Val {
                fast = fast.Next
            }
        } else {
            node.Next = &ListNode{Val: slow.Val}
            node = node.Next
        }
        slow = fast
    }
    return result.Next
}