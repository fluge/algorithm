package leetcode

/**
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。


给定 1->2->3->4, 你应该返回 2->1->4->3.

思路 每3个节点一组，进行交换
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//可能的根节点
	node := &ListNode{}
	root := node
	for head != nil && head.Next != nil {
		node1 := head.Next.Next
		root.Next = head.Next
		root.Next.Next = head
		root = root.Next.Next
		root.Next = nil
		head = node1
	}

	if head != nil {
		root.Next = head
	}
	return node.Next
}
