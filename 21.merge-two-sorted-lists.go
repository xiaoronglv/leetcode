/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	for node := result; !(l1 == nil && l2 == nil); node = node.Next {
		if l1 == nil {
			node.Next = l2
			break
		} else if l2 == nil {
			node.Next = l1
			break
		} else if l1.Val <= l2.Val {
			node.Next = l1
			l1 = l1.Next	
		} else if l2.Val < l1.Val {
			node.Next = l2
			l2 = l2.Next
		}
	}
	return result.Next
}

