package leetcode

/*
21,合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	head := root
	for l1 != nil || l2 != nil {
		if l1 == nil {
			root.Next = l2
			break
		}
		if l2 == nil {
			root.Next = l1
			break
		}
		if l1.Val >= l2.Val {
			root.Next = l2
			l2 = l2.Next
		} else {
			root.Next = l1
			l1 = l1.Next
		}
		root = root.Next
	}
	return head.Next
}
